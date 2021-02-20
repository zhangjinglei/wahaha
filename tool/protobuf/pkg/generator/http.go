package generator

import (
	"errors"
	"fmt"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/extensions/permission"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/tag"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/typemap"
	"google.golang.org/genproto/googleapis/api/annotations"
)

// HTTPInfo http info for method
type HTTPInfo struct {
	App string
	Permission  permission.Permission
	PermissionCode string
	HttpMethod   string
	Path         string
	LegacyPath   string
	NewPath      string
	IsLegacyPath bool
	Title        string
	Description  string
	// is http path added in the google.api.http option ?
	HasExplicitHTTPPath bool
}

type googleMethodOptionInfo struct {
	Method      string
	PathPattern string
	HTTPRule    *annotations.HttpRule
}

// GetHTTPInfo http info of method
func GetHTTPInfo(
	file *descriptor.FileDescriptorProto,
	service *descriptor.ServiceDescriptorProto,
	method *descriptor.MethodDescriptorProto,
	reg *typemap.Registry) *HTTPInfo {
	var (
		title            string
		desc             string
		httpMethod       string
		newPath          string
		explicitHTTPPath bool=true
		perm =permission.Permission_NeedPerm
		permcode string=""
		app=""
	)
	comment, _ := reg.MethodComments(file, service, method)
	//tags := tag.GetTagsInComment(comment.Leading)
	cleanComments := tag.GetCommentWithoutTag(comment.Leading)
	if len(cleanComments) > 0 {
		title = strings.Trim(cleanComments[0], "\n\r ")
		if len(cleanComments) > 1 {
			descLines := cleanComments[1:]
			desc = strings.Trim(strings.Join(descLines, "\n"), "\r\n ")
		} else {
			desc = ""
		}
	} else {
		title = ""
	}

	parsePermission, err2 := ParsePermission(method)
	if err2!=nil {
		//没有定义http扩展
		//不生成http接口
		explicitHTTPPath=false
	}else {
		app=parsePermission.GetApp()
		if parsePermission.GetPerm()==permission.Permission_IgnoreLogin{
			app=""
		}
		_,ok:=parsePermission.GetPattern().(*permission.HttpRule_Get)
		if ok{
			httpMethod = "GET"
			newPath=strings.TrimSpace(parsePermission.GetGet())
		}else {
			_,ok:=parsePermission.GetPattern().(*permission.HttpRule_Post)
			if ok{
				httpMethod = "POST"
				newPath=strings.TrimSpace(parsePermission.GetPost())
			}
		}
		if newPath==""{
			newPath = "/" + file.GetPackage() + "." + service.GetName() + "/" + method.GetName()
		}

		//println("================", parsePermission)
		if parsePermission.GetPerm() == permission.Permission_NeedPerm {
			if strings.TrimSpace(parsePermission.GetPermcode()) == "" {
				panic(errors.New("缺少权限码定义permcode:"+file.GetName()+"=>"+service.GetName() + "." + method.GetName() ))
			}
			if strings.TrimSpace(parsePermission.GetPermgroup()) == "" {
				panic(errors.New("缺少权限分类permgroup定义:"+file.GetName()+"=>"+service.GetName() + "." + method.GetName() ))
			}
		}
		if parsePermission.GetPerm() != permission.Permission_IgnoreLogin{
			if strings.TrimSpace(parsePermission.GetApp()) == "" {
				panic(errors.New("缺少适用系统app定义:"+file.GetName()+"=>"+service.GetName() + "." + method.GetName() ))
			}
		}


		perm=parsePermission.GetPerm()
		permcode=strings.TrimSpace(parsePermission.GetPermcode())
		if parsePermission.GetPerm()==permission.Permission_IgnoreLogin{
			title+="✨适用【所有系统】┈┈┈┈┈✅无需登录,无需权限"
		}else if parsePermission.GetPerm()==permission.Permission_LoginWithNoPermission{
			title+=`✨适用【`+parsePermission.GetApp()+`】`+`┈┈┈┈┈✅✔需登录,无需权限`
		}else {
			title+=`✨适用【`+parsePermission.GetApp()+`】`+`┈┈┈┈┈✅权限码:`+permcode
		}


	}



//END:
	var p = newPath
	param := &HTTPInfo{
		App:app,
		HttpMethod: httpMethod,
		Permission:perm,
		PermissionCode:permcode,
		Path:                p,
		NewPath:             newPath,
		IsLegacyPath:        false,
		Title:               title,
		Description:         desc,
		HasExplicitHTTPPath: explicitHTTPPath,
	}
	if title == "" {
		param.Title = param.Path
	}
	return param
}

func (t *Base) GetHttpInfoCached(file *descriptor.FileDescriptorProto,
	service *descriptor.ServiceDescriptorProto,
	method *descriptor.MethodDescriptorProto) *HTTPInfo {
	key := file.GetPackage() + service.GetName() + method.GetName()
	httpInfo, ok := t.httpInfoCache[key]
	if !ok {
		httpInfo = GetHTTPInfo(file, service, method, t.Reg)
		t.httpInfoCache[key] = httpInfo
	}
	return httpInfo
}

func ParsePermission(method *descriptor.MethodDescriptorProto) (*permission.HttpRule, error) {
	ext, err := proto.GetExtension(method.GetOptions(), permission.E_Http)
	if err!=nil{
		return nil,err
	}
	rule := ext.(*permission.HttpRule)
	return rule,nil
}

// ParseBMMethod parse BMMethodDescriptor form method descriptor proto
func ParseBMMethod(method *descriptor.MethodDescriptorProto) (*googleMethodOptionInfo, error) {
	ext, err := proto.GetExtension(method.GetOptions(), annotations.E_Http)
	if err != nil {
		return nil, fmt.Errorf("get extension error: %s", err)
	}
	rule := ext.(*annotations.HttpRule)
	var httpMethod string
	var pathPattern string
	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		pathPattern = pattern.Get
		httpMethod = http.MethodGet
	case *annotations.HttpRule_Put:
		pathPattern = pattern.Put
		httpMethod = http.MethodPut
	case *annotations.HttpRule_Post:
		pathPattern = pattern.Post
		httpMethod = http.MethodPost
	case *annotations.HttpRule_Patch:
		pathPattern = pattern.Patch
		httpMethod = http.MethodPatch
	case *annotations.HttpRule_Delete:
		pathPattern = pattern.Delete
		httpMethod = http.MethodDelete
	default:
		return nil, fmt.Errorf("unsupport http pattern %s", rule.Pattern)
	}
	bmMethod := &googleMethodOptionInfo{
		Method:      httpMethod,
		PathPattern: pathPattern,
		HTTPRule:    rule,
	}
	return bmMethod, nil
}

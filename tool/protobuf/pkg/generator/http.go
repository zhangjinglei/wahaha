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
		//定义了http扩展，但是是No
		//不生成http接口
		if parsePermission.GetMethod() == permission.HttpMethod_No {
			explicitHTTPPath = false
		}

		println("================", parsePermission)
		if parsePermission.GetMethod() != permission.HttpMethod_No &&
			parsePermission.GetPerm() == permission.Permission_NeedPerm &&
			strings.TrimSpace(parsePermission.GetPermcode()) == "" {
			panic(errors.New(service.GetName() + "." + method.GetName() + "缺少权限码定义"))
		}
		httpMethod = strings.ToUpper(parsePermission.Method.String())
	}
	//googleOptionInfo, err := ParseBMMethod(method)
	//if err!=nil{
	//	println(err.Error())
	//}
	//if err == nil {
	//	httpMethod = strings.ToUpper(googleOptionInfo.Method)
	//	p := googleOptionInfo.PathPattern
	//	if p != "" {
	//		explicitHTTPPath = true
	//		newPath = p
	//		goto END
	//	}
	//}
	//
	//if httpMethod == "" {
	//	// resolve http method
	//	httpMethod = tag.GetTagValue("method", tags)
	//	if httpMethod == "" {
	//		httpMethod = "GET"
	//	} else {
	//		httpMethod = strings.ToUpper(httpMethod)
	//	}
	//}

	newPath = "/" + file.GetPackage() + "." + service.GetName() + "/" + method.GetName()
//END:
	var p = newPath
	param := &HTTPInfo{HttpMethod: httpMethod,
		Permission:parsePermission.GetPerm(),
		PermissionCode:strings.TrimSpace(parsePermission.GetPermcode()),
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

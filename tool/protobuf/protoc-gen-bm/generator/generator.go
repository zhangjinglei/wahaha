package generator

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/generator"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/naming"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/tag"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/typemap"
	"github.com/zhangjinglei/wahaha/tool/protobuf/pkg/utils"
)

type bm struct {
	generator.Base
	filesHandled int
}

// BmGenerator BM generator.
func BmGenerator() *bm {
	t := &bm{}
	return t
}

// Generate ...
func (t *bm) Generate(in *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	t.Setup(in)

	// Showtime! Generate the response.
	resp := new(plugin.CodeGeneratorResponse)
	for _, f := range t.GenFiles {
		respFile := t.generateForFile(f)
		if respFile != nil {
			resp.File = append(resp.File, respFile)
		}
	}
	return resp
}

func (t *bm) generateForFile(file *descriptor.FileDescriptorProto) *plugin.CodeGeneratorResponse_File {
	resp := new(plugin.CodeGeneratorResponse_File)

	t.generateFileHeader(file, t.GenPkgName)
	t.generateImports(file)
	t.generatePathConstants(file)
	count := 0
	for i, service := range file.Service {
		count += t.generateBMInterface(file, service)
		t.generateBMRoute(file, service, i)
	}

	resp.Name = proto.String(naming.GenFileName(file, ".bm.go"))
	resp.Content = proto.String(t.FormattedOutput())
	t.Output.Reset()

	t.filesHandled++
	return resp
}

func (t *bm) generatePathConstants(file *descriptor.FileDescriptorProto) {
	t.P()
	for _, service := range file.Service {
		name := naming.ServiceName(service)
		for _, method := range service.Method {
			if !t.ShouldGenForMethod(file, service, method) {
				continue
			}
			apiInfo := t.GetHttpInfoCached(file, service, method)
			t.P(`var Path`, name, naming.MethodName(method), ` = "`, apiInfo.Path, `"`)
		}
		t.P()
	}
}

func (t *bm) generateFileHeader(file *descriptor.FileDescriptorProto, pkgName string) {
	t.P("// Code generated by protoc-gen-bm ", generator.Version, ", DO NOT EDIT.")
	t.P("// source: ", file.GetName())
	t.P()
	if t.filesHandled == 0 {
		comment, err := t.Reg.FileComments(file)
		if err == nil && comment.Leading != "" {
			// doc for the first file
			t.P("/*")
			t.P("Package ", t.GenPkgName, " is a generated blademaster stub package.")
			t.P("This code was generated with kratos/tool/protobuf/protoc-gen-bm ", generator.Version, ".")
			t.P()
			for _, line := range strings.Split(comment.Leading, "\n") {
				line = strings.TrimPrefix(line, " ")
				// ensure we don't escape from the block comment
				line = strings.Replace(line, "*/", "* /", -1)
				t.P(line)
			}
			t.P()
			t.P("It is generated from these files:")
			for _, f := range t.GenFiles {
				t.P("\t", f.GetName())
			}
			t.P("*/")
		}
	}
	t.P(`package `, pkgName)
	t.P()
}

func (t *bm) generateImports(file *descriptor.FileDescriptorProto) {
	//if len(file.Service) == 0 {
	//	return
	//}
	t.P(`import (`)
	//t.P(`	`,t.pkgs["context"], ` "context"`)
	t.P(`	"context"`)
	t.P()
	t.P(`	bm "github.com/zhangjinglei/wahaha/pkg/net/http/blademaster"`)
	t.P(`	"github.com/zhangjinglei/wahaha/pkg/net/http/blademaster/binding"`)

	t.P(`)`)
	// It's legal to import a message and use it as an input or output for a
	// method. Make sure to import the package of any such message. First, dedupe
	// them.
	deps := make(map[string]string) // Map of package name to quoted import path.
	deps = t.DeduceDeps(file)
	for pkg, importPath := range deps {
		t.P(`import `, pkg, ` `, importPath)
	}
	t.P()
	t.P(`// to suppressed 'imported but not used warning'`)
	t.P(`var _ *bm.Context`)
	t.P(`var _ context.Context`)
	t.P(`var _ binding.StructValidator`)

}

// Big header comments to makes it easier to visually parse a generated file.
func (t *bm) sectionComment(sectionTitle string) {
	t.P()
	t.P(`// `, strings.Repeat("=", len(sectionTitle)))
	t.P(`// `, sectionTitle)
	t.P(`// `, strings.Repeat("=", len(sectionTitle)))
	t.P()
}

func (t *bm) generateBMRoute(
	file *descriptor.FileDescriptorProto,
	service *descriptor.ServiceDescriptorProto,
	index int) {
	// old mode is generate xx.route.go in the http pkg
	// new mode is generate route code in the same .bm.go
	// route rule /x{department}/{project-name}/{path_prefix}/method_name
	// generate each route method
	servName := naming.ServiceName(service)
	versionPrefix := naming.GetVersionPrefix(t.GenPkgName)
	svcName := utils.LcFirst(utils.CamelCase(versionPrefix)) + servName + "Svc"
	t.P(`var `, svcName, ` `, servName, `BMServer`)

	type methodInfo struct {
		midwares      []string
		routeFuncName string
		apiInfo       *generator.HTTPInfo
		methodName    string
	}
	var methList []methodInfo
	var allMidwareMap = make(map[string]bool)
	//var isLegacyPkg = false
	for _, method := range service.Method {
		if !t.ShouldGenForMethod(file, service, method) {
			continue
		}
		var midwares []string
		comments, _ := t.Reg.MethodComments(file, service, method)
		tags := tag.GetTagsInComment(comments.Leading)
		if tag.GetTagValue("dynamic", tags) == "true" {
			continue
		}
		apiInfo := t.GetHttpInfoCached(file, service, method)
		//isLegacyPkg = apiInfo.IsLegacyPath
		//httpMethod, legacyPath, path := getHttpInfo(file, service, method, t.reg)
		//if legacyPath != "" {
		//	isLegacyPkg = true
		//}

		midStr := tag.GetTagValue("midware", tags)
		if midStr != "" {
			midwares = strings.Split(midStr, ",")
			for _, m := range midwares {
				allMidwareMap[m] = true
			}
		}

		methName := naming.MethodName(method)
		inputType := t.GoTypeName(method.GetInputType())

		routeName := utils.LcFirst(utils.CamelCase(servName) +
			utils.CamelCase(methName))

		methList = append(methList, methodInfo{
			apiInfo:       apiInfo,
			midwares:      midwares,
			routeFuncName: routeName,
			methodName:    method.GetName(),
		})

		t.P(fmt.Sprintf("func %s (c *bm.Context) {", routeName))
		t.P(`	p := new(`, inputType, `)`)
		requestBinding := ""
		if t.hasHeaderTag(t.Reg.MessageDefinition(method.GetInputType())) {
			requestBinding = ", binding.Request"
		}
		t.P(`	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))` +
			requestBinding + `); err != nil {`)
		t.P(`		return`)
		t.P(`	}`)
		t.P(`	resp, err := `, svcName, `.`, methName, `(c, p)`)
		t.P(`	c.JSON(resp, err)`)
		t.P(`}`)
		t.P(``)
	}

	// generate route group
	var midList []string
	for m := range allMidwareMap {
		midList = append(midList, m+" bm.HandlerFunc")
	}

	sort.Strings(midList)

	//// 注册老的路由的方法
	//if isLegacyPkg {
	//	funcName := `Register` + utils.CamelCase(versionPrefix) + servName + `Service`
	//	t.P(`// `, funcName, ` Register the blademaster route with middleware map`)
	//	t.P(`// midMap is the middleware map, the key is defined in proto`)
	//	t.P(`func `, funcName, `(e *bm.Engine, svc `, servName, "BMServer, midMap map[string]bm.HandlerFunc)", ` {`)
	//	var keys []string
	//	for m := range allMidwareMap {
	//		keys = append(keys, m)
	//	}
	//	// to keep generated code consistent
	//	sort.Strings(keys)
	//	for _, m := range keys {
	//		t.P(m, ` := midMap["`, m, `"]`)
	//	}
	//
	//	t.P(svcName, ` = svc`)
	//	for _, methInfo := range methList {
	//		var midArgStr string
	//		if len(methInfo.midwares) == 0 {
	//			midArgStr = ""
	//		} else {
	//			midArgStr = strings.Join(methInfo.midwares, ", ") + ", "
	//		}
	//		t.P(`e.`, methInfo.apiInfo.HttpMethod, `("`, methInfo.apiInfo.LegacyPath, `", `, midArgStr, methInfo.routeFuncName, `)`)
	//	}
	//	t.P(`	}`)
	//} else {
	// 新的注册路由的方法
	var bmFuncName = fmt.Sprintf("Register%sBMServer", servName)
	t.P(`// `, bmFuncName, ` Register the blademaster route`)
	t.P(`func `, bmFuncName, `(e *bm.Engine, server `, servName, `BMServer) {`)
	t.P(svcName, ` = server`)
	for _, methInfo := range methList {
		//comments, _ := t.Reg.MethodComments(file, service, methInfo)
		//tags := tag.GetTagsInComment(comments.Leading)
		//if tag.GetTagValue("dynamic", tags) == "true" {
		//	continue
		//}
		//t.P("//zhangjinglei")
		//t.P(`permission["`+methInfo.apiInfo.NewPath+`"]=[]string{"a","b","c"}`)
		//if methInfo.apiInfo.Permission!=permission.Permission_IgnoreLogin{
		//	t.P(`e.`, methInfo.apiInfo.HttpMethod, `("`, methInfo.apiInfo.NewPath, `",e.AuthMid("`+methInfo.apiInfo.App+`","`+methInfo.apiInfo.NewPath+`"),`, methInfo.routeFuncName, ` )`)
		//}else {
		//	t.P(`e.`, methInfo.apiInfo.HttpMethod, `("`, methInfo.apiInfo.NewPath, `",`, methInfo.routeFuncName, ` )`)
		//}

		t.P(`e.`, methInfo.apiInfo.HttpMethod, `("`, methInfo.apiInfo.NewPath, `",`, methInfo.routeFuncName, ` )`)

	}
	for _, methInfo := range methList {
		description := strings.ReplaceAll(methInfo.apiInfo.Description, `"`, "")
		description = strings.ReplaceAll(description, "\r", "")
		description = strings.ReplaceAll(description, "\n", "")
		svcdescription := strings.ReplaceAll(methInfo.apiInfo.ServiceDescription, `"`, "")
		svcdescription = strings.ReplaceAll(svcdescription, "\r", "")
		svcdescription = strings.ReplaceAll(svcdescription, "\n", "")
		t.P(`e.AuthMid("`+methInfo.apiInfo.App+`","`, methInfo.apiInfo.NewPath, `",`, strconv.Itoa(int(methInfo.apiInfo.Permission)), `,"`+description+`"`, `,"`+methInfo.apiInfo.PackageServiceName+`"`, `,"`+svcdescription+`"`, ` )`)
	}
	t.P(`	}`)
	//}
}

func (t *bm) hasHeaderTag(md *typemap.MessageDefinition) bool {
	if md.Descriptor.Field == nil {
		return false
	}
	for _, f := range md.Descriptor.Field {
		t := tag.GetMoreTags(f)
		if t != nil {
			st := reflect.StructTag(*t)
			if st.Get("request") != "" {
				return true
			}
			if st.Get("header") != "" {
				return true
			}
		}
	}
	return false
}

func (t *bm) generateBMInterface(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto) int {
	count := 0
	servName := naming.ServiceName(service)
	t.P("// " + servName + "BMServer is the server API for " + servName + " service.")

	comments, err := t.Reg.ServiceComments(file, service)
	if err == nil {
		t.PrintComments(comments)
	}
	t.P(`type `, servName, `BMServer interface {`)
	for _, method := range service.Method {

		if !t.ShouldGenForMethod(file, service, method) {
			continue
		}
		count++
		t.generateInterfaceMethod(file, service, method, comments)
		t.P()
	}
	t.P(`}`)
	return count
}

func (t *bm) generateInterfaceMethod(file *descriptor.FileDescriptorProto,
	service *descriptor.ServiceDescriptorProto,
	method *descriptor.MethodDescriptorProto,
	comments typemap.DefinitionComments) {
	comments, err := t.Reg.MethodComments(file, service, method)

	methName := naming.MethodName(method)
	outputType := t.GoTypeName(method.GetOutputType())
	inputType := t.GoTypeName(method.GetInputType())
	tags := tag.GetTagsInComment(comments.Leading)
	if tag.GetTagValue("dynamic", tags) == "true" {
		return
	}

	if err == nil {
		t.PrintComments(comments)
	}

	respDynamic := tag.GetTagValue("dynamic_resp", tags) == "true"
	if respDynamic {
		t.P(fmt.Sprintf(`	%s(ctx context.Context, req *%s) (resp interface{}, err error)`,
			methName, inputType))
	} else {
		t.P(fmt.Sprintf(`	%s(ctx context.Context, req *%s) (resp *%s, err error)`,
			methName, inputType, outputType))
	}
}

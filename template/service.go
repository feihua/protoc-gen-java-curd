package template

import (
	"github/feihua/protoc-gen-java-curd/util"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateServiceFile(gen *protogen.Plugin, file *protogen.File, t string) {
	for _, service := range file.Services {
		filename := "./generate/service/" + service.GoName + "Service.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		g.P("package ", file.Desc.Package(), ".service;")
		g.P() // 换行
		g.P("import java.util.List;\n")
		methodParams := make(map[string]string)
		for _, method := range service.Methods {
			methodParams[method.Input.GoIdent.GoName] = "0"
			methodParams[method.Output.GoIdent.GoName] = "1"
		}
		for param := range methodParams {
			g.P("import ", file.Desc.Package(), ".vo.", param, ";")
		}
		g.P()
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("/**")
		g.P(" * 描述: ", serviceComment)
		g.P(" * 作者：", "demo")
		g.P(" * 日期：", t)
		g.P(" */")
		g.P("public interface ", service.GoName, "Service {\n")
		for _, method := range service.Methods {
			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			g.P("\t/**")
			g.P("\t * ", methodComment)
			g.P("\t * ")
			g.P("\t * @param ", util.FirstLower(method.Input.GoIdent.GoName), " 请求参数")
			g.P("\t * @return ", method.Output.GoIdent)
			g.P("\t * @author ", "demo")
			g.P("\t * @date ", t)
			g.P("\t */")
			g.P("\t", method.Output.GoIdent, " ", util.FirstLower(method.GoName), "(", method.Input.GoIdent, " ", util.FirstLower(method.Input.GoIdent.GoName), ");\n")
		}
		g.P("}")

	}

}

func GenerateServiceImplFile(gen *protogen.Plugin, file *protogen.File, t string) {

	for _, service := range file.Services {
		filename := "./generate/service/impl/" + service.GoName + "ServiceImpl.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		g.P("package ", file.Desc.Package(), ".service.impl;")
		g.P() // 换行
		g.P("import java.util.List;\n")
		g.P("import java.util.stream.Collectors;\n")
		g.P("import org.springframework.beans.BeanUtils;")
		g.P("import org.springframework.beans.factory.annotation.Autowired;")
		g.P("import org.springframework.stereotype.Service;\n")
		g.P("import ", file.Desc.Package(), ".biz.", service.GoName, "Biz;\n")
		g.P("import ", file.Desc.Package(), ".service.", service.GoName, "Service;\n")
		methodParams := make(map[string]string)
		for _, method := range service.Methods {
			methodParams[method.Input.GoIdent.GoName] = "0"
			methodParams[method.Output.GoIdent.GoName] = "1"
		}
		for param := range methodParams {
			g.P("import ", file.Desc.Package(), ".entity.", param, ";")
		}
		for param := range methodParams {
			g.P("import ", file.Desc.Package(), ".vo.", param, ";")
		}
		g.P()
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("/**")
		g.P(" * 描述: ", serviceComment)
		g.P(" * 作者：", "demo")
		g.P(" * 日期：", t)
		g.P(" */")
		g.P("@Service")
		g.P("public class ", service.GoName, "ServiceImpl implements ", service.GoName, "Service {\n")
		g.P("\t@Autowired")
		g.P("\tprivate ", service.GoName, "Biz ", util.FirstLower(service.GoName), "Biz;\n")
		for _, method := range service.Methods {

			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			g.P("\t/**")
			g.P("\t * ", methodComment)
			g.P("\t * ")
			g.P("\t * @param ", util.FirstLower(method.Input.GoIdent.GoName), " 请求参数")
			g.P("\t * @return ", method.Output.GoIdent)
			g.P("\t * @author ", "demo")
			g.P("\t * @date ", t)
			g.P("\t */")
			g.P("\t@Override")
			g.P("\tpublic ", method.Output.GoIdent, " ", util.FirstLower(method.GoName), "(", method.Input.GoIdent, " ", util.FirstLower(method.Input.GoIdent.GoName), ")", " {")
			//g.P("		return ", FirstLower(service.GoName), "Biz.", FirstLower(method.GoName), "(", FirstLower(method.Input.GoIdent.GoName), ");")
			//HelloReply helloReply = greeterTestDao.sayHelloAgain(HelloRequest.builder().build());
			//
			g.P("\t\t", method.Output.GoIdent, " ", util.FirstLower(method.Output.GoIdent.GoName), " = ", util.FirstLower(service.GoName), "Biz.", util.FirstLower(method.GoName), "(", method.Input.GoIdent.GoName, ".builder().build());")
			//return HelloReplyVo.builder().build();
			g.P("\t\treturn ", method.Output.GoIdent, "builder().build();")
			g.P("	}\n")
		}
		g.P("}")

	}

}

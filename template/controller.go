package template

import (
	"github/feihua/protoc-gen-java-curd/util"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateControllerFile(gen *protogen.Plugin, file *protogen.File, t string) {
	//循环解析service
	for _, service := range file.Services {
		filename := "./generate/controller/v2/" + service.GoName + "Controller.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)

		//文件头导入模板代码
		g.P("package ", file.Desc.Package(), ".controller;\n")
		g.P("import io.swagger.annotations.Api;")
		g.P("import io.swagger.annotations.ApiOperation;\n")
		g.P("import java.util.List;\n")
		g.P("import javax.validation.Valid;\n")
		g.P("import org.springframework.beans.factory.annotation.Autowired;")
		g.P("import org.springframework.web.bind.annotation.PostMapping;")
		g.P("import org.springframework.web.bind.annotation.RequestBody;")
		g.P("import org.springframework.web.bind.annotation.RequestMapping;")
		g.P("import org.springframework.web.bind.annotation.RestController;\n")
		g.P("import ", file.Desc.Package(), ".service.", service.GoName, "Service;\n")
		methodParams := make(map[string]string)
		for _, method := range service.Methods {
			methodParams[method.Input.GoIdent.GoName] = "0"
			methodParams[method.Output.GoIdent.GoName] = "1"
		}
		for param := range methodParams {
			g.P("import ", file.Desc.Package(), ".vo.", param, ";")
		}
		g.P()

		//获取service注释
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("/**")
		g.P(" * 描述: ", serviceComment)
		g.P(" * 作者：", "demo")
		g.P(" * 日期：", t)
		g.P(" */")
		g.P("@Api(tags = \"", serviceComment, "\")")
		g.P("@RestController")
		g.P("@RequestMapping(\"/", util.FirstLower(service.GoName), "\")")
		g.P("public class ", service.GoName, "Controller {\n")
		g.P("\t@Autowired")
		g.P("\tprivate ", service.GoName, "Service ", util.FirstLower(service.GoName), "Service;\n")
		for _, method := range service.Methods {

			//获取method注释
			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			//   /**
			//    * 添加支付宝支付配置
			//    *
			//    * @param record 请求参数
			//    * @return Result<Integer>
			//    * @author 刘飞华
			//    * @date: 2024-09-23 09:54:03
			//    */
			g.P("\t/**")
			g.P("\t * ", methodComment)
			g.P("\t * ")
			g.P("\t * @param ", util.FirstLower(method.Input.GoIdent.GoName), " 请求参数")
			g.P("\t * @return ", method.Output.GoIdent)
			g.P("\t * @author ", "demo")
			g.P("\t * @date ", t)
			g.P("\t */")
			g.P("\t@ApiOperation(\"", methodComment, "\")")
			g.P("\t@PostMapping(\"/", util.FirstLower(method.GoName), "\")")
			g.P("\tpublic ", method.Output.GoIdent, " ", util.FirstLower(method.GoName), "(@RequestBody @Valid ", method.Input.GoIdent, " ", util.FirstLower(method.Input.GoIdent.GoName), ")", " {")
			g.P("\t\treturn ", util.FirstLower(service.GoName), "Service.", util.FirstLower(method.GoName), "(", util.FirstLower(method.Input.GoIdent.GoName), ");")
			g.P("\t}\n")
		}
		g.P("}")

	}

}

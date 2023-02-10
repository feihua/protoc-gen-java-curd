package template

import (
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func GenerateControllerFile(gen *protogen.Plugin, file *protogen.File) {
	for _, service := range file.Services {
		filename := "./generate/controller/" + service.GoName + "Controller.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".controller;")
		g.P() // 换行
		g.P("import io.swagger.annotations.Api;")
		g.P("import io.swagger.annotations.ApiOperation;\n")
		g.P("import java.util.List;\n")
		g.P("import javax.validation.Valid;\n")
		g.P("import org.springframework.beans.factory.annotation.Autowired;")
		g.P("import org.springframework.web.bind.annotation.PostMapping;")
		g.P("import org.springframework.web.bind.annotation.RequestBody;")
		g.P("import org.springframework.web.bind.annotation.RequestMapping;")
		g.P("import org.springframework.web.bind.annotation.RestController;\n")
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("@Api(tags = \"", serviceComment, "\")")
		g.P("@RestController")
		g.P("@RequestMapping(\"/", FirstLower(service.GoName), "\")")
		g.P("public class ", service.GoName, "Controller {\n")
		g.P("	@Autowired")
		g.P("	private ", service.GoName, "Service service\n")
		for _, method := range service.Methods {

			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			//g.P(methodComment)
			g.P("	@ApiOperation(\"", methodComment, "\")")
			g.P("	@PostMapping(\"/", FirstLower(method.GoName), "\")")
			g.P("	public ", method.Output.GoIdent, " ", FirstLower(method.GoName), "(@RequestBody @Valid ", method.Input.GoIdent, " ", FirstLower(method.Input.GoIdent.GoName), ")", " {")
			g.P("		return service.", FirstLower(method.GoName), "(", FirstLower(method.Input.GoIdent.GoName), ")")
			g.P("	}\n")
		}
		// 输出 }
		g.P("}")

	}

}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

package template

import (
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func GenerateControllerFile(gen *protogen.Plugin, file *protogen.File) {
	for _, service := range file.Services {
		filename := "./generate/controller/" + service.GoName + "Controller.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
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
			g.P("import ", file.Desc.Package(), ".vo.", param, "Vo;")
		}
		g.P()
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("@Api(tags = \"", serviceComment, "\")")
		g.P("@RestController")
		g.P("@RequestMapping(\"/", FirstLower(service.GoName), "\")")
		g.P("public class ", service.GoName, "Controller {\n")
		g.P("\t@Autowired")
		g.P("\tprivate ", service.GoName, "Service ", FirstLower(service.GoName), "Service;\n")
		for _, method := range service.Methods {

			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			g.P("\t@ApiOperation(\"", methodComment, "\")")
			g.P("\t@PostMapping(\"/", FirstLower(method.GoName), "\")")
			g.P("\tpublic ", method.Output.GoIdent, "Vo ", FirstLower(method.GoName), "(@RequestBody @Valid ", method.Input.GoIdent, "Vo ", FirstLower(method.Input.GoIdent.GoName), ")", " {")
			g.P("\t\treturn ", FirstLower(service.GoName), "Service.", FirstLower(method.GoName), "(", FirstLower(method.Input.GoIdent.GoName), ");")
			g.P("\t}\n")
		}
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

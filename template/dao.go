package template

import "google.golang.org/protobuf/compiler/protogen"

func GenerateDaoFile(gen *protogen.Plugin, file *protogen.File) {
	for _, service := range file.Services {
		filename := "./generate/dao/" + service.GoName + "Dao.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		g.P("package ", file.Desc.Package(), ".dao;\n")
		g.P("import org.apache.ibatis.annotations.Mapper;\n")
		g.P("import java.util.List;\n")
		methodParams := make(map[string]string)
		for _, method := range service.Methods {
			methodParams[method.Input.GoIdent.GoName] = "0"
			methodParams[method.Output.GoIdent.GoName] = "1"
		}
		for param := range methodParams {
			g.P("import ", file.Desc.Package(), ".entity.", param, ";")
		}
		g.P()
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("@Mapper")
		g.P("public interface ", service.GoName, "Dao {\n")
		for _, method := range service.Methods {
			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[0 : len(methodComment)-2]
			g.P("\t", methodComment)
			g.P("\t", method.Output.GoIdent, " ", FirstLower(method.GoName), "(", method.Input.GoIdent, " ", FirstLower(method.Input.GoIdent.GoName), ");\n")
		}
		g.P("}")

	}

}

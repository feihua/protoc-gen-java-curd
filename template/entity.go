package template

import "google.golang.org/protobuf/compiler/protogen"

func GenerateEntityFile(gen *protogen.Plugin, file *protogen.File) {
	for _, m := range file.Messages {
		filename := "./generate/entity/" + m.GoIdent.GoName + ".java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".entity;")
		g.P() // 换行
		// 输出 type m.GoIdent struct {
		g.P("import java.io.Serializable;")
		g.P("import java.util.Date;\n")
		g.P("import lombok.AllArgsConstructor;")
		g.P("import lombok.Builder;")
		g.P("import lombok.Data;")
		g.P("import lombok.NoArgsConstructor;\n")
		g.P("@Data")
		g.P("@Builder")
		g.P("@NoArgsConstructor")
		g.P("@AllArgsConstructor")
		g.P("public class ", m.GoIdent, " implements Serializable {\n")
		for _, field := range m.Fields {
			//leadingComment := field.Comments.Leading.String()
			trailingComment := field.Comments.Trailing.String()
			trailingComment = trailingComment[0 : len(trailingComment)-1]
			g.P("	", trailingComment)
			g.P("	private ", field.Desc.Kind(), " ", field.Desc.JSONName(), ";")
			// 输出 行首注释
			//g.P(leadingComment)
			// 输出 行内容
			//g.P(line)
		}
		// 输出 }
		g.P("}")
	}

	//file.Services[0].Methods[0].
	//g.P() // 换行
}

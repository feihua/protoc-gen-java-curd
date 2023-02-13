package template

import "google.golang.org/protobuf/compiler/protogen"

func GenerateVoFile(gen *protogen.Plugin, file *protogen.File) {
	for _, m := range file.Messages {
		filename := "./generate/vo/" + m.GoIdent.GoName + "Vo.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".vo;")
		g.P() // 换行

		// 输出 type m.GoIdent struct {
		g.P("import io.swagger.annotations.ApiModel;")
		g.P("import io.swagger.annotations.ApiModelProperty;\n")
		g.P("import java.io.Serializable;")
		g.P("import java.util.Date;\n")
		g.P("import lombok.AllArgsConstructor;")
		g.P("import lombok.Builder;")
		g.P("import lombok.Data;")
		g.P("import lombok.NoArgsConstructor;\n")
		serviceComment := m.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("@ApiModel(\"", serviceComment, "\")")
		g.P("@Data")
		g.P("@Builder")
		g.P("@NoArgsConstructor")
		g.P("@AllArgsConstructor")
		g.P("public class ", m.GoIdent, "Vo implements Serializable {\n")
		for _, field := range m.Fields {
			//leadingComment := field.Comments.Leading.String()
			trailingComment := field.Comments.Trailing.String()
			trailingComment = trailingComment[3 : len(trailingComment)-1]
			g.P("\t@ApiModelProperty(\"", trailingComment, "\")")
			g.P("\tprivate ", ProtoTypeToJavaType[field.Desc.Kind().String()], " ", field.Desc.JSONName(), ";\n")
		}
		g.P("}")
	}

}

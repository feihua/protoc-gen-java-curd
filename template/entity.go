package template

import "google.golang.org/protobuf/compiler/protogen"

func GenerateEntityFile(gen *protogen.Plugin, file *protogen.File, t string) {
	for _, m := range file.Messages {
		filename := "./generate/entity/" + m.GoIdent.GoName + ".java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
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
		serviceComment := m.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("/**")
		g.P(" * 描述: ", serviceComment)
		g.P(" * 作者：", "demo")
		g.P(" * 日期：", t)
		g.P(" */")
		g.P("public class ", m.GoIdent, " implements Serializable {\n")
		for _, field := range m.Fields {
			trailingComment := field.Comments.Trailing.String()
			trailingComment = trailingComment[0 : len(trailingComment)-1]
			g.P("\t", trailingComment)
			g.P("\tprivate ", ProtoTypeToJavaType[field.Desc.Kind().String()], " ", field.Desc.JSONName(), ";")
		}
		g.P("}")
	}

}

var ProtoTypeToJavaType = map[string]string{
	"double": "double",
	"float":  "float",
	"int64":  "long",
	"int32":  "int",
	"uint64": "long",
	"uint32": "int",
	"sint64": "long",
	"sint32": "int",
	"bool":   "boolean",
	"string": "String",
}

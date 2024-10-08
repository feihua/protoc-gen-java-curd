package template

import (
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func GenerateVoV3File(gen *protogen.Plugin, file *protogen.File, t string) {
	for _, m := range file.Messages {
		filename := "./generate/vo/v3/req/" + m.GoIdent.GoName + ".java"
		if strings.Contains(m.GoIdent.GoName, "Resp") {
			filename = "./generate/vo/v3/resp/" + m.GoIdent.GoName + ".java"
		}

		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".vo;")
		g.P() // 换行

		// 输出 type m.GoIdent struct {
		g.P("import io.swagger.v3.oas.annotations.media.Schema;\n")
		g.P("import java.io.Serializable;")
		g.P("import java.util.Date;\n")
		g.P("import lombok.AllArgsConstructor;")
		g.P("import lombok.Builder;")
		g.P("import lombok.Data;")
		g.P("import lombok.NoArgsConstructor;\n")
		serviceComment := m.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("/**")
		g.P(" * 描述: ", serviceComment)
		g.P(" * 作者：", "demo")
		g.P(" * 日期：", t)
		g.P(" */")
		g.P("@Schema(description = \"", serviceComment, "\"")
		g.P("@Data")
		g.P("@Builder")
		g.P("@NoArgsConstructor")
		g.P("@AllArgsConstructor")
		g.P("public class ", m.GoIdent, " implements Serializable {\n")
		g.P("\tprivate static final long serialVersionUID = 1L;\n")
		for _, field := range m.Fields {
			//leadingComment := field.Comments.Leading.String()
			trailingComment := field.Comments.Trailing.String()
			trailingComment = trailingComment[3 : len(trailingComment)-1]
			if strings.Contains(m.GoIdent.GoName, "Req") {
				g.P("\t@Schema(description = \"", trailingComment, "\", requiredMode = Schema.RequiredMode.REQUIRED)")
			} else {
				g.P("\t@Schema(description = \"", trailingComment, "\")")
			}

			g.P("\tprivate ", ProtoTypeToJavaType[field.Desc.Kind().String()], " ", field.Desc.JSONName(), ";\n")
		}
		g.P("}")
	}

}

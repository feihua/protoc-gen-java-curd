package template

import "google.golang.org/protobuf/compiler/protogen"

func GenerateXmlFile(gen *protogen.Plugin, file *protogen.File) {

	for _, service := range file.Services {
		filename := "./generate/mapper/" + service.GoName + "Mapper.xml"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		g.P("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
		g.P("<!DOCTYPE mapper PUBLIC \"-//mybatis.org//DTD Mapper 3.0//EN\" \"https://mybatis.org/dtd/mybatis-3-mapper.dtd\">")
		g.P("<mapper namespace=\"com.test.dao.xx\">")
		g.P("\t<resultMap id=\"BaseResultMap\" type=\"com.test.entity.xx\">\n        <!-- <result column=\"id\" property=\"id\" jdbcType=\"BIGINT\"/> -->\n    </resultMap>\n\n    <sql id=\"Base_Column_List\">\n        \n    </sql>\n")
		for _, method := range service.Methods {
			//<select id="query" parameterType="com.test.entity.xx" resultMap="BaseResultMap">
			//
			serviceComment := method.Comments.Leading.String()
			serviceComment = serviceComment[3 : len(serviceComment)-2]
			g.P("    <!-- ", serviceComment, " -->")
			g.P("    <select id=\"", FirstLower(method.GoName), "\" parameterType=\"com.test.entity.xx\" resultMap=\"BaseResultMap\">")
			g.P()
			g.P("    </select>\n")
		}
		g.P("</mapper>")
	}
}

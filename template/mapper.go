package template

import (
	"github/feihua/protoc-gen-java-curd/util"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func GenerateXmlFile(gen *protogen.Plugin, file *protogen.File, t string) {

	for _, service := range file.Services {
		filename := "./generate/mapper/" + service.GoName + "Mapper.xml"
		PackageName := file.Desc.Package()
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		g.P("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
		g.P("<!DOCTYPE mapper PUBLIC \"-//mybatis.org//DTD Mapper 3.0//EN\" \"https://mybatis.org/dtd/mybatis-3-mapper.dtd\">")
		g.P("<mapper namespace=\"", PackageName, ".dao.", service.GoName, "Dao\">")
		//g.P("\t<resultMap id=\"BaseResultMap\" type=\"com.test.entity.xx\">\n        <!-- <result column=\"id\" property=\"id\" jdbcType=\"BIGINT\"/> -->\n    </resultMap>\n\n    <sql id=\"Base_Column_List\">\n        \n    </sql>\n")
		for _, method := range service.Methods {
			serviceComment := method.Comments.Leading.String()
			serviceComment = serviceComment[3 : len(serviceComment)-2]
			g.P("\t<!-- ", serviceComment, " -->")
			methodName := util.FirstLower(method.GoName)
			inputParam := method.Input.GoIdent.GoName
			outParam := method.Output.GoIdent.GoName

			if strings.HasPrefix(methodName, "insert") || strings.HasPrefix(methodName, "save") || strings.HasPrefix(methodName, "add") {
				g.P("\t<insert id=\"", methodName, "\" parameterType=\"", PackageName, ".entity.", inputParam, "\">\n")
				g.P("\t</insert>\n")
			} else if strings.HasPrefix(methodName, "update") || strings.HasPrefix(methodName, "upd") {
				g.P("\t<update id=\"", methodName, "\" parameterType=\"", PackageName, ".entity.", inputParam, "\">\n")
				g.P("\t</update>\n")
			} else if strings.HasPrefix(methodName, "del") || strings.HasPrefix(methodName, "delete") || strings.HasPrefix(methodName, "remove") {
				g.P("\t<delete id=\"", methodName, "\" parameterType=\"", PackageName, ".entity.", inputParam, "\">\n")
				g.P("\t</delete>\n")
			} else if strings.HasPrefix(methodName, "query") || strings.HasPrefix(methodName, "get") || strings.HasPrefix(methodName, "find") {
				g.P("\t<select id=\"", methodName, "\" parameterType=\"", PackageName, ".entity.", inputParam, "\" resultType=\"", PackageName, ".entity.", outParam, "\">\n")
				g.P("\t</select>\n")
			} else {
				g.P("\t<select id=\"", methodName, "\" parameterType=\"", PackageName, ".entity.", inputParam, "\" resultType=\"", PackageName, ".entity.", outParam, "\">\n")
				g.P("\t</select>\n")
			}
		}
		g.P("</mapper>")
	}
}

package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			//generateFile(gen, f)
			generateEntityFile(gen, f)
			generateDaoFile(gen, f)
			generateXmlFile(gen, f)
			generateServiceFile(gen, f)
			generateServiceImplFile(gen, f)
			generateControllerFile(gen, f)
			generateVoFile(gen, f)
		}
		return nil
	})
}

// 生成.struct.go文件，参数为 输出插件gen，以及读取的文件file
func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + ".struct.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	// 输出 package packageName
	g.P("package ", file.GoPackageName)
	g.P() // 换行

	for _, m := range file.Messages {
		// 输出 type m.GoIdent struct {
		g.P("type ", m.GoIdent, " struct {")
		for _, field := range m.Fields {
			//leadingComment := field.Comments.Leading.String()
			trailingComment := field.Comments.Trailing.String()
			line := fmt.Sprintf("%s %s `json:\"%s\"` %s", field.GoName, field.Desc.Kind(), field.Desc.JSONName(), trailingComment)
			// 输出 行首注释
			//g.P(leadingComment)
			// 输出 行内容
			g.P(line)
		}
		// 输出 }
		g.P("}")
	}

	//file.Services[0].Methods[0].
	g.P() // 换行
	g.Content()
}

func generateEntityFile(gen *protogen.Plugin, file *protogen.File) {
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

func generateDaoFile(gen *protogen.Plugin, file *protogen.File) {
	for _, service := range file.Services {
		filename := "./generate/dao/" + service.GoName + "Dao.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".dao;")
		g.P() // 换
		g.P("import java.util.List;\n")
		g.P("import java.util.stream.Collectors;\n")
		//g.P("import org.springframework.beans.BeanUtils;")
		//g.P("import org.springframework.beans.factory.annotation.Autowired;")
		//g.P("import org.springframework.stereotype.Service;\n")
		//g.P("import org.springframework.web.bind.annotation.RequestMapping;")
		//g.P("import org.springframework.web.bind.annotation.RestController;\n")
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		//g.P("@Service")
		g.P("public class ", service.GoName, "Dao {\n")
		//g.P("	@Autowired")
		//g.P("	private ", service.GoName, "Dao ", FirstLower(service.GoName), "Dao\n")
		for _, method := range service.Methods {

			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			//g.P(methodComment)
			//g.P("	@ApiOperation(\"", methodComment, "\")")
			//g.P("	@PostMapping(\"/", FirstLower(method.GoName), "\")")
			//g.P("	@Override")
			g.P("	public ", method.Output.GoIdent, " ", FirstLower(method.GoName), "(", method.Input.GoIdent, " ", FirstLower(method.Input.GoIdent.GoName), ")\n")
			//g.P("	}\n")
		}
		// 输出 }
		g.P("}")

	}

	//g.P() // 换行
}

func generateXmlFile(gen *protogen.Plugin, file *protogen.File) {

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

func generateServiceFile(gen *protogen.Plugin, file *protogen.File) {
	for _, service := range file.Services {
		filename := "./generate/service/" + service.GoName + "Service.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".service;")
		g.P() // 换行
		g.P("import java.util.List;\n")
		g.P("import java.util.stream.Collectors;\n")
		//g.P("import org.springframework.beans.BeanUtils;")
		//g.P("import org.springframework.beans.factory.annotation.Autowired;")
		//g.P("import org.springframework.stereotype.Service;\n")
		//g.P("import org.springframework.web.bind.annotation.RequestMapping;")
		//g.P("import org.springframework.web.bind.annotation.RestController;\n")
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		//g.P("@Service")
		g.P("public class ", service.GoName, "Service {\n")
		//g.P("	@Autowired")
		//g.P("	private ", service.GoName, "Dao ", FirstLower(service.GoName), "Dao\n")
		for _, method := range service.Methods {

			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			//g.P(methodComment)
			//g.P("	@ApiOperation(\"", methodComment, "\")")
			//g.P("	@PostMapping(\"/", FirstLower(method.GoName), "\")")
			//g.P("	@Override")
			g.P("	public ", method.Output.GoIdent, " ", FirstLower(method.GoName), "(", method.Input.GoIdent, " ", FirstLower(method.Input.GoIdent.GoName), ")\n")
			//g.P("	}\n")
		}
		// 输出 }
		g.P("}")

	}

}

func generateServiceImplFile(gen *protogen.Plugin, file *protogen.File) {

	for _, service := range file.Services {
		filename := "./generate/service/impl/" + service.GoName + "ServiceImpl.java"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		// 输出 package packageName
		g.P("package ", file.Desc.Package(), ".service.impl;")
		g.P() // 换行
		g.P("import java.util.List;\n")
		g.P("import java.util.stream.Collectors;\n")
		g.P("import org.springframework.beans.BeanUtils;")
		g.P("import org.springframework.beans.factory.annotation.Autowired;")
		g.P("import org.springframework.stereotype.Service;\n")
		//g.P("import org.springframework.web.bind.annotation.RequestMapping;")
		//g.P("import org.springframework.web.bind.annotation.RestController;\n")
		serviceComment := service.Comments.Leading.String()
		serviceComment = serviceComment[3 : len(serviceComment)-2]
		g.P("@Service")
		g.P("public class ", service.GoName, "ServiceImpl implements ", service.GoName, "Service {\n")
		g.P("	@Autowired")
		g.P("	private ", service.GoName, "Dao ", FirstLower(service.GoName), "Dao\n")
		for _, method := range service.Methods {

			methodComment := method.Comments.Leading.String()
			methodComment = methodComment[3 : len(methodComment)-2]
			//g.P(methodComment)
			//g.P("	@ApiOperation(\"", methodComment, "\")")
			//g.P("	@PostMapping(\"/", FirstLower(method.GoName), "\")")
			g.P("	@Override")
			g.P("	public ", method.Output.GoIdent, " ", FirstLower(method.GoName), "(", method.Input.GoIdent, " ", FirstLower(method.Input.GoIdent.GoName), ")", " {")
			g.P("		return ", FirstLower(service.GoName), "Dao.", FirstLower(method.GoName), "(", FirstLower(method.Input.GoIdent.GoName), ")")
			g.P("	}\n")
		}
		// 输出 }
		g.P("}")

	}

}

func generateControllerFile(gen *protogen.Plugin, file *protogen.File) {
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

func generateVoFile(gen *protogen.Plugin, file *protogen.File) {
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
		g.P("@Data")
		g.P("@Builder")
		g.P("@NoArgsConstructor")
		g.P("@AllArgsConstructor")
		g.P("public class ", m.GoIdent, " implements Serializable {\n")
		for _, field := range m.Fields {
			//leadingComment := field.Comments.Leading.String()
			trailingComment := field.Comments.Trailing.String()
			trailingComment = trailingComment[3 : len(trailingComment)-1]
			g.P("	@ApiModelProperty(\"", trailingComment, "\")")
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

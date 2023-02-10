package template

import "google.golang.org/protobuf/compiler/protogen"

func GenerateServiceFile(gen *protogen.Plugin, file *protogen.File) {
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

func GenerateServiceImplFile(gen *protogen.Plugin, file *protogen.File) {

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

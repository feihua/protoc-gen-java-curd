package main

import (
	"github/feihua/protoc-gen-java-curd/template"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		// 这个循环遍历所有要生成的proto文件
		for _, f := range gen.Files {
			//如果该文件不需要生成,则跳过
			if !f.Generate {
				continue
			}
			//如果需要生成，就把文件的相关信息传递给生成器
			generateFile(gen, f)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, f *protogen.File) {
	template.GenerateEntityFile(gen, f)
	template.GenerateDaoFile(gen, f)
	template.GenerateXmlFile(gen, f)
	template.GenerateServiceFile(gen, f)
	template.GenerateServiceImplFile(gen, f)
	template.GenerateControllerFile(gen, f)
	template.GenerateVoFile(gen, f)
}

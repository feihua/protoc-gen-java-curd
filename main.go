package main

import (
	"github/feihua/protoc-gen-java-curd/template"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
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

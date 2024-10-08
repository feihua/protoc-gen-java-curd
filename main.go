package main

import (
	"github/feihua/protoc-gen-java-curd/template"
	"google.golang.org/protobuf/compiler/protogen"
	"time"
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
	t := time.Now().Format("2006-01-02 15:04:05")
	template.GenerateEntityFile(gen, f, t)
	template.GenerateDaoFile(gen, f, t)
	template.GenerateXmlFile(gen, f, t)
	template.GenerateServiceFile(gen, f, t)
	template.GenerateBizFile(gen, f, t)
	template.GenerateBizImplFile(gen, f, t)
	template.GenerateServiceImplFile(gen, f, t)
	template.GenerateControllerFile(gen, f, t)
	template.GenerateControllerV3File(gen, f, t)
	template.GenerateVoFile(gen, f, t)
	template.GenerateVoV3File(gen, f, t)
}

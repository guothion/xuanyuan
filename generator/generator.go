package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// go:embed templates/*.tmpl
var embeddedTemplates embed.FS

func GenerateFile(templateFile, outputFile string, data TemplateData, templateDir string) error {
	var tmpl *template.Template
	var err error

	if _, err := os.Stat(templateDir); err == nil {
		tmplPath := filepath.Join(templateDir, templateFile)
		tmpl, err = template.ParseFiles(tmplPath)
		if err != nil {
			return fmt.Errorf("解析外部模板失败：%w", err)
		}
	} else {
		// 回到 embed 模板
		tmpl, err = template.ParseFS(embeddedTemplates, "templates/"+templateFile)
		if err != nil {
			return fmt.Errorf("解析嵌入模板失败：%w", err)
		}
	}

	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		return err
	}

	if _, err := os.Stat(outputFile); err == nil {
		fmt.Printf("⚠️文件已存在：%s，是否覆盖？[y/N]:", outputFile)
		var input string
		fmt.Scanln(&input)
		if input != "y" && input != "Y" {
			fmt.Println("跳过生成。")
			return nil
		}
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("执行模板失败：%w", err)
	}

	fmt.Printf("✅生成成功：%s\n", outputFile)
	return nil
}

func ParseFields(fieldsStr string) []Field {
	if fieldsStr == "" {
		return nil
	}

	var fields []Field
	for _, f := range strings.Split(fieldsStr, ",") {
		kv := strings.Split(f, ":")
		if len(kv) != 2 {
			continue
		}
		name, typ := kv[0], kv[1]
		jsonTag := strings.ToLower(name[:1] + name[1:])
		gormTag := ""
		if name == "ID" {
			gormTag = "primarykey"
		}
		fields = append(fields, Field{
			Name:    name,
			Type:    typ,
			JsonTag: jsonTag,
			GormTag: gormTag,
		})
	}
	return fields
}

package cmd

import (
	"fmt"
	"github.com/guothion/xuanyuan/generator"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	structName     string
	packageName    string
	fields         string
	description    string
	author         string
	outputDir      string
	templateDir    string
	withTimestamps bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "生成 Gin 项目代码",
	Long:  `支持通过命名参数生成 Model/Handler/Router`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		genType := args[0]

		// 解析字段
		fieldList := generator.ParseFields(fields)
		if genType == "model" && len(fieldList) == 0 {
			fmt.Fprintln(os.Stderr, "❌错误：生成 model必须提供--fields")
			os.Exit(1)
		}

		data := generator.TemplateData{
			StructName:    structName,
			PackageName:   packageName,
			RouteName:     toSnakeCase(structName),
			Fields:        fieldList,
			Description:   description,
			HasTimestamps: withTimestamps,
			Author:        author,
			Year:          time.Now().Year(),
		}

		if packageName == "" {
			data.PackageName = strings.ToLower(structName[:1]) + structName[1:]
		}

		templateFile := fmt.Sprintf("%s.go.tmpl", genType)
		outputFile := fmt.Sprintf("%s/%s.go", outputDir, toSnakeCase(structName))

		err := generator.GenerateFile(templateFile, outputFile, data, templateDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ 生成失败: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.Flags().StringVar(&structName, "name", "", "结构体名称（必填）")
	generateCmd.MarkFlagRequired("name")
	generateCmd.Flags().StringVar(&packageName, "package", "", "包名（可选，默认为结构体小驼峰）")

	generateCmd.Flags().StringVar(&fields, "fields", "", "字段列表，格式: name:type,name:type")

	generateCmd.Flags().StringVar(&description, "description", "", "描述信息")

	generateCmd.Flags().StringVar(&author, "author", getUserHomeDir(), "作者名")

	generateCmd.Flags().StringVar(&outputDir, "output", "internal", "输出目录")

	generateCmd.Flags().StringVar(&templateDir, "template-dir", "./templates", "自定义模板目录")

	generateCmd.Flags().BoolVar(&withTimestamps, "with-timestamps", false, "是否包含 created_at/updated_at 字段")

	//RootCmd.AddCommand(generateCmd)
}

// 获取用户主目录作为默认作者
func getUserHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERNAME")
	}
	return os.Getenv("USER")
}

// 驼峰转蛇形：StallApplication -> stall_application
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// 驼峰转短横线：StallApplication -> stall-application
func toKebabCase(s string) string {
	return strings.ReplaceAll(toSnakeCase(s), "_", "-")
}

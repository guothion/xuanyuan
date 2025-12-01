package generator

type Field struct {
	Name    string
	Type    string
	JsonTag string
	GormTag string
}

type TemplateData struct {
	StructName    string // 结构体名：StallApplication
	PackageName   string // 包名：model
	RouteName     string // 路由名：stall-application
	Fields        []Field
	Description   string // 描述
	HasTimestamps bool   // 是否有 created_at 等
	Author        string // 作者
	Year          int    // 年份
}

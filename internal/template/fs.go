package template

import (
	"embed"
	"html/template"
)

var (
	//go:embed resources
	dir embed.FS

	tmplIndex  *template.Template
	tmplDetail *template.Template
)

func init() {
	tmplIndex = template.Must(template.ParseFS(dir, "*/base.gohtml", "*/index.gohtml"))
	tmplDetail = template.Must(template.ParseFS(dir, "*/base.gohtml", "*/detail.gohtml"))
}

func Index() *template.Template {
	return tmplIndex
}

func Detail() *template.Template {
	return tmplDetail
}

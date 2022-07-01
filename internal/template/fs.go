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
	debug      bool
)

func init() {
	tmplIndex = template.Must(template.ParseFS(dir, "*/base.gohtml", "*/index.gohtml"))
	tmplDetail = template.Must(template.ParseFS(dir, "*/base.gohtml", "*/detail.gohtml"))
}

func Init(needReload bool) {
	debug = needReload
}

func Index() *template.Template {
	if debug {
		tmplIndex = template.Must(template.ParseFiles("internal/template/resources/base.gohtml", "internal/template/resources/index.gohtml"))
	}

	return tmplIndex
}

func Detail() *template.Template {
	if debug {
		tmplDetail = template.Must(template.ParseFiles("internal/template/resources/base.gohtml", "internal/template/resources/detail.gohtml"))
	}

	return tmplDetail
}

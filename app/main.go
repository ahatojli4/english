package main

import (
	"github.com/ahatojli4/english/internal/http"
	_ "github.com/ahatojli4/english/internal/template"
)

func main() {
	//fmt.Println(posts.GetFileNames())

	//tmplIndex := template.GetTemplate("base.gohtml")

	err := http.Start("8080")
	if err != nil {
		panic(err)
	}
}

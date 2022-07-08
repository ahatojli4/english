package main

import (
	"flag"

	"github.com/ahatojli4/english/internal/http"
	"github.com/ahatojli4/english/internal/template"
)

var (
	port            = flag.String("port", "8080", "port to listen")
	reloadTemplates = flag.Bool("rt", false, "reload templates")
)

func main() {
	flag.Parse()
	template.Init(*reloadTemplates)

	//fmt.Println(notes.GetFileNames())

	//tmplIndex := template.GetTemplate("base.gohtml")

	err := http.Start(*port)
	if err != nil {
		panic(err)
	}
}

package http

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gomarkdown/markdown"

	"github.com/ahatojli4/english/internal/posts"
	t "github.com/ahatojli4/english/internal/template"
)

var routes = map[string]http.HandlerFunc{
	"/":      index,
	"/list/": detail,
}

func Start(port string) error {
	for route, handlerFunc := range routes {
		http.HandleFunc(route, handlerFunc)
	}

	return http.ListenAndServe(":"+port, nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	tmpl := t.Index()
	postNames := posts.GetFileNames()
	//fmt.Println(postNames)
	/*for i := 0; i < len(tmpl.DefinedTemplates()); i++ {

	}*/
	err := tmpl.ExecuteTemplate(writer, "base", postNames)
	if err != nil {
		//todo: add logger
		return
	}
	//err = tmpl.ExecuteTemplate(writer, "main", postNames)
	//if err != nil {
	//	//todo: add logger
	//	return
	//}
}

func detail(writer http.ResponseWriter, request *http.Request) {
	slug, _, _ := strings.Cut(
		strings.TrimPrefix(request.URL.Path, "/list/"),
		"/",
	)
	if len(slug) == 0 {
		index(writer, request)
		return
	}
	content := posts.GetContent(slug)
	html := markdown.ToHTML(content, nil, nil)
	tmpl := t.Detail()
	err := tmpl.ExecuteTemplate(writer, "base", template.HTML(html))
	if err != nil {
		//todo: add logger
		return
	}
	//err = tmpl.ExecuteTemplate(writer, "main", string(content))
	//if err != nil {
	//	//todo: add logger
	//	return
	//}
}

package http

import (
	"html/template"
	"net/http"
	"strings"

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
	noteList := posts.GetNoteList()
	data := t.DataList{
		Items: make([]t.DataDetail, 0, len(noteList)),
	}
	for _, note := range noteList {
		data.Items = append(data.Items, t.DataDetail{
			Title:      note.GetTitle(),
			FileName:   note.GetFileName(),
			DateCreate: note.GetDateCreated(),
		})
	}
	err := tmpl.ExecuteTemplate(writer, "base", data)
	if err != nil {
		//todo: add logger
		return
	}
	//err = tmpl.ExecuteTemplate(writer, "main", noteList)
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
	note := posts.GetNote(slug)
	data := t.DataDetail{
		Title:      note.GetTitle(),
		DateCreate: note.GetDateCreated(),
		Content:    template.HTML(note.GetContent()),
	}
	tmpl := t.Detail()
	err := tmpl.ExecuteTemplate(writer, "base", data)
	if err != nil {
		//todo: add logger
		return
	}
	//err = tmpl.ExecuteTemplate(writer, "main", string(note))
	//if err != nil {
	//	//todo: add logger
	//	return
	//}
}

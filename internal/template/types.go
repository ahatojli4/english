package template

import (
	"html/template"
	"time"
)

type DataDetail struct {
	Title      string
	FileName   string
	DateCreate time.Time
	Content    template.HTML
}

type DataList struct {
	Items []DataDetail
}

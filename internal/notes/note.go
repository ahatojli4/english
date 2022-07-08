package notes

import (
	"fmt"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type note struct {
	fileName    string
	title       string
	dateCreated time.Time
	rawContent  []byte
	content     []byte
}

func newNote(fileName string, rawContent []byte) *note {
	n := &note{
		fileName:   fileName,
		rawContent: rawContent,
	}
	astTree := markdown.Parse(rawContent, parser.New())
	toRemove := make([]ast.Node, 0)
	ast.WalkFunc(astTree, func(node ast.Node, entering bool) ast.WalkStatus {
		switch h := node.(type) {
		case *ast.Heading:
			if h.IsSpecial && !entering {
				parseSpecialHeader(n, string(h.Literal))
				toRemove = append(toRemove, h)

				return ast.GoToNext
			}
		default:
		}

		return ast.GoToNext
	})
	for i := range toRemove {
		ast.RemoveFromTree(toRemove[i])
	}
	opts := html.RendererOptions{
		Flags: html.CommonFlags,
	}
	renderer := html.NewRenderer(opts)
	rRes := markdown.Render(astTree, renderer)
	n.content = rRes

	return n
}

type byDate []*note

func (n byDate) Len() int {
	return len(n)
}

func (n byDate) Less(i, j int) bool {
	return n[i].dateCreated.Before(n[j].dateCreated)
}

func (n byDate) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type Note interface {
	GetTitle() string
	GetFileName() string
	GetDateCreated() time.Time
	GetContent() []byte
}

func (n *note) GetTitle() string {
	return n.title
}

func (n *note) GetDateCreated() time.Time {
	return n.dateCreated
}

func (n *note) GetContent() []byte {
	return n.content
}

func (n *note) GetFileName() string {
	return n.fileName
}

func parseSpecialHeader(p *note, s string) {
	elements := strings.Split(s, ":")
	if len(elements) != 2 {
		fmt.Println("Invalid special header:", s)
		return
	}
	key, value := strings.TrimSpace(elements[0]), strings.TrimSpace(elements[1])
	switch key {
	case "title":
		p.title = value
	case "date_create":
		date, err := time.Parse("2006-01-02", value)
		if err != nil {
			fmt.Println(err)
			// todo: add logger (wrong format)
			return
		}
		p.dateCreated = date
	default:
		fmt.Println("unsupported meta info")
		// todo: add logger
	}
}

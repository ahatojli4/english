package posts

import (
	"embed"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

var (
	//go:embed resources
	dir embed.FS

	notes    []*note
	mapNotes map[string]int
)

func init() {
	entries, err := dir.ReadDir("resources")
	if err != nil {
		// todo: add logger
		panic(err)
	}
	notes = make([]*note, 0, len(entries))
	mapNotes = make(map[string]int, len(entries))
	for i, entry := range entries {
		file, err := dir.ReadFile(filepath.Join("resources", entry.Name()))
		if err != nil {
			// todo: add logger
			fmt.Println(err)
			continue
		}
		if entry.IsDir() {
			// todo: add logger
			fmt.Println("recursive walking dir doesn't support")
			continue
		}
		entryName := strings.ToLower(entry.Name())
		n := newNote(entryName, file)
		notes = append(notes, n)
		mapNotes[entryName] = i
	}
	sort.Sort(byDate(notes))
	for i, n := range notes {
		mapNotes[n.fileName] = i
	}
}

func GetNoteList() []Note {
	res := make([]Note, 0, len(notes))
	for i := range notes {
		res = append(res, Note(notes[i]))
	}

	return res
}

func GetNote(fileName string) Note {
	index, ok := mapNotes[strings.ToLower(fileName)]
	if !ok {
		return &note{
			fileName: fileName,
			content:  []byte("404"),
		}
	}

	return notes[index]
}

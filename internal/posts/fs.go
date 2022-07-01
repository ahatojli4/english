package posts

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	//go:embed resources
	dir embed.FS

	mapFiles map[string][]byte
)

func init() {
	entries, err := dir.ReadDir("resources")
	if err != nil {
		// todo: add logger
		panic(err)
	}
	mapFiles = make(map[string][]byte, len(entries))
	for _, entry := range entries {
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
		mapFiles[strings.ToLower(entry.Name())] = file
	}
}

// todo: add sorting by date
func GetFileNames() []string {
	keys := make([]string, 0, len(mapFiles))
	for k := range mapFiles {
		keys = append(keys, k)
	}

	return keys
}

func GetContent(fileName string) []byte {
	content, ok := mapFiles[strings.ToLower(fileName)]
	if !ok {
		return []byte("404")
	}

	return content
}

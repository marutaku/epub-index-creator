package book

import (
	"os"
	"path"
)

type Page struct {
	Id   string
	Path string
}

func NewPage(id, path string) *Page {
	return &Page{
		Id:   id,
		Path: path,
	}
}

func (p *Page) Content(baseDir string) (string, error) {
	contentsBytes, err := os.ReadFile(path.Join(baseDir, p.Path))
	if err != nil {
		return "", err
	}
	return string(contentsBytes), nil
}

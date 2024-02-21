package book

import (
	"errors"
	"os"
	"regexp"
	"strings"
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

func (p *Page) Content() (string, error) {
	contentsBytes, err := os.ReadFile(p.Path)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`(?s)<body.*?>(.*)</body>`)
	res := re.FindAllStringSubmatch(string(contentsBytes), 1)
	if len(res) == 0 {
		return "", errors.New("body tag not found")
	}
	return strings.TrimSpace(res[0][1]), nil
}

package domain

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/marutaku/epub-index-creator/ent"
)

type Page struct {
	Id   int
	Path string
}

func NewPage(id int, path string) *Page {
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

func (p *Page) Title() (string, error) {
	contentsBytes, err := os.ReadFile(p.Path)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	res := re.FindAllStringSubmatch(string(contentsBytes), 1)
	if len(res) == 0 {
		return "", errors.New("title tag not found")
	}
	return res[0][1], nil
}

func (p *Page) ExtractKeywords() ([]string, error) {
	contents, err := p.Content()
	if err != nil {
		return nil, err
	}
	// Extract words that inside <b> or <strong> tag
	re := regexp.MustCompile(`(?i)<(b|strong)>(.*?)</(b|strong)>`)
	keywordsWithTag := re.FindAllString(contents, -1)
	keywords := make([]string, 0, len(keywordsWithTag))
	for _, keywordWithTag := range keywordsWithTag {
		re = regexp.MustCompile(`(?i)<(b|strong)>(.*?)</(b|strong)>`)
		res := re.FindAllStringSubmatch(keywordWithTag, 1)
		keywords = append(keywords, res[0][2])
	}
	return keywords, nil
}

func NewPageFromEnt(entPage *ent.Page) *Page {
	return NewPage(entPage.ID, entPage.Path)
}

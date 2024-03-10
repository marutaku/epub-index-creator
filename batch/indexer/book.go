package indexer

import (
	"os"
	"path"

	"github.com/antchfx/xmlquery"
	"github.com/marutaku/epub-index-creator/domain"
)

type HTMLBook struct {
	ISBN      string
	Title     string
	Language  string
	Author    string
	Publisher string
	Pages     []*HTMLPage
}

func NewHTMLBook(isbn, title, language, author, publisher string, pages []*HTMLPage) *HTMLBook {
	return &HTMLBook{
		ISBN:      isbn,
		Title:     title,
		Language:  language,
		Author:    author,
		Publisher: publisher,
		Pages:     pages,
	}
}

func NewHTMLBookFromOPF(filepath string) (*HTMLBook, error) {
	baseDir := path.Dir(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	doc, err := xmlquery.Parse(file)
	if err != nil {
		return nil, err
	}
	isbn := xmlquery.FindOne(doc, "//package/metadata/dc:identifier").InnerText()
	title := xmlquery.FindOne(doc, "//package/metadata/dc:title").InnerText()
	language := xmlquery.FindOne(doc, "//package/metadata/dc:language").InnerText()
	author := xmlquery.FindOne(doc, "//package/metadata/dc:creator").InnerText()
	publisher := xmlquery.FindOne(doc, "//package/metadata/dc:publisher").InnerText()
	items := xmlquery.Find(doc, "//package/manifest/item[@media-type='application/xhtml+xml']")
	pages := make([]*HTMLPage, 0, len(items))
	for index, item := range items {
		pagePath := item.SelectAttr("href")
		pages = append(pages, NewHTMLPage(index+1, path.Join(baseDir, pagePath)))
	}
	return NewHTMLBook(isbn, title, language, author, publisher, pages), nil
}

func (b *HTMLBook) ToDomain() (*domain.Book, error) {
	pages := make([]*domain.Page, 0, len(b.Pages))
	for _, page := range b.Pages {
		pageDomain, err := page.ToDomain()
		if err != nil {
			return nil, err
		}
		pages = append(pages, pageDomain)
	}
	return domain.NewBook(b.ISBN, b.Title, b.Language, b.Author, b.Publisher, pages), nil
}

package controller

import (
	"context"
	"log"

	epubIndexCreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
)

// epub_index_creator service example implementation.
// The example methods log the requests and return zero values.
type epubIndexCreatorsrvc struct {
	logger *log.Logger
}

// NewEpubIndexCreator returns the epub_index_creator service implementation.
func NewEpubIndexCreator(logger *log.Logger) epubIndexCreator.Service {
	return &epubIndexCreatorsrvc{logger}
}

// List implements List.
func (s *epubIndexCreatorsrvc) ListBooks(ctx context.Context, p *epubIndexCreator.ListBooksPayload) (res []*epubIndexCreator.Book, err error) {
	res = []*epubIndexCreator.Book{}
	s.logger.Print("epubIndexCreator.ListBooks")
	return
}

func (s *epubIndexCreatorsrvc) FindBook(ctx context.Context, p *epubIndexCreator.FindBookPayload) (res *epubIndexCreator.Book, err error) {
	res = &epubIndexCreator.Book{}
	s.logger.Print("epubIndexCreator.FindBook")
	return
}

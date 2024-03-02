package calc

import (
	"context"
	"log"

	epubindexcreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
)

// epub_index_creator service example implementation.
// The example methods log the requests and return zero values.
type epubIndexCreatorsrvc struct {
	logger *log.Logger
}

// NewEpubIndexCreator returns the epub_index_creator service implementation.
func NewEpubIndexCreator(logger *log.Logger) epubindexcreator.Service {
	return &epubIndexCreatorsrvc{logger}
}

// List implements List.
func (s *epubIndexCreatorsrvc) List(ctx context.Context, p *epubindexcreator.ListPayload) (res *epubindexcreator.Book, err error) {
	res = &epubindexcreator.Book{}
	s.logger.Print("epubIndexCreator.List")
	return
}

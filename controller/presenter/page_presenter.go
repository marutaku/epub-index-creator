package presenter

import (
	"github.com/marutaku/epub-index-creator/domain"
	epubIndexCreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
)

func ConvertPageToResponse(page *domain.Page) *epubIndexCreator.PageResponse {
	return &epubIndexCreator.PageResponse{
		ID:    page.Id,
		Title: page.Title,
	}
}

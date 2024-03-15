package presenter

import (
	"github.com/marutaku/epub-index-creator/domain"
	epubIndexCreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
)

func ConvertKeywordToResponse(keyword *domain.Keyword) *epubIndexCreator.KeywordResponse {
	return &epubIndexCreator.KeywordResponse{
		ID:      keyword.ID,
		Keyword: keyword.Keyword,
	}
}

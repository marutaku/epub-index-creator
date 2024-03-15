package usecase

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/infra/repository"
)

type KeywordUsecaseInterface interface {
	ListKeywords(ctx context.Context, limit, offset int) ([]*domain.Keyword, error)
	ListKeywordsByPageID(ctx context.Context, pageID int) ([]*domain.Keyword, error)
}

type keywordUsecase struct{}

func NewKeywordUsecase() KeywordUsecaseInterface {
	return &keywordUsecase{}
}

func (*keywordUsecase) ListKeywords(ctx context.Context, limit, offset int) ([]*domain.Keyword, error) {
	keywordRep := repository.NewKeywordRepository()
	keywords, err := keywordRep.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

func (*keywordUsecase) ListKeywordsByPageID(ctx context.Context, pageID int) ([]*domain.Keyword, error) {
	keywordRep := repository.NewKeywordRepository()
	keywords, err := keywordRep.FindAllByPageID(ctx, pageID)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

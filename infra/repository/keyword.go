package repository

import (
	"context"
	"fmt"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"

	pageModel "github.com/marutaku/epub-index-creator/ent/page"
	"github.com/marutaku/epub-index-creator/ent/predicate"
)

type KeywordRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]*domain.Keyword, error)
	FindAllByPageID(ctx context.Context, pageID int) ([]*domain.Keyword, error)
}

type KeywordRepositoryImpl struct{}

func NewKeywordRepository() KeywordRepository {
	return &KeywordRepositoryImpl{}
}

func (*KeywordRepositoryImpl) FindAll(ctx context.Context, limit, offset int) ([]*domain.Keyword, error) {
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		return nil, fmt.Errorf("ent.TxFromContext(ctx) is nil")
	}
	keywords, err := tx.Keyword.Query().Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*domain.Keyword
	for _, keyword := range keywords {
		result = append(result, domain.NewKeywordFromEnt(keyword))
	}
	return result, nil
}

func (*KeywordRepositoryImpl) FindAllByPageID(ctx context.Context, pageID int) ([]*domain.Keyword, error) {
	tx := ent.TxFromContext(ctx)
	keywords, err := tx.Keyword.Query().Where(predicate.Keyword(pageModel.ID(pageID))).All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*domain.Keyword
	for _, keyword := range keywords {
		result = append(result, domain.NewKeywordFromEnt(keyword))
	}
	return result, nil
}

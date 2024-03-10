package repository

import (
	"context"
	"fmt"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"
)

type PageRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]*domain.Page, error)
	FindByID(ctx context.Context, id int) (*domain.Page, error)
	Save(ctx context.Context, path string) (*domain.Page, error)
	Update(ctx context.Context, id int, page *domain.Page) (*domain.Page, error)
	Delete(ctx context.Context, id int) error
}

type PageRepositoryImpl struct{}

func NewPageRepository() PageRepository {
	return &PageRepositoryImpl{}
}

func (r *PageRepositoryImpl) FindAll(ctx context.Context, limit, offset int) ([]*domain.Page, error) {
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		return nil, fmt.Errorf("ent.TxFromContext(ctx) is nil")
	}
	pages, err := tx.Page.Query().Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*domain.Page
	for _, page := range pages {
		result = append(result, domain.NewPageFromEnt(page))
	}
	return result, nil
}

func (r *PageRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Page, error) {
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		return nil, fmt.Errorf("ent.TxFromContext(ctx) is nil")
	}
	page, err := tx.Page.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.NewPageFromEnt(page), nil
}

func (r *PageRepositoryImpl) Save(ctx context.Context, path string) (*domain.Page, error) {
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		return nil, fmt.Errorf("ent.TxFromContext(ctx) is nil")
	}
	page, err := tx.Page.Create().SetPath(path).Save(ctx)
	if err != nil {
		return nil, err
	}
	return domain.NewPageFromEnt(page), nil
}

func (r *PageRepositoryImpl) Update(ctx context.Context, id int, page *domain.Page) (*domain.Page, error) {
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		return nil, fmt.Errorf("ent.TxFromContext(ctx) is nil")
	}
	newEntPage, err := tx.Page.UpdateOneID(id).SetPath(page.Path).Save(ctx)
	if err != nil {
		return nil, err
	}
	return domain.NewPageFromEnt(newEntPage), nil
}

func (r *PageRepositoryImpl) Delete(ctx context.Context, id int) error {
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		return fmt.Errorf("ent.TxFromContext(ctx) is nil")
	}
	return tx.Page.DeleteOneID(id).Exec(ctx)
}

package usecase

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/infra/repository"
)

type PageUsecaseInterface interface {
	ListPages(ctx context.Context, limit, offset int) ([]*domain.Page, error)
	FindPage(ctx context.Context, id int) (*domain.Page, error)
	CreatePage(ctx context.Context, title string) (*domain.Page, error)
	UpdatePage(ctx context.Context, id int, title string) (*domain.Page, error)
	DeletePage(ctx context.Context, id int) error
}

type pageUsecase struct{}

func NewPageUsecase() PageUsecaseInterface {
	return &pageUsecase{}
}

func (*pageUsecase) ListPages(ctx context.Context, limit, offset int) ([]*domain.Page, error) {
	pageRepo := repository.NewPageRepository()
	pages, err := pageRepo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func (*pageUsecase) FindPage(ctx context.Context, id int) (*domain.Page, error) {
	pageRepo := repository.NewPageRepository()
	page, err := pageRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (*pageUsecase) CreatePage(ctx context.Context, path string) (*domain.Page, error) {
	pageRepo := repository.NewPageRepository()
	page, err := pageRepo.Save(ctx, path)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (*pageUsecase) UpdatePage(ctx context.Context, id int, path string) (*domain.Page, error) {
	pageRepo := repository.NewPageRepository()
	page, err := pageRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	page.Path = path
	newPage, err := pageRepo.Update(ctx, page.Id, page)
	if err != nil {
		return nil, err
	}
	return newPage, nil
}

func (*pageUsecase) DeletePage(ctx context.Context, id int) error {
	pageRepo := repository.NewPageRepository()
	err := pageRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

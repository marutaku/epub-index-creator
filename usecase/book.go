package usecase

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
)

type BookUsecaseInterface interface {
	ListBooks(ctx context.Context, limit, offset int) ([]*domain.Book, error)
	FindBook(ctx context.Context, isbn string) (*domain.Book, error)
	Save(ctx context.Context, book domain.Book) error
}

type bookUsecase struct{}

func NewBookUsecase() *bookUsecase {
	return &bookUsecase{}
}

func (*bookUsecase) ListBooks(ctx context.Context, limit, offset int) ([]*domain.Book, error) {
	return nil, nil
}

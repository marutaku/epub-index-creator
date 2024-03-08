package usecase

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/infra/repository"
)

type BookUsecaseInterface interface {
	ListBooks(ctx context.Context, limit, offset int) ([]*domain.Book, error)
	FindBook(ctx context.Context, isbn string) (*domain.Book, error)
	CreateBook(ctx context.Context, isbn string, title string, author string, language string, publisher string) (*domain.Book, error)
	UpdateBook(ctx context.Context, isbn string, title string, author string, language string, publisher string) (*domain.Book, error)
	DeleteBook(ctx context.Context, isbn string) error
}

type bookUsecase struct{}

func NewBookUsecase() BookUsecaseInterface {
	return &bookUsecase{}
}

func (*bookUsecase) ListBooks(ctx context.Context, limit, offset int) ([]*domain.Book, error) {
	bookRepo := repository.NewBookRepository()
	books, err := bookRepo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (*bookUsecase) FindBook(ctx context.Context, isbn string) (*domain.Book, error) {
	bookRepo := repository.NewBookRepository()
	book, err := bookRepo.FindByISBN(ctx, isbn)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*bookUsecase) CreateBook(ctx context.Context, isbn string, title string, author string, language string, publisher string) (*domain.Book, error) {
	bookRepo := repository.NewBookRepository()
	book := domain.NewBook(isbn, title, author, language, publisher, make([]*domain.Page, 0))
	err := bookRepo.Save(ctx, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*bookUsecase) UpdateBook(ctx context.Context, isbn string, title string, author string, language string, publisher string) (*domain.Book, error) {
	bookRepo := repository.NewBookRepository()
	book, err := bookRepo.FindByISBN(ctx, isbn)
	if err != nil {
		return nil, err
	}
	book.Title = title
	book.Author = author
	book.Language = language
	book.Publisher = publisher
	err = bookRepo.Update(ctx, isbn, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*bookUsecase) DeleteBook(ctx context.Context, isbn string) error {
	bookRepo := repository.NewBookRepository()
	err := bookRepo.Delete(ctx, isbn)
	if err != nil {
		return err
	}
	return nil
}

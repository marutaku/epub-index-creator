package repository

import (
	"context"

	"github.com/marutaku/epub-index-creator/domain"
	"github.com/marutaku/epub-index-creator/ent"
	bookModel "github.com/marutaku/epub-index-creator/ent/book"
)

type BookRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]*domain.Book, error)
	FindByISBN(ctx context.Context, isbn string) (*domain.Book, error)
	Save(ctx context.Context, book *domain.Book) error
}

type BookRepositoryImpl struct{}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) FindAll(ctx context.Context, limit, offset int) ([]*domain.Book, error) {
	tx := ent.TxFromContext(ctx)
	books, err := tx.Book.Query().Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*domain.Book
	for _, book := range books {
		result = append(result, domain.NewBookFromEnt(book))
	}
	return result, nil
}

func (r *BookRepositoryImpl) FindByISBN(ctx context.Context, isbn string) (*domain.Book, error) {
	tx := ent.TxFromContext(ctx)
	book, err := tx.Book.Query().Where(bookModel.Isbn(isbn)).First(ctx)
	if err != nil {
		return nil, err
	}
	return domain.NewBookFromEnt(book), nil
}

func (r *BookRepositoryImpl) Save(ctx context.Context, book *domain.Book) error {
	tx := ent.TxFromContext(ctx)
	entBook, err := tx.Book.Create().
		SetIsbn(book.ISBN).
		SetTitle(book.Title).
		SetLanguage(book.Language).
		SetAuthor(book.Author).
		SetPublisher(book.Publisher).
		Save(ctx)
	if err != nil {
		return err
	}
	for _, page := range book.Pages {
		title, err := page.Title()
		if err != nil {
			return err
		}
		_, err = tx.Page.Create().
			SetBook(entBook).
			SetTitle(title).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

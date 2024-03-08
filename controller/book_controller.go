package controller

import (
	"context"
	"log"

	"github.com/marutaku/epub-index-creator/domain"
	epubIndexCreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
	"github.com/marutaku/epub-index-creator/usecase"
)

// epub_index_creator service example implementation.
// The example methods log the requests and return zero values.
type epubIndexCreatorsrvc struct {
	logger      *log.Logger
	bookUsecase usecase.BookUsecaseInterface
}

// NewEpubIndexCreator returns the epub_index_creator service implementation.
func NewEpubIndexCreator(logger *log.Logger) epubIndexCreator.Service {
	return &epubIndexCreatorsrvc{logger: logger, bookUsecase: usecase.NewBookUsecase()}
}

func convertBookToResponse(book *domain.Book) *epubIndexCreator.Book {
	return &epubIndexCreator.Book{
		Isbn:   epubIndexCreator.ISBN(book.ISBN),
		Title:  book.Title,
		Author: book.Author,
		Pages:  make([]*epubIndexCreator.Page, 0), // TODO: Implement
	}
}

// List implements List.
func (s *epubIndexCreatorsrvc) ListBooks(ctx context.Context, p *epubIndexCreator.ListBooksPayload) (res []*epubIndexCreator.Book, err error) {
	res = []*epubIndexCreator.Book{}
	s.logger.Print("epubIndexCreator.ListBooks")
	books, err := s.bookUsecase.ListBooks(ctx, p.Limit, p.Offset)
	for _, book := range books {
		res = append(res, convertBookToResponse(book))
	}
	return
}

func (s *epubIndexCreatorsrvc) FindBook(ctx context.Context, p *epubIndexCreator.FindBookPayload) (res *epubIndexCreator.Book, err error) {
	res = &epubIndexCreator.Book{}
	s.logger.Print("epubIndexCreator.FindBook")
	book, err := s.bookUsecase.FindBook(ctx, string(p.Isbn))
	if err != nil {
		return nil, err
	}
	res = convertBookToResponse(book)
	return
}

func (s *epubIndexCreatorsrvc) CreateBook(ctx context.Context, p *epubIndexCreator.Book) (res *epubIndexCreator.Book, err error) {
	s.logger.Print("epubIndexCreator.CreateBook")
	book, err := s.bookUsecase.CreateBook(ctx, string(p.Isbn), p.Title, p.Author, p.Language, p.Publisher)
	if err != nil {
		return nil, err
	}
	res = convertBookToResponse(book)
	return
}

func (s *epubIndexCreatorsrvc) UpdateBook(ctx context.Context, p *epubIndexCreator.UpdateBookPayload) (res *epubIndexCreator.Book, err error) {
	res = &epubIndexCreator.Book{}
	s.logger.Print("epubIndexCreator.UpdateBook")
	book, err := s.bookUsecase.UpdateBook(ctx, string(p.Isbn), p.Title, p.Author, p.Language, p.Publisher)
	if err != nil {
		return nil, err
	}
	res = convertBookToResponse(book)
	return
}

func (s *epubIndexCreatorsrvc) DeleteBook(ctx context.Context, p *epubIndexCreator.DeleteBookPayload) (err error) {
	s.logger.Print("epubIndexCreator.DeleteBook")
	err = s.bookUsecase.DeleteBook(ctx, string(p.Isbn))
	if err != nil {
		return err
	}
	return
}

func (s *epubIndexCreatorsrvc) CreatePage(ctx context.Context, p *epubIndexCreator.CreatePagePayload) (res *epubIndexCreator.Page, err error) {
	res = &epubIndexCreator.Page{}
	s.logger.Print("epubIndexCreator.CreatePage")
	return
}

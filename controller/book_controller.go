package controller

import (
	"context"
	"log"

	"github.com/marutaku/epub-index-creator/controller/presenter"
	epubIndexCreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
	"github.com/marutaku/epub-index-creator/usecase"
)

// epub_index_creator service example implementation.
// The example methods log the requests and return zero values.
type epubIndexCreatorsrvc struct {
	logger         *log.Logger
	bookUsecase    usecase.BookUsecaseInterface
	pageUsecase    usecase.PageUsecaseInterface
	keywordUsecase usecase.KeywordUsecaseInterface
}

// NewEpubIndexCreator returns the epub_index_creator service implementation.
func NewEpubIndexCreator(logger *log.Logger) epubIndexCreator.Service {
	return &epubIndexCreatorsrvc{logger: logger, bookUsecase: usecase.NewBookUsecase(), pageUsecase: usecase.NewPageUsecase(), keywordUsecase: usecase.NewKeywordUsecase()}
}

// List implements List.
func (s *epubIndexCreatorsrvc) ListBooks(ctx context.Context, p *epubIndexCreator.ListBooksPayload) (res []*epubIndexCreator.BookResponse, err error) {
	res = []*epubIndexCreator.BookResponse{}
	s.logger.Print("epubIndexCreator.ListBooks")
	books, err := s.bookUsecase.ListBooks(ctx, p.Limit, p.Offset)
	for _, book := range books {
		res = append(res, presenter.ConvertBookToResponse(book))
	}
	return
}

func (s *epubIndexCreatorsrvc) FindBook(ctx context.Context, p *epubIndexCreator.FindBookPayload) (res *epubIndexCreator.BookResponse, err error) {
	res = &epubIndexCreator.BookResponse{}
	s.logger.Print("epubIndexCreator.FindBook")
	book, err := s.bookUsecase.FindBook(ctx, string(p.Isbn))
	if err != nil {
		return nil, err
	}
	res = presenter.ConvertBookToResponse(book)
	return
}

func (s *epubIndexCreatorsrvc) CreateBook(ctx context.Context, p *epubIndexCreator.BookRequest) (res *epubIndexCreator.BookResponse, err error) {
	s.logger.Print("epubIndexCreator.CreateBook")
	book, err := s.bookUsecase.CreateBook(ctx, string(p.Isbn), p.Title, p.Author, p.Language, p.Publisher)
	if err != nil {
		return nil, err
	}
	res = presenter.ConvertBookToResponse(book)
	return
}

func (s *epubIndexCreatorsrvc) UpdateBook(ctx context.Context, p *epubIndexCreator.BookRequest) (res *epubIndexCreator.BookResponse, err error) {
	res = &epubIndexCreator.BookResponse{}
	s.logger.Print("epubIndexCreator.UpdateBook")
	book, err := s.bookUsecase.UpdateBook(ctx, string(p.Isbn), p.Title, p.Author, p.Language, p.Publisher)
	if err != nil {
		return nil, err
	}
	res = presenter.ConvertBookToResponse(book)
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

func (s *epubIndexCreatorsrvc) ListPages(ctx context.Context, p *epubIndexCreator.ListPagesPayload) (res []*epubIndexCreator.PageResponse, err error) {
	res = []*epubIndexCreator.PageResponse{}
	s.logger.Print("epubIndexCreator.ListPages")
	pages, err := s.pageUsecase.ListPages(ctx, p.Limit, p.Offset)
	for _, page := range pages {
		res = append(res, presenter.ConvertPageToResponse(page))
	}
	return
}

func (s *epubIndexCreatorsrvc) FindPage(ctx context.Context, p *epubIndexCreator.FindPagePayload) (res *epubIndexCreator.PageResponse, err error) {
	res = &epubIndexCreator.PageResponse{}
	s.logger.Print("epubIndexCreator.FindPage")
	page, err := s.pageUsecase.FindPage(ctx, p.PageID)
	if err != nil {
		return nil, err
	}
	res = presenter.ConvertPageToResponse(page)
	return
}

func (s *epubIndexCreatorsrvc) CreatePage(ctx context.Context, p *epubIndexCreator.CreatePagePayload) (res *epubIndexCreator.PageResponse, err error) {
	res = &epubIndexCreator.PageResponse{}
	s.logger.Print("epubIndexCreator.CreatePage")
	page, err := s.pageUsecase.CreatePage(ctx, p.Page.Title)
	if err != nil {
		return nil, err
	}
	res = presenter.ConvertPageToResponse(page)
	return
}

func (s *epubIndexCreatorsrvc) UpdatePage(ctx context.Context, p *epubIndexCreator.UpdatePagePayload) (res *epubIndexCreator.PageResponse, err error) {
	res = &epubIndexCreator.PageResponse{}
	s.logger.Print("epubIndexCreator.UpdatePage")
	page, err := s.pageUsecase.UpdatePage(ctx, p.PageID, p.Page.Title)
	if err != nil {
		return nil, err
	}
	res = presenter.ConvertPageToResponse(page)
	return
}

func (s *epubIndexCreatorsrvc) DeletePage(ctx context.Context, p *epubIndexCreator.DeletePagePayload) (err error) {
	s.logger.Print("epubIndexCreator.DeletePage")
	err = s.pageUsecase.DeletePage(ctx, p.PageID)
	if err != nil {
		return err
	}
	return
}

func (s *epubIndexCreatorsrvc) ListKeywordsInPage(ctx context.Context, p *epubIndexCreator.ListKeywordsInPagePayload) (res []*epubIndexCreator.KeywordResponse, err error) {
	res = []*epubIndexCreator.KeywordResponse{}
	s.logger.Print("epubIndexCreator.ListKeywords")
	keywords, err := s.keywordUsecase.ListKeywordsByPageID(ctx, p.PageID)
	for _, keyword := range keywords {
		res = append(res, presenter.ConvertKeywordToResponse(keyword))
	}
	return
}

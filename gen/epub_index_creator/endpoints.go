// Code generated by goa v3.15.0, DO NOT EDIT.
//
// epub_index_creator endpoints
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package epubindexcreator

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "epub_index_creator" service endpoints.
type Endpoints struct {
	ListBooks  goa.Endpoint
	FindBook   goa.Endpoint
	CreateBook goa.Endpoint
	UpdateBook goa.Endpoint
	DeleteBook goa.Endpoint
	ListPages  goa.Endpoint
	FindPage   goa.Endpoint
	CreatePage goa.Endpoint
	UpdatePage goa.Endpoint
	DeletePage goa.Endpoint
}

// NewEndpoints wraps the methods of the "epub_index_creator" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		ListBooks:  NewListBooksEndpoint(s),
		FindBook:   NewFindBookEndpoint(s),
		CreateBook: NewCreateBookEndpoint(s),
		UpdateBook: NewUpdateBookEndpoint(s),
		DeleteBook: NewDeleteBookEndpoint(s),
		ListPages:  NewListPagesEndpoint(s),
		FindPage:   NewFindPageEndpoint(s),
		CreatePage: NewCreatePageEndpoint(s),
		UpdatePage: NewUpdatePageEndpoint(s),
		DeletePage: NewDeletePageEndpoint(s),
	}
}

// Use applies the given middleware to all the "epub_index_creator" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListBooks = m(e.ListBooks)
	e.FindBook = m(e.FindBook)
	e.CreateBook = m(e.CreateBook)
	e.UpdateBook = m(e.UpdateBook)
	e.DeleteBook = m(e.DeleteBook)
	e.ListPages = m(e.ListPages)
	e.FindPage = m(e.FindPage)
	e.CreatePage = m(e.CreatePage)
	e.UpdatePage = m(e.UpdatePage)
	e.DeletePage = m(e.DeletePage)
}

// NewListBooksEndpoint returns an endpoint function that calls the method
// "ListBooks" of service "epub_index_creator".
func NewListBooksEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListBooksPayload)
		return s.ListBooks(ctx, p)
	}
}

// NewFindBookEndpoint returns an endpoint function that calls the method
// "FindBook" of service "epub_index_creator".
func NewFindBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*FindBookPayload)
		return s.FindBook(ctx, p)
	}
}

// NewCreateBookEndpoint returns an endpoint function that calls the method
// "CreateBook" of service "epub_index_creator".
func NewCreateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*BookRequest)
		return s.CreateBook(ctx, p)
	}
}

// NewUpdateBookEndpoint returns an endpoint function that calls the method
// "UpdateBook" of service "epub_index_creator".
func NewUpdateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*BookRequest)
		return s.UpdateBook(ctx, p)
	}
}

// NewDeleteBookEndpoint returns an endpoint function that calls the method
// "DeleteBook" of service "epub_index_creator".
func NewDeleteBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteBookPayload)
		return nil, s.DeleteBook(ctx, p)
	}
}

// NewListPagesEndpoint returns an endpoint function that calls the method
// "ListPages" of service "epub_index_creator".
func NewListPagesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPagesPayload)
		return s.ListPages(ctx, p)
	}
}

// NewFindPageEndpoint returns an endpoint function that calls the method
// "FindPage" of service "epub_index_creator".
func NewFindPageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*FindPagePayload)
		return s.FindPage(ctx, p)
	}
}

// NewCreatePageEndpoint returns an endpoint function that calls the method
// "CreatePage" of service "epub_index_creator".
func NewCreatePageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePagePayload)
		return s.CreatePage(ctx, p)
	}
}

// NewUpdatePageEndpoint returns an endpoint function that calls the method
// "UpdatePage" of service "epub_index_creator".
func NewUpdatePageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePagePayload)
		return s.UpdatePage(ctx, p)
	}
}

// NewDeletePageEndpoint returns an endpoint function that calls the method
// "DeletePage" of service "epub_index_creator".
func NewDeletePageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePagePayload)
		return nil, s.DeletePage(ctx, p)
	}
}

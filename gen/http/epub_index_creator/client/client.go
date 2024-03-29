// Code generated by goa v3.15.0, DO NOT EDIT.
//
// epub_index_creator client HTTP transport
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the epub_index_creator service endpoint HTTP clients.
type Client struct {
	// ListBooks Doer is the HTTP client used to make requests to the ListBooks
	// endpoint.
	ListBooksDoer goahttp.Doer

	// FindBook Doer is the HTTP client used to make requests to the FindBook
	// endpoint.
	FindBookDoer goahttp.Doer

	// CreateBook Doer is the HTTP client used to make requests to the CreateBook
	// endpoint.
	CreateBookDoer goahttp.Doer

	// UpdateBook Doer is the HTTP client used to make requests to the UpdateBook
	// endpoint.
	UpdateBookDoer goahttp.Doer

	// DeleteBook Doer is the HTTP client used to make requests to the DeleteBook
	// endpoint.
	DeleteBookDoer goahttp.Doer

	// ListPages Doer is the HTTP client used to make requests to the ListPages
	// endpoint.
	ListPagesDoer goahttp.Doer

	// FindPage Doer is the HTTP client used to make requests to the FindPage
	// endpoint.
	FindPageDoer goahttp.Doer

	// CreatePage Doer is the HTTP client used to make requests to the CreatePage
	// endpoint.
	CreatePageDoer goahttp.Doer

	// UpdatePage Doer is the HTTP client used to make requests to the UpdatePage
	// endpoint.
	UpdatePageDoer goahttp.Doer

	// DeletePage Doer is the HTTP client used to make requests to the DeletePage
	// endpoint.
	DeletePageDoer goahttp.Doer

	// ListKeywordsInPage Doer is the HTTP client used to make requests to the
	// ListKeywordsInPage endpoint.
	ListKeywordsInPageDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the epub_index_creator service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListBooksDoer:          doer,
		FindBookDoer:           doer,
		CreateBookDoer:         doer,
		UpdateBookDoer:         doer,
		DeleteBookDoer:         doer,
		ListPagesDoer:          doer,
		FindPageDoer:           doer,
		CreatePageDoer:         doer,
		UpdatePageDoer:         doer,
		DeletePageDoer:         doer,
		ListKeywordsInPageDoer: doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// ListBooks returns an endpoint that makes HTTP requests to the
// epub_index_creator service ListBooks server.
func (c *Client) ListBooks() goa.Endpoint {
	var (
		encodeRequest  = EncodeListBooksRequest(c.encoder)
		decodeResponse = DecodeListBooksResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListBooksRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListBooksDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "ListBooks", err)
		}
		return decodeResponse(resp)
	}
}

// FindBook returns an endpoint that makes HTTP requests to the
// epub_index_creator service FindBook server.
func (c *Client) FindBook() goa.Endpoint {
	var (
		decodeResponse = DecodeFindBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildFindBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FindBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "FindBook", err)
		}
		return decodeResponse(resp)
	}
}

// CreateBook returns an endpoint that makes HTTP requests to the
// epub_index_creator service CreateBook server.
func (c *Client) CreateBook() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateBookRequest(c.encoder)
		decodeResponse = DecodeCreateBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreateBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "CreateBook", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateBook returns an endpoint that makes HTTP requests to the
// epub_index_creator service UpdateBook server.
func (c *Client) UpdateBook() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateBookRequest(c.encoder)
		decodeResponse = DecodeUpdateBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "UpdateBook", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteBook returns an endpoint that makes HTTP requests to the
// epub_index_creator service DeleteBook server.
func (c *Client) DeleteBook() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeleteBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "DeleteBook", err)
		}
		return decodeResponse(resp)
	}
}

// ListPages returns an endpoint that makes HTTP requests to the
// epub_index_creator service ListPages server.
func (c *Client) ListPages() goa.Endpoint {
	var (
		encodeRequest  = EncodeListPagesRequest(c.encoder)
		decodeResponse = DecodeListPagesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListPagesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListPagesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "ListPages", err)
		}
		return decodeResponse(resp)
	}
}

// FindPage returns an endpoint that makes HTTP requests to the
// epub_index_creator service FindPage server.
func (c *Client) FindPage() goa.Endpoint {
	var (
		decodeResponse = DecodeFindPageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildFindPageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FindPageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "FindPage", err)
		}
		return decodeResponse(resp)
	}
}

// CreatePage returns an endpoint that makes HTTP requests to the
// epub_index_creator service CreatePage server.
func (c *Client) CreatePage() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreatePageRequest(c.encoder)
		decodeResponse = DecodeCreatePageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreatePageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreatePageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "CreatePage", err)
		}
		return decodeResponse(resp)
	}
}

// UpdatePage returns an endpoint that makes HTTP requests to the
// epub_index_creator service UpdatePage server.
func (c *Client) UpdatePage() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdatePageRequest(c.encoder)
		decodeResponse = DecodeUpdatePageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdatePageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdatePageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "UpdatePage", err)
		}
		return decodeResponse(resp)
	}
}

// DeletePage returns an endpoint that makes HTTP requests to the
// epub_index_creator service DeletePage server.
func (c *Client) DeletePage() goa.Endpoint {
	var (
		decodeResponse = DecodeDeletePageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeletePageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeletePageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "DeletePage", err)
		}
		return decodeResponse(resp)
	}
}

// ListKeywordsInPage returns an endpoint that makes HTTP requests to the
// epub_index_creator service ListKeywordsInPage server.
func (c *Client) ListKeywordsInPage() goa.Endpoint {
	var (
		decodeResponse = DecodeListKeywordsInPageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListKeywordsInPageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListKeywordsInPageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("epub_index_creator", "ListKeywordsInPage", err)
		}
		return decodeResponse(resp)
	}
}

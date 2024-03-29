// Code generated by goa v3.15.0, DO NOT EDIT.
//
// epub_index_creator HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	epubindexcreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListBooksRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "ListBooks" endpoint
func (c *Client) BuildListBooksRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListBooksEpubIndexCreatorPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "ListBooks", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListBooksRequest returns an encoder for requests sent to the
// epub_index_creator ListBooks server.
func EncodeListBooksRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*epubindexcreator.ListBooksPayload)
		if !ok {
			return goahttp.ErrInvalidType("epub_index_creator", "ListBooks", "*epubindexcreator.ListBooksPayload", v)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("offset", fmt.Sprintf("%v", p.Offset))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListBooksResponse returns a decoder for responses returned by the
// epub_index_creator ListBooks endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeListBooksResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListBooksResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "ListBooks", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateBookResponseResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "ListBooks", err)
			}
			res := NewListBooksBookResponseOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "ListBooks", resp.StatusCode, string(body))
		}
	}
}

// BuildFindBookRequest instantiates a HTTP request object with method and path
// set to call the "epub_index_creator" service "FindBook" endpoint
func (c *Client) BuildFindBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn string
	)
	{
		p, ok := v.(*epubindexcreator.FindBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "FindBook", "*epubindexcreator.FindBookPayload", v)
		}
		isbn = string(p.Isbn)
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FindBookEpubIndexCreatorPath(isbn)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "FindBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFindBookResponse returns a decoder for responses returned by the
// epub_index_creator FindBook endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeFindBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FindBookOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "FindBook", err)
			}
			err = ValidateFindBookOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "FindBook", err)
			}
			res := NewFindBookBookResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "FindBook", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateBookRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "CreateBook" endpoint
func (c *Client) BuildCreateBookRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateBookEpubIndexCreatorPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "CreateBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateBookRequest returns an encoder for requests sent to the
// epub_index_creator CreateBook server.
func EncodeCreateBookRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*epubindexcreator.BookRequest)
		if !ok {
			return goahttp.ErrInvalidType("epub_index_creator", "CreateBook", "*epubindexcreator.BookRequest", v)
		}
		body := NewCreateBookRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("epub_index_creator", "CreateBook", err)
		}
		return nil
	}
}

// DecodeCreateBookResponse returns a decoder for responses returned by the
// epub_index_creator CreateBook endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeCreateBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateBookOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "CreateBook", err)
			}
			err = ValidateCreateBookOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "CreateBook", err)
			}
			res := NewCreateBookBookResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "CreateBook", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateBookRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "UpdateBook" endpoint
func (c *Client) BuildUpdateBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn string
	)
	{
		p, ok := v.(*epubindexcreator.BookRequest)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "UpdateBook", "*epubindexcreator.BookRequest", v)
		}
		isbn = string(p.Isbn)
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateBookEpubIndexCreatorPath(isbn)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "UpdateBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateBookRequest returns an encoder for requests sent to the
// epub_index_creator UpdateBook server.
func EncodeUpdateBookRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*epubindexcreator.BookRequest)
		if !ok {
			return goahttp.ErrInvalidType("epub_index_creator", "UpdateBook", "*epubindexcreator.BookRequest", v)
		}
		body := NewUpdateBookRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("epub_index_creator", "UpdateBook", err)
		}
		return nil
	}
}

// DecodeUpdateBookResponse returns a decoder for responses returned by the
// epub_index_creator UpdateBook endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeUpdateBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateBookOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "UpdateBook", err)
			}
			err = ValidateUpdateBookOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "UpdateBook", err)
			}
			res := NewUpdateBookBookResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "UpdateBook", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteBookRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "DeleteBook" endpoint
func (c *Client) BuildDeleteBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn string
	)
	{
		p, ok := v.(*epubindexcreator.DeleteBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "DeleteBook", "*epubindexcreator.DeleteBookPayload", v)
		}
		isbn = string(p.Isbn)
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteBookEpubIndexCreatorPath(isbn)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "DeleteBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteBookResponse returns a decoder for responses returned by the
// epub_index_creator DeleteBook endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeDeleteBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "DeleteBook", resp.StatusCode, string(body))
		}
	}
}

// BuildListPagesRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "ListPages" endpoint
func (c *Client) BuildListPagesRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn string
	)
	{
		p, ok := v.(*epubindexcreator.ListPagesPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "ListPages", "*epubindexcreator.ListPagesPayload", v)
		}
		isbn = string(p.Isbn)
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListPagesEpubIndexCreatorPath(isbn)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "ListPages", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListPagesRequest returns an encoder for requests sent to the
// epub_index_creator ListPages server.
func EncodeListPagesRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*epubindexcreator.ListPagesPayload)
		if !ok {
			return goahttp.ErrInvalidType("epub_index_creator", "ListPages", "*epubindexcreator.ListPagesPayload", v)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("offset", fmt.Sprintf("%v", p.Offset))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListPagesResponse returns a decoder for responses returned by the
// epub_index_creator ListPages endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeListPagesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListPagesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "ListPages", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidatePageResponseResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "ListPages", err)
			}
			res := NewListPagesPageResponseOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "ListPages", resp.StatusCode, string(body))
		}
	}
}

// BuildFindPageRequest instantiates a HTTP request object with method and path
// set to call the "epub_index_creator" service "FindPage" endpoint
func (c *Client) BuildFindPageRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn   string
		pageID int
	)
	{
		p, ok := v.(*epubindexcreator.FindPagePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "FindPage", "*epubindexcreator.FindPagePayload", v)
		}
		isbn = string(p.Isbn)
		pageID = p.PageID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FindPageEpubIndexCreatorPath(isbn, pageID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "FindPage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFindPageResponse returns a decoder for responses returned by the
// epub_index_creator FindPage endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeFindPageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FindPageOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "FindPage", err)
			}
			err = ValidateFindPageOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "FindPage", err)
			}
			res := NewFindPagePageResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "FindPage", resp.StatusCode, string(body))
		}
	}
}

// BuildCreatePageRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "CreatePage" endpoint
func (c *Client) BuildCreatePageRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn   string
		pageID int
	)
	{
		p, ok := v.(*epubindexcreator.CreatePagePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "CreatePage", "*epubindexcreator.CreatePagePayload", v)
		}
		isbn = string(p.Isbn)
		pageID = p.PageID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreatePageEpubIndexCreatorPath(isbn, pageID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "CreatePage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreatePageRequest returns an encoder for requests sent to the
// epub_index_creator CreatePage server.
func EncodeCreatePageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*epubindexcreator.CreatePagePayload)
		if !ok {
			return goahttp.ErrInvalidType("epub_index_creator", "CreatePage", "*epubindexcreator.CreatePagePayload", v)
		}
		body := NewCreatePageRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("epub_index_creator", "CreatePage", err)
		}
		return nil
	}
}

// DecodeCreatePageResponse returns a decoder for responses returned by the
// epub_index_creator CreatePage endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeCreatePageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreatePageOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "CreatePage", err)
			}
			err = ValidateCreatePageOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "CreatePage", err)
			}
			res := NewCreatePagePageResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "CreatePage", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdatePageRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "UpdatePage" endpoint
func (c *Client) BuildUpdatePageRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn   string
		pageID int
	)
	{
		p, ok := v.(*epubindexcreator.UpdatePagePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "UpdatePage", "*epubindexcreator.UpdatePagePayload", v)
		}
		isbn = string(p.Isbn)
		pageID = p.PageID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdatePageEpubIndexCreatorPath(isbn, pageID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "UpdatePage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdatePageRequest returns an encoder for requests sent to the
// epub_index_creator UpdatePage server.
func EncodeUpdatePageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*epubindexcreator.UpdatePagePayload)
		if !ok {
			return goahttp.ErrInvalidType("epub_index_creator", "UpdatePage", "*epubindexcreator.UpdatePagePayload", v)
		}
		body := NewUpdatePageRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("epub_index_creator", "UpdatePage", err)
		}
		return nil
	}
}

// DecodeUpdatePageResponse returns a decoder for responses returned by the
// epub_index_creator UpdatePage endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeUpdatePageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdatePageOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "UpdatePage", err)
			}
			err = ValidateUpdatePageOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "UpdatePage", err)
			}
			res := NewUpdatePagePageResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "UpdatePage", resp.StatusCode, string(body))
		}
	}
}

// BuildDeletePageRequest instantiates a HTTP request object with method and
// path set to call the "epub_index_creator" service "DeletePage" endpoint
func (c *Client) BuildDeletePageRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn   string
		pageID int
	)
	{
		p, ok := v.(*epubindexcreator.DeletePagePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "DeletePage", "*epubindexcreator.DeletePagePayload", v)
		}
		isbn = p.Isbn
		pageID = p.PageID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeletePageEpubIndexCreatorPath(isbn, pageID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "DeletePage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeletePageResponse returns a decoder for responses returned by the
// epub_index_creator DeletePage endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeDeletePageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "DeletePage", resp.StatusCode, string(body))
		}
	}
}

// BuildListKeywordsInPageRequest instantiates a HTTP request object with
// method and path set to call the "epub_index_creator" service
// "ListKeywordsInPage" endpoint
func (c *Client) BuildListKeywordsInPageRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		isbn   string
		pageID int
	)
	{
		p, ok := v.(*epubindexcreator.ListKeywordsInPagePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("epub_index_creator", "ListKeywordsInPage", "*epubindexcreator.ListKeywordsInPagePayload", v)
		}
		isbn = string(p.Isbn)
		pageID = p.PageID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListKeywordsInPageEpubIndexCreatorPath(isbn, pageID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("epub_index_creator", "ListKeywordsInPage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListKeywordsInPageResponse returns a decoder for responses returned by
// the epub_index_creator ListKeywordsInPage endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeListKeywordsInPageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListKeywordsInPageResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("epub_index_creator", "ListKeywordsInPage", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateKeywordResponseResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("epub_index_creator", "ListKeywordsInPage", err)
			}
			res := NewListKeywordsInPageKeywordResponseOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("epub_index_creator", "ListKeywordsInPage", resp.StatusCode, string(body))
		}
	}
}

// unmarshalBookResponseResponseToEpubindexcreatorBookResponse builds a value
// of type *epubindexcreator.BookResponse from a value of type
// *BookResponseResponse.
func unmarshalBookResponseResponseToEpubindexcreatorBookResponse(v *BookResponseResponse) *epubindexcreator.BookResponse {
	res := &epubindexcreator.BookResponse{
		Isbn:      epubindexcreator.ISBN(*v.Isbn),
		Title:     *v.Title,
		Author:    *v.Author,
		Language:  *v.Language,
		Publisher: *v.Publisher,
	}
	res.Pages = make([]*epubindexcreator.PageResponse, len(v.Pages))
	for i, val := range v.Pages {
		res.Pages[i] = unmarshalPageResponseResponseToEpubindexcreatorPageResponse(val)
	}

	return res
}

// unmarshalPageResponseResponseToEpubindexcreatorPageResponse builds a value
// of type *epubindexcreator.PageResponse from a value of type
// *PageResponseResponse.
func unmarshalPageResponseResponseToEpubindexcreatorPageResponse(v *PageResponseResponse) *epubindexcreator.PageResponse {
	res := &epubindexcreator.PageResponse{
		ID:    *v.ID,
		Title: *v.Title,
	}
	res.Keywords = make([]string, len(v.Keywords))
	for i, val := range v.Keywords {
		res.Keywords[i] = val
	}

	return res
}

// unmarshalPageResponseResponseBodyToEpubindexcreatorPageResponse builds a
// value of type *epubindexcreator.PageResponse from a value of type
// *PageResponseResponseBody.
func unmarshalPageResponseResponseBodyToEpubindexcreatorPageResponse(v *PageResponseResponseBody) *epubindexcreator.PageResponse {
	res := &epubindexcreator.PageResponse{
		ID:    *v.ID,
		Title: *v.Title,
	}
	res.Keywords = make([]string, len(v.Keywords))
	for i, val := range v.Keywords {
		res.Keywords[i] = val
	}

	return res
}

// marshalEpubindexcreatorCreatePageRequestToCreatePageRequestRequestBody
// builds a value of type *CreatePageRequestRequestBody from a value of type
// *epubindexcreator.CreatePageRequest.
func marshalEpubindexcreatorCreatePageRequestToCreatePageRequestRequestBody(v *epubindexcreator.CreatePageRequest) *CreatePageRequestRequestBody {
	res := &CreatePageRequestRequestBody{
		Title: v.Title,
	}
	if v.Keywords != nil {
		res.Keywords = make([]string, len(v.Keywords))
		for i, val := range v.Keywords {
			res.Keywords[i] = val
		}
	}

	return res
}

// marshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest
// builds a value of type *epubindexcreator.CreatePageRequest from a value of
// type *CreatePageRequestRequestBody.
func marshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest(v *CreatePageRequestRequestBody) *epubindexcreator.CreatePageRequest {
	res := &epubindexcreator.CreatePageRequest{
		Title: v.Title,
	}
	if v.Keywords != nil {
		res.Keywords = make([]string, len(v.Keywords))
		for i, val := range v.Keywords {
			res.Keywords[i] = val
		}
	}

	return res
}

// unmarshalKeywordResponseResponseToEpubindexcreatorKeywordResponse builds a
// value of type *epubindexcreator.KeywordResponse from a value of type
// *KeywordResponseResponse.
func unmarshalKeywordResponseResponseToEpubindexcreatorKeywordResponse(v *KeywordResponseResponse) *epubindexcreator.KeywordResponse {
	res := &epubindexcreator.KeywordResponse{
		ID:      *v.ID,
		Keyword: *v.Keyword,
	}

	return res
}

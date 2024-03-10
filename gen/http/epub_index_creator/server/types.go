// Code generated by goa v3.15.0, DO NOT EDIT.
//
// epub_index_creator HTTP server types
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package server

import (
	epubindexcreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
	goa "goa.design/goa/v3/pkg"
)

// CreateBookRequestBody is the type of the "epub_index_creator" service
// "CreateBook" endpoint HTTP request body.
type CreateBookRequestBody struct {
	// ISBN of the book
	Isbn *string `form:"isbn,omitempty" json:"isbn,omitempty" xml:"isbn,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Language of the book
	Language *string `form:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// Publisher of the book
	Publisher *string `form:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
}

// UpdateBookRequestBody is the type of the "epub_index_creator" service
// "UpdateBook" endpoint HTTP request body.
type UpdateBookRequestBody struct {
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Language of the book
	Language *string `form:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// Publisher of the book
	Publisher *string `form:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
}

// CreatePageRequestBody is the type of the "epub_index_creator" service
// "CreatePage" endpoint HTTP request body.
type CreatePageRequestBody struct {
	Page *CreatePageRequestRequestBody `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
}

// UpdatePageRequestBody is the type of the "epub_index_creator" service
// "UpdatePage" endpoint HTTP request body.
type UpdatePageRequestBody struct {
	Page *CreatePageRequestRequestBody `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
}

// ListBooksResponseBody is the type of the "epub_index_creator" service
// "ListBooks" endpoint HTTP response body.
type ListBooksResponseBody []*BookResponseResponse

// FindBookOKResponseBody is the type of the "epub_index_creator" service
// "FindBook" endpoint HTTP response body.
type FindBookOKResponseBody struct {
	Isbn string `form:"isbn" json:"isbn" xml:"isbn"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Language of the book
	Language string `form:"language" json:"language" xml:"language"`
	// Publisher of the book
	Publisher string `form:"publisher" json:"publisher" xml:"publisher"`
	// Pages of the book
	Pages []*PageResponseResponseBody `form:"pages" json:"pages" xml:"pages"`
}

// CreateBookOKResponseBody is the type of the "epub_index_creator" service
// "CreateBook" endpoint HTTP response body.
type CreateBookOKResponseBody struct {
	Isbn string `form:"isbn" json:"isbn" xml:"isbn"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Language of the book
	Language string `form:"language" json:"language" xml:"language"`
	// Publisher of the book
	Publisher string `form:"publisher" json:"publisher" xml:"publisher"`
	// Pages of the book
	Pages []*PageResponseResponseBody `form:"pages" json:"pages" xml:"pages"`
}

// UpdateBookOKResponseBody is the type of the "epub_index_creator" service
// "UpdateBook" endpoint HTTP response body.
type UpdateBookOKResponseBody struct {
	Isbn string `form:"isbn" json:"isbn" xml:"isbn"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Language of the book
	Language string `form:"language" json:"language" xml:"language"`
	// Publisher of the book
	Publisher string `form:"publisher" json:"publisher" xml:"publisher"`
	// Pages of the book
	Pages []*PageResponseResponseBody `form:"pages" json:"pages" xml:"pages"`
}

// ListPagesResponseBody is the type of the "epub_index_creator" service
// "ListPages" endpoint HTTP response body.
type ListPagesResponseBody []*PageResponseResponse

// FindPageOKResponseBody is the type of the "epub_index_creator" service
// "FindPage" endpoint HTTP response body.
type FindPageOKResponseBody struct {
	// ID of the page
	ID int `form:"id" json:"id" xml:"id"`
	// Title of the page
	Title string `form:"title" json:"title" xml:"title"`
	// Keywords of the page
	Keywords []string `form:"keywords" json:"keywords" xml:"keywords"`
}

// CreatePageOKResponseBody is the type of the "epub_index_creator" service
// "CreatePage" endpoint HTTP response body.
type CreatePageOKResponseBody struct {
	// ID of the page
	ID int `form:"id" json:"id" xml:"id"`
	// Title of the page
	Title string `form:"title" json:"title" xml:"title"`
	// Keywords of the page
	Keywords []string `form:"keywords" json:"keywords" xml:"keywords"`
}

// UpdatePageOKResponseBody is the type of the "epub_index_creator" service
// "UpdatePage" endpoint HTTP response body.
type UpdatePageOKResponseBody struct {
	// ID of the page
	ID int `form:"id" json:"id" xml:"id"`
	// Title of the page
	Title string `form:"title" json:"title" xml:"title"`
	// Keywords of the page
	Keywords []string `form:"keywords" json:"keywords" xml:"keywords"`
}

// BookResponseResponse is used to define fields on response body types.
type BookResponseResponse struct {
	Isbn string `form:"isbn" json:"isbn" xml:"isbn"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Language of the book
	Language string `form:"language" json:"language" xml:"language"`
	// Publisher of the book
	Publisher string `form:"publisher" json:"publisher" xml:"publisher"`
	// Pages of the book
	Pages []*PageResponseResponse `form:"pages" json:"pages" xml:"pages"`
}

// PageResponseResponse is used to define fields on response body types.
type PageResponseResponse struct {
	// ID of the page
	ID int `form:"id" json:"id" xml:"id"`
	// Title of the page
	Title string `form:"title" json:"title" xml:"title"`
	// Keywords of the page
	Keywords []string `form:"keywords" json:"keywords" xml:"keywords"`
}

// PageResponseResponseBody is used to define fields on response body types.
type PageResponseResponseBody struct {
	// ID of the page
	ID int `form:"id" json:"id" xml:"id"`
	// Title of the page
	Title string `form:"title" json:"title" xml:"title"`
	// Keywords of the page
	Keywords []string `form:"keywords" json:"keywords" xml:"keywords"`
}

// CreatePageRequestRequestBody is used to define fields on request body types.
type CreatePageRequestRequestBody struct {
	// Title of the page
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Keywords of the page
	Keywords []string `form:"keywords,omitempty" json:"keywords,omitempty" xml:"keywords,omitempty"`
}

// NewListBooksResponseBody builds the HTTP response body from the result of
// the "ListBooks" endpoint of the "epub_index_creator" service.
func NewListBooksResponseBody(res []*epubindexcreator.BookResponse) ListBooksResponseBody {
	body := make([]*BookResponseResponse, len(res))
	for i, val := range res {
		body[i] = marshalEpubindexcreatorBookResponseToBookResponseResponse(val)
	}
	return body
}

// NewFindBookOKResponseBody builds the HTTP response body from the result of
// the "FindBook" endpoint of the "epub_index_creator" service.
func NewFindBookOKResponseBody(res *epubindexcreator.BookResponse) *FindBookOKResponseBody {
	body := &FindBookOKResponseBody{
		Isbn:      string(res.Isbn),
		Title:     res.Title,
		Author:    res.Author,
		Language:  res.Language,
		Publisher: res.Publisher,
	}
	if res.Pages != nil {
		body.Pages = make([]*PageResponseResponseBody, len(res.Pages))
		for i, val := range res.Pages {
			body.Pages[i] = marshalEpubindexcreatorPageResponseToPageResponseResponseBody(val)
		}
	} else {
		body.Pages = []*PageResponseResponseBody{}
	}
	return body
}

// NewCreateBookOKResponseBody builds the HTTP response body from the result of
// the "CreateBook" endpoint of the "epub_index_creator" service.
func NewCreateBookOKResponseBody(res *epubindexcreator.BookResponse) *CreateBookOKResponseBody {
	body := &CreateBookOKResponseBody{
		Isbn:      string(res.Isbn),
		Title:     res.Title,
		Author:    res.Author,
		Language:  res.Language,
		Publisher: res.Publisher,
	}
	if res.Pages != nil {
		body.Pages = make([]*PageResponseResponseBody, len(res.Pages))
		for i, val := range res.Pages {
			body.Pages[i] = marshalEpubindexcreatorPageResponseToPageResponseResponseBody(val)
		}
	} else {
		body.Pages = []*PageResponseResponseBody{}
	}
	return body
}

// NewUpdateBookOKResponseBody builds the HTTP response body from the result of
// the "UpdateBook" endpoint of the "epub_index_creator" service.
func NewUpdateBookOKResponseBody(res *epubindexcreator.BookResponse) *UpdateBookOKResponseBody {
	body := &UpdateBookOKResponseBody{
		Isbn:      string(res.Isbn),
		Title:     res.Title,
		Author:    res.Author,
		Language:  res.Language,
		Publisher: res.Publisher,
	}
	if res.Pages != nil {
		body.Pages = make([]*PageResponseResponseBody, len(res.Pages))
		for i, val := range res.Pages {
			body.Pages[i] = marshalEpubindexcreatorPageResponseToPageResponseResponseBody(val)
		}
	} else {
		body.Pages = []*PageResponseResponseBody{}
	}
	return body
}

// NewListPagesResponseBody builds the HTTP response body from the result of
// the "ListPages" endpoint of the "epub_index_creator" service.
func NewListPagesResponseBody(res []*epubindexcreator.PageResponse) ListPagesResponseBody {
	body := make([]*PageResponseResponse, len(res))
	for i, val := range res {
		body[i] = marshalEpubindexcreatorPageResponseToPageResponseResponse(val)
	}
	return body
}

// NewFindPageOKResponseBody builds the HTTP response body from the result of
// the "FindPage" endpoint of the "epub_index_creator" service.
func NewFindPageOKResponseBody(res *epubindexcreator.PageResponse) *FindPageOKResponseBody {
	body := &FindPageOKResponseBody{
		ID:    res.ID,
		Title: res.Title,
	}
	if res.Keywords != nil {
		body.Keywords = make([]string, len(res.Keywords))
		for i, val := range res.Keywords {
			body.Keywords[i] = val
		}
	} else {
		body.Keywords = []string{}
	}
	return body
}

// NewCreatePageOKResponseBody builds the HTTP response body from the result of
// the "CreatePage" endpoint of the "epub_index_creator" service.
func NewCreatePageOKResponseBody(res *epubindexcreator.PageResponse) *CreatePageOKResponseBody {
	body := &CreatePageOKResponseBody{
		ID:    res.ID,
		Title: res.Title,
	}
	if res.Keywords != nil {
		body.Keywords = make([]string, len(res.Keywords))
		for i, val := range res.Keywords {
			body.Keywords[i] = val
		}
	} else {
		body.Keywords = []string{}
	}
	return body
}

// NewUpdatePageOKResponseBody builds the HTTP response body from the result of
// the "UpdatePage" endpoint of the "epub_index_creator" service.
func NewUpdatePageOKResponseBody(res *epubindexcreator.PageResponse) *UpdatePageOKResponseBody {
	body := &UpdatePageOKResponseBody{
		ID:    res.ID,
		Title: res.Title,
	}
	if res.Keywords != nil {
		body.Keywords = make([]string, len(res.Keywords))
		for i, val := range res.Keywords {
			body.Keywords[i] = val
		}
	} else {
		body.Keywords = []string{}
	}
	return body
}

// NewListBooksPayload builds a epub_index_creator service ListBooks endpoint
// payload.
func NewListBooksPayload(limit int, offset int) *epubindexcreator.ListBooksPayload {
	v := &epubindexcreator.ListBooksPayload{}
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewFindBookPayload builds a epub_index_creator service FindBook endpoint
// payload.
func NewFindBookPayload(isbn string) *epubindexcreator.FindBookPayload {
	v := &epubindexcreator.FindBookPayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v
}

// NewCreateBookBookRequest builds a epub_index_creator service CreateBook
// endpoint payload.
func NewCreateBookBookRequest(body *CreateBookRequestBody) *epubindexcreator.BookRequest {
	v := &epubindexcreator.BookRequest{
		Isbn:      epubindexcreator.ISBN(*body.Isbn),
		Title:     *body.Title,
		Author:    *body.Author,
		Language:  *body.Language,
		Publisher: *body.Publisher,
	}

	return v
}

// NewUpdateBookBookRequest builds a epub_index_creator service UpdateBook
// endpoint payload.
func NewUpdateBookBookRequest(body *UpdateBookRequestBody, isbn string) *epubindexcreator.BookRequest {
	v := &epubindexcreator.BookRequest{
		Title:     *body.Title,
		Author:    *body.Author,
		Language:  *body.Language,
		Publisher: *body.Publisher,
	}
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v
}

// NewDeleteBookPayload builds a epub_index_creator service DeleteBook endpoint
// payload.
func NewDeleteBookPayload(isbn string) *epubindexcreator.DeleteBookPayload {
	v := &epubindexcreator.DeleteBookPayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v
}

// NewListPagesPayload builds a epub_index_creator service ListPages endpoint
// payload.
func NewListPagesPayload(isbn string, limit int, offset int) *epubindexcreator.ListPagesPayload {
	v := &epubindexcreator.ListPagesPayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewFindPagePayload builds a epub_index_creator service FindPage endpoint
// payload.
func NewFindPagePayload(isbn string, pageID int) *epubindexcreator.FindPagePayload {
	v := &epubindexcreator.FindPagePayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.PageID = pageID

	return v
}

// NewCreatePagePayload builds a epub_index_creator service CreatePage endpoint
// payload.
func NewCreatePagePayload(body *CreatePageRequestBody, isbn string) *epubindexcreator.CreatePagePayload {
	v := &epubindexcreator.CreatePagePayload{}
	v.Page = unmarshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest(body.Page)
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v
}

// NewUpdatePagePayload builds a epub_index_creator service UpdatePage endpoint
// payload.
func NewUpdatePagePayload(body *UpdatePageRequestBody, isbn string) *epubindexcreator.UpdatePagePayload {
	v := &epubindexcreator.UpdatePagePayload{}
	v.Page = unmarshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest(body.Page)
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v
}

// NewDeletePagePayload builds a epub_index_creator service DeletePage endpoint
// payload.
func NewDeletePagePayload(isbn string, pageID int) *epubindexcreator.DeletePagePayload {
	v := &epubindexcreator.DeletePagePayload{}
	v.Isbn = isbn
	v.PageID = pageID

	return v
}

// ValidateCreateBookRequestBody runs the validations defined on
// CreateBookRequestBody
func ValidateCreateBookRequestBody(body *CreateBookRequestBody) (err error) {
	if body.Isbn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("isbn", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.Language == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("language", "body"))
	}
	if body.Publisher == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("publisher", "body"))
	}
	if body.Isbn != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.isbn", *body.Isbn, "^[0-9]{13}$"))
	}
	return
}

// ValidateUpdateBookRequestBody runs the validations defined on
// UpdateBookRequestBody
func ValidateUpdateBookRequestBody(body *UpdateBookRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.Language == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("language", "body"))
	}
	if body.Publisher == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("publisher", "body"))
	}
	return
}

// ValidateCreatePageRequestBody runs the validations defined on
// CreatePageRequestBody
func ValidateCreatePageRequestBody(body *CreatePageRequestBody) (err error) {
	if body.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
	}
	if body.Page != nil {
		if err2 := ValidateCreatePageRequestRequestBody(body.Page); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdatePageRequestBody runs the validations defined on
// UpdatePageRequestBody
func ValidateUpdatePageRequestBody(body *UpdatePageRequestBody) (err error) {
	if body.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
	}
	if body.Page != nil {
		if err2 := ValidateCreatePageRequestRequestBody(body.Page); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCreatePageRequestRequestBody runs the validations defined on
// CreatePageRequestRequestBody
func ValidateCreatePageRequestRequestBody(body *CreatePageRequestRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	return
}

// Code generated by goa v3.15.0, DO NOT EDIT.
//
// epub_index_creator HTTP client CLI support package
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	epubindexcreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
	goa "goa.design/goa/v3/pkg"
)

// BuildListBooksPayload builds the payload for the epub_index_creator
// ListBooks endpoint from CLI flags.
func BuildListBooksPayload(epubIndexCreatorListBooksLimit string, epubIndexCreatorListBooksOffset string) (*epubindexcreator.ListBooksPayload, error) {
	var err error
	var limit int
	{
		if epubIndexCreatorListBooksLimit != "" {
			var v int64
			v, err = strconv.ParseInt(epubIndexCreatorListBooksLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var offset int
	{
		if epubIndexCreatorListBooksOffset != "" {
			var v int64
			v, err = strconv.ParseInt(epubIndexCreatorListBooksOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &epubindexcreator.ListBooksPayload{}
	v.Limit = limit
	v.Offset = offset

	return v, nil
}

// BuildFindBookPayload builds the payload for the epub_index_creator FindBook
// endpoint from CLI flags.
func BuildFindBookPayload(epubIndexCreatorFindBookIsbn string) (*epubindexcreator.FindBookPayload, error) {
	var err error
	var isbn string
	{
		isbn = epubIndexCreatorFindBookIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	v := &epubindexcreator.FindBookPayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v, nil
}

// BuildCreateBookPayload builds the payload for the epub_index_creator
// CreateBook endpoint from CLI flags.
func BuildCreateBookPayload(epubIndexCreatorCreateBookBody string) (*epubindexcreator.BookRequest, error) {
	var err error
	var body CreateBookRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorCreateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"Nam quibusdam vero rem aliquam voluptatibus.\",\n      \"isbn\": \"1554585076269\",\n      \"language\": \"Eligendi ipsa porro.\",\n      \"publisher\": \"Qui in omnis quaerat odit.\",\n      \"title\": \"Eligendi placeat.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.isbn", body.Isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	v := &epubindexcreator.BookRequest{
		Isbn:      epubindexcreator.ISBN(body.Isbn),
		Title:     body.Title,
		Author:    body.Author,
		Language:  body.Language,
		Publisher: body.Publisher,
	}

	return v, nil
}

// BuildUpdateBookPayload builds the payload for the epub_index_creator
// UpdateBook endpoint from CLI flags.
func BuildUpdateBookPayload(epubIndexCreatorUpdateBookBody string, epubIndexCreatorUpdateBookIsbn string) (*epubindexcreator.BookRequest, error) {
	var err error
	var body UpdateBookRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorUpdateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"Aut incidunt adipisci consectetur.\",\n      \"language\": \"Repudiandae dolorem est eligendi sit velit.\",\n      \"publisher\": \"Reiciendis est est ut nihil.\",\n      \"title\": \"Repudiandae beatae et non consequatur dolore ad.\"\n   }'")
		}
	}
	var isbn string
	{
		isbn = epubIndexCreatorUpdateBookIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	v := &epubindexcreator.BookRequest{
		Title:     body.Title,
		Author:    body.Author,
		Language:  body.Language,
		Publisher: body.Publisher,
	}
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v, nil
}

// BuildDeleteBookPayload builds the payload for the epub_index_creator
// DeleteBook endpoint from CLI flags.
func BuildDeleteBookPayload(epubIndexCreatorDeleteBookIsbn string) (*epubindexcreator.DeleteBookPayload, error) {
	var err error
	var isbn string
	{
		isbn = epubIndexCreatorDeleteBookIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	v := &epubindexcreator.DeleteBookPayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)

	return v, nil
}

// BuildListPagesPayload builds the payload for the epub_index_creator
// ListPages endpoint from CLI flags.
func BuildListPagesPayload(epubIndexCreatorListPagesIsbn string, epubIndexCreatorListPagesLimit string, epubIndexCreatorListPagesOffset string) (*epubindexcreator.ListPagesPayload, error) {
	var err error
	var isbn string
	{
		isbn = epubIndexCreatorListPagesIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	var limit int
	{
		if epubIndexCreatorListPagesLimit != "" {
			var v int64
			v, err = strconv.ParseInt(epubIndexCreatorListPagesLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var offset int
	{
		if epubIndexCreatorListPagesOffset != "" {
			var v int64
			v, err = strconv.ParseInt(epubIndexCreatorListPagesOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &epubindexcreator.ListPagesPayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.Limit = limit
	v.Offset = offset

	return v, nil
}

// BuildFindPagePayload builds the payload for the epub_index_creator FindPage
// endpoint from CLI flags.
func BuildFindPagePayload(epubIndexCreatorFindPageIsbn string, epubIndexCreatorFindPagePageID string) (*epubindexcreator.FindPagePayload, error) {
	var err error
	var isbn string
	{
		isbn = epubIndexCreatorFindPageIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	var pageID int
	{
		var v int64
		v, err = strconv.ParseInt(epubIndexCreatorFindPagePageID, 10, strconv.IntSize)
		pageID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for pageID, must be INT")
		}
	}
	v := &epubindexcreator.FindPagePayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.PageID = pageID

	return v, nil
}

// BuildCreatePagePayload builds the payload for the epub_index_creator
// CreatePage endpoint from CLI flags.
func BuildCreatePagePayload(epubIndexCreatorCreatePageBody string, epubIndexCreatorCreatePageIsbn string, epubIndexCreatorCreatePagePageID string) (*epubindexcreator.CreatePagePayload, error) {
	var err error
	var body CreatePageRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorCreatePageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"page\": {\n         \"keywords\": [\n            \"Introduction\",\n            \"Chapter 1\",\n            \"Chapter 2\"\n         ],\n         \"title\": \"Introduction\"\n      }\n   }'")
		}
		if body.Page == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var isbn string
	{
		isbn = epubIndexCreatorCreatePageIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	var pageID int
	{
		var v int64
		v, err = strconv.ParseInt(epubIndexCreatorCreatePagePageID, 10, strconv.IntSize)
		pageID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for pageID, must be INT")
		}
	}
	v := &epubindexcreator.CreatePagePayload{}
	if body.Page != nil {
		v.Page = marshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest(body.Page)
	}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.PageID = pageID

	return v, nil
}

// BuildUpdatePagePayload builds the payload for the epub_index_creator
// UpdatePage endpoint from CLI flags.
func BuildUpdatePagePayload(epubIndexCreatorUpdatePageBody string, epubIndexCreatorUpdatePageIsbn string, epubIndexCreatorUpdatePagePageID string) (*epubindexcreator.UpdatePagePayload, error) {
	var err error
	var body UpdatePageRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorUpdatePageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"page\": {\n         \"keywords\": [\n            \"Introduction\",\n            \"Chapter 1\",\n            \"Chapter 2\"\n         ],\n         \"title\": \"Introduction\"\n      }\n   }'")
		}
		if body.Page == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var isbn string
	{
		isbn = epubIndexCreatorUpdatePageIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	var pageID int
	{
		var v int64
		v, err = strconv.ParseInt(epubIndexCreatorUpdatePagePageID, 10, strconv.IntSize)
		pageID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for pageID, must be INT")
		}
	}
	v := &epubindexcreator.UpdatePagePayload{}
	if body.Page != nil {
		v.Page = marshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest(body.Page)
	}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.PageID = pageID

	return v, nil
}

// BuildDeletePagePayload builds the payload for the epub_index_creator
// DeletePage endpoint from CLI flags.
func BuildDeletePagePayload(epubIndexCreatorDeletePageIsbn string, epubIndexCreatorDeletePagePageID string) (*epubindexcreator.DeletePagePayload, error) {
	var err error
	var isbn string
	{
		isbn = epubIndexCreatorDeletePageIsbn
	}
	var pageID int
	{
		var v int64
		v, err = strconv.ParseInt(epubIndexCreatorDeletePagePageID, 10, strconv.IntSize)
		pageID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for pageID, must be INT")
		}
	}
	v := &epubindexcreator.DeletePagePayload{}
	v.Isbn = isbn
	v.PageID = pageID

	return v, nil
}

// BuildListKeywordsInPagePayload builds the payload for the epub_index_creator
// ListKeywordsInPage endpoint from CLI flags.
func BuildListKeywordsInPagePayload(epubIndexCreatorListKeywordsInPageIsbn string, epubIndexCreatorListKeywordsInPagePageID string) (*epubindexcreator.ListKeywordsInPagePayload, error) {
	var err error
	var isbn string
	{
		isbn = epubIndexCreatorListKeywordsInPageIsbn
		err = goa.MergeErrors(err, goa.ValidatePattern("isbn", isbn, "^[0-9]{13}$"))
		if err != nil {
			return nil, err
		}
	}
	var pageID int
	{
		var v int64
		v, err = strconv.ParseInt(epubIndexCreatorListKeywordsInPagePageID, 10, strconv.IntSize)
		pageID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for pageID, must be INT")
		}
	}
	v := &epubindexcreator.ListKeywordsInPagePayload{}
	v.Isbn = epubindexcreator.ISBN(isbn)
	v.PageID = pageID

	return v, nil
}

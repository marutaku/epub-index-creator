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

	epubindexcreator "github.com/marutaku/epub-index-creator/gen/epub_index_creator"
	goa "goa.design/goa/v3/pkg"
)

// BuildListBooksPayload builds the payload for the epub_index_creator
// ListBooks endpoint from CLI flags.
func BuildListBooksPayload(epubIndexCreatorListBooksBody string) (*epubindexcreator.ListBooksPayload, error) {
	var err error
	var body ListBooksRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorListBooksBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"limit\": 47,\n      \"offset\": 8534742228231063417\n   }'")
		}
	}
	v := &epubindexcreator.ListBooksPayload{
		Limit:  body.Limit,
		Offset: body.Offset,
	}
	{
		var zero int
		if v.Limit == zero {
			v.Limit = 100
		}
	}
	{
		var zero int
		if v.Offset == zero {
			v.Offset = 0
		}
	}

	return v, nil
}

// BuildFindBookPayload builds the payload for the epub_index_creator FindBook
// endpoint from CLI flags.
func BuildFindBookPayload(epubIndexCreatorFindBookIsbn string) (*epubindexcreator.FindBookPayload, error) {
	var isbn string
	{
		isbn = epubIndexCreatorFindBookIsbn
	}
	v := &epubindexcreator.FindBookPayload{}
	v.Isbn = isbn

	return v, nil
}

// BuildCreateBookPayload builds the payload for the epub_index_creator
// CreateBook endpoint from CLI flags.
func BuildCreateBookPayload(epubIndexCreatorCreateBookBody string) (*epubindexcreator.Book, error) {
	var err error
	var body CreateBookRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorCreateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"Magni tempora ducimus omnis qui.\",\n      \"isbn\": \"Ex asperiores aspernatur enim.\",\n      \"language\": \"Natus voluptatum voluptatem corporis.\",\n      \"pages\": [\n         {\n            \"keywords\": [\n               {\n                  \"keyword\": \"Temporibus itaque nostrum.\"\n               },\n               {\n                  \"keyword\": \"Temporibus itaque nostrum.\"\n               },\n               {\n                  \"keyword\": \"Temporibus itaque nostrum.\"\n               }\n            ],\n            \"title\": \"Nemo eaque aut et exercitationem placeat ad.\"\n         },\n         {\n            \"keywords\": [\n               {\n                  \"keyword\": \"Temporibus itaque nostrum.\"\n               },\n               {\n                  \"keyword\": \"Temporibus itaque nostrum.\"\n               },\n               {\n                  \"keyword\": \"Temporibus itaque nostrum.\"\n               }\n            ],\n            \"title\": \"Nemo eaque aut et exercitationem placeat ad.\"\n         }\n      ],\n      \"publisher\": \"Dignissimos sed qui et consequatur aspernatur.\",\n      \"title\": \"Laboriosam omnis tenetur iusto animi mollitia.\"\n   }'")
		}
		if body.Pages == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("pages", "body"))
		}
		for _, e := range body.Pages {
			if e != nil {
				if err2 := ValidatePageRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &epubindexcreator.Book{
		Isbn:      body.Isbn,
		Title:     body.Title,
		Author:    body.Author,
		Language:  body.Language,
		Publisher: body.Publisher,
	}
	if body.Pages != nil {
		v.Pages = make([]*epubindexcreator.Page, len(body.Pages))
		for i, val := range body.Pages {
			v.Pages[i] = marshalPageRequestBodyToEpubindexcreatorPage(val)
		}
	} else {
		v.Pages = []*epubindexcreator.Page{}
	}

	return v, nil
}

// BuildUpdateBookPayload builds the payload for the epub_index_creator
// UpdateBook endpoint from CLI flags.
func BuildUpdateBookPayload(epubIndexCreatorUpdateBookBody string, epubIndexCreatorUpdateBookIsbn string) (*epubindexcreator.UpdateBookPayload, error) {
	var err error
	var body UpdateBookRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorUpdateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"Magnam voluptatem occaecati possimus non facere.\",\n      \"language\": \"Non sit et eum officia doloribus.\",\n      \"publisher\": \"Ipsum deserunt incidunt non qui non.\",\n      \"title\": \"Suscipit facilis quasi qui ab.\"\n   }'")
		}
	}
	var isbn string
	{
		isbn = epubIndexCreatorUpdateBookIsbn
	}
	v := &epubindexcreator.UpdateBookPayload{
		Title:     body.Title,
		Author:    body.Author,
		Language:  body.Language,
		Publisher: body.Publisher,
	}
	v.Isbn = isbn

	return v, nil
}

// BuildDeleteBookPayload builds the payload for the epub_index_creator
// DeleteBook endpoint from CLI flags.
func BuildDeleteBookPayload(epubIndexCreatorDeleteBookBody string, epubIndexCreatorDeleteBookIsbn string) (*epubindexcreator.DeleteBookPayload, error) {
	var err error
	var body DeleteBookRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorDeleteBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"book\": {\n         \"author\": \"Atque deserunt sit dolores aspernatur quis error.\",\n         \"isbn\": \"Libero consectetur sunt.\",\n         \"language\": \"Earum ea nihil delectus repellat architecto.\",\n         \"pages\": [\n            {\n               \"keywords\": [\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  },\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  },\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  }\n               ],\n               \"title\": \"Nemo eaque aut et exercitationem placeat ad.\"\n            },\n            {\n               \"keywords\": [\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  },\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  },\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  }\n               ],\n               \"title\": \"Nemo eaque aut et exercitationem placeat ad.\"\n            },\n            {\n               \"keywords\": [\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  },\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  },\n                  {\n                     \"keyword\": \"Temporibus itaque nostrum.\"\n                  }\n               ],\n               \"title\": \"Nemo eaque aut et exercitationem placeat ad.\"\n            }\n         ],\n         \"publisher\": \"Totam officiis quibusdam nobis nemo iusto.\",\n         \"title\": \"Nesciunt nisi in enim.\"\n      }\n   }'")
		}
		if body.Book == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("book", "body"))
		}
		if body.Book != nil {
			if err2 := ValidateBookRequestBody(body.Book); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var isbn string
	{
		isbn = epubIndexCreatorDeleteBookIsbn
	}
	v := &epubindexcreator.DeleteBookPayload{}
	if body.Book != nil {
		v.Book = marshalBookRequestBodyToEpubindexcreatorBook(body.Book)
	}
	v.Isbn = isbn

	return v, nil
}

// BuildCreatePagePayload builds the payload for the epub_index_creator
// CreatePage endpoint from CLI flags.
func BuildCreatePagePayload(epubIndexCreatorCreatePageBody string, epubIndexCreatorCreatePageIsbn string) (*epubindexcreator.CreatePagePayload, error) {
	var err error
	var body CreatePageRequestBody
	{
		err = json.Unmarshal([]byte(epubIndexCreatorCreatePageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"page\": {\n         \"keywords\": [\n            \"Et culpa.\",\n            \"Quia blanditiis dolor sint aut quia.\",\n            \"Eveniet doloremque consequatur dolores.\"\n         ],\n         \"title\": \"Dolor quis sed laborum quam aut reprehenderit.\"\n      }\n   }'")
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
	}
	v := &epubindexcreator.CreatePagePayload{}
	if body.Page != nil {
		v.Page = marshalCreatePageRequestRequestBodyToEpubindexcreatorCreatePageRequest(body.Page)
	}
	v.Isbn = isbn

	return v, nil
}

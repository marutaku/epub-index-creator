package design

import (
	. "goa.design/goa/v3/dsl"
)

var ISBN = Type("ISBN", String, func() {
	Description("ISBN of the book")
	Pattern("^[0-9]{13}$")
	Example("9783161484100")
})

var BookRequest = Type("BookRequest", func() {
	Attribute("isbn", ISBN, "ISBN of the book")
	Attribute("title", String, "Title of the book")
	Attribute("author", String, "Author of the book")
	Attribute("language", String, "Language of the book")
	Attribute("publisher", String, "Publisher of the book")
	Required("isbn", "title", "author", "language", "publisher")
})

var BookResponse = Type("BookResponse", func() {
	Field(1, "isbn", ISBN)
	Field(2, "title", String, func() {
		Description("Title of the book")
		Example("The Go Programming Language")
	})
	Field(3, "author", String, func() {
		Description("Author of the book")
		Example("Alan A. A. Donovan, Brian W. Kernighan")
	})
	Field(4, "language", String, func() {
		Description("Language of the book")
		Example("English")
	})
	Field(5, "publisher", String, func() {
		Description("Publisher of the book")
		Example("Addison-Wesley")
	})
	Field(6, "pages", ArrayOf(PageResponse), func() {
		Description("Pages of the book")
	})
	Required("isbn", "title", "author", "language", "publisher", "pages")
})

var PageRequest = Type("CreatePageRequest", func() {
	Field(1, "title", String, func() {
		Description("Title of the page")
		Example("Introduction")
	})
	Field(2, "keywords", ArrayOf(String), func() {
		Description("Keywords of the page")
		Example([]string{"Introduction", "Chapter 1", "Chapter 2"})
	})
	Required("title")
})

var PageResponse = Type("PageResponse", func() {
	Field(1, "title", String, func() {
		Description("Title of the page")
		Example("Introduction")
	})
	Field(2, "keywords", ArrayOf(String), func() {
		Description("Keywords of the page")
		Example([]string{"Introduction", "Chapter 1", "Chapter 2"})
	})
	Required("title", "keywords")
})

var KeywordResponse = Type("Keyword", func() {
	Field(1, "keyword", String, func() {
		Description("Keyword of the page")
		Example("Introduction")
	})
	Required("keyword")
})

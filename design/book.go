package design

import (
	. "goa.design/goa/v3/dsl"
)

var Book = Type("Book", func() {
	Field(1, "isbn", String, "ISBN of the book")
	Field(2, "title", String, "Title of the book")
	Field(3, "author", String, "Author of the book")
	Field(4, "pages", ArrayOf(Page), "Pages of the book")
	Required("isbn", "title", "author", "pages")
})

var Page = Type("Page", func() {
	Field(1, "title", String, "Title of the page")
	Field(2, "keywords", ArrayOf(Keyword), "Keywords of the page")
	Required("title", "keywords")
})

var Keyword = Type("Keyword", func() {
	Field(1, "keyword", String, "Keyword of the page")
	Required("keyword")
})

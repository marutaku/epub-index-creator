package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("epub-index-creator", func() {
	Title("Epup Index Creator")
	Server("web", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("epub_index_creator", func() {
	Method("ListBooks", func() {
		Payload(func() {
			Attribute("limit", Int, "Maximum number of books to return", func() {
				Minimum(1)
				Maximum(100)
				Default(100)
			})
			Attribute("offset", Int, "Field to paginate books", func() {
				Minimum(0)
				Default(0)
			})
		})
		Result(ArrayOf(Book))

		HTTP(func() {
			GET("/books")
		})
	})
	Method("FindBook", func() {
		Payload(func() {
			Attribute("isbn", String, "ISBN of the book")
			Required("isbn")
		})

		Result(Book)

		HTTP(func() {
			GET("/books/{isbn}")
		})
	})

	Method("CreateBook", func() {
		Payload(Book)

		Result(Book)

		HTTP(func() {
			POST("/books")
		})
	})

	Method("UpdateBook", func() {
		Payload(func() {
			Attribute("isbn", String, "ISBN of the book")
			Attribute("title", String, "Title of the book")
			Attribute("author", String, "Author of the book")
			Attribute("publisher", String, "Publisher of the book")
			Required("isbn", "title", "author", "publisher")
		})
		Result(Book)
		HTTP(func() {
			PUT("/books/{isbn}")
		})
	})

	Method("DeleteBook", func() {
		Payload(func() {
			Attribute("isbn", String, "ISBN of the book")
			Attribute("book", Book)
			Required("isbn", "book")
		})
		Result(Empty)
		HTTP(func() {
			DELETE("/books/{isbn}")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

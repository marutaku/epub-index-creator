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
		Result(ArrayOf(BookResponse))

		HTTP(func() {
			GET("/books")
			Param("limit")
			Param("offset")
		})
	})
	Method("FindBook", func() {
		Payload(func() {
			Attribute("isbn", ISBN)
			Required("isbn")
		})

		Result(BookResponse)

		HTTP(func() {
			GET("/books/{isbn}")
			Response(StatusOK)
			Response(StatusNotFound)
		})
	})

	Method("CreateBook", func() {
		Payload(BookRequest)

		Result(BookResponse)

		HTTP(func() {
			POST("/books")
			Response(StatusOK)
			Response(StatusBadRequest)
		})
	})

	Method("UpdateBook", func() {
		Payload(BookRequest)
		Result(BookResponse)
		HTTP(func() {
			PUT("/books/{isbn}")
			Response(StatusOK)
			Response(StatusNotFound)
		})
	})

	Method("DeleteBook", func() {
		Payload(func() {
			Attribute("isbn", ISBN)
			Required("isbn")
		})
		Result(Empty)
		HTTP(func() {
			DELETE("/books/{isbn}")
			Response(StatusOK)
			Response(StatusNotFound)
		})
	})

	Method("CreatePage", func() {
		Payload(func() {
			Attribute("isbn", ISBN)
			Attribute("page", PageRequest)
			Required("isbn", "page")
		})
		Result(PageResponse)
		HTTP(func() {
			POST("/books/{isbn}/pages")
			Response(StatusOK)
			Response(StatusNotFound)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

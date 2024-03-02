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
	Method("List", func() {
		Payload(func() {
			Attribute("isbn", String, "ISBN of the book")
			Required("isbn")
		})

		Result(Book)

		HTTP(func() {
			GET("/books/{isbn}")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("epub_index_creator", func() {
	Description("The calc service performs operations on numbers.")

	Method("List", func() {
		Payload(func() {
			Attribute("isbn", String, "ISBN of the book")
			Required("isbn")
		})

		Result(Int)

		HTTP(func() {
			GET("/books/{isbn}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

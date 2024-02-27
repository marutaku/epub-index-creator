package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("isbn").NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("language").NotEmpty(),
		field.String("author").NotEmpty(),
		field.String("publisher").NotEmpty(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Keyword.Type),
	}
}

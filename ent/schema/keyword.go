package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Keyword holds the schema definition for the Keyword entity.
type Keyword struct {
	ent.Schema
}

// Fields of the Keyword.
func (Keyword) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").NotEmpty(),
	}
}

// Edges of the Keyword.
func (Keyword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("page", Page.Type).Ref("keywords").Unique(),
	}
}

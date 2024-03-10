// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/marutaku/epub-index-creator/ent/book"
	"github.com/marutaku/epub-index-creator/ent/page"
)

// Page is the model entity for the Page schema.
type Page struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PageQuery when eager-loading is set.
	Edges        PageEdges `json:"edges"`
	book_pages   *int
	selectValues sql.SelectValues
}

// PageEdges holds the relations/edges for other nodes in the graph.
type PageEdges struct {
	// Book holds the value of the book edge.
	Book *Book `json:"book,omitempty"`
	// Keywords holds the value of the keywords edge.
	Keywords []*Keyword `json:"keywords,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BookOrErr returns the Book value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PageEdges) BookOrErr() (*Book, error) {
	if e.loadedTypes[0] {
		if e.Book == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: book.Label}
		}
		return e.Book, nil
	}
	return nil, &NotLoadedError{edge: "book"}
}

// KeywordsOrErr returns the Keywords value or an error if the edge
// was not loaded in eager-loading.
func (e PageEdges) KeywordsOrErr() ([]*Keyword, error) {
	if e.loadedTypes[1] {
		return e.Keywords, nil
	}
	return nil, &NotLoadedError{edge: "keywords"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Page) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			values[i] = new(sql.NullInt64)
		case page.FieldTitle, page.FieldPath:
			values[i] = new(sql.NullString)
		case page.ForeignKeys[0]: // book_pages
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Page fields.
func (pa *Page) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case page.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pa.Title = value.String
			}
		case page.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				pa.Path = value.String
			}
		case page.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field book_pages", value)
			} else if value.Valid {
				pa.book_pages = new(int)
				*pa.book_pages = int(value.Int64)
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Page.
// This includes values selected through modifiers, order, etc.
func (pa *Page) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryBook queries the "book" edge of the Page entity.
func (pa *Page) QueryBook() *BookQuery {
	return NewPageClient(pa.config).QueryBook(pa)
}

// QueryKeywords queries the "keywords" edge of the Page entity.
func (pa *Page) QueryKeywords() *KeywordQuery {
	return NewPageClient(pa.config).QueryKeywords(pa)
}

// Update returns a builder for updating this Page.
// Note that you need to call Page.Unwrap() before calling this method if this Page
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Page) Update() *PageUpdateOne {
	return NewPageClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Page entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Page) Unwrap() *Page {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Page is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Page) String() string {
	var builder strings.Builder
	builder.WriteString("Page(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("title=")
	builder.WriteString(pa.Title)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(pa.Path)
	builder.WriteByte(')')
	return builder.String()
}

// Pages is a parsable slice of Page.
type Pages []*Page

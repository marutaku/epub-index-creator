// Code generated by ent, DO NOT EDIT.

package keyword

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the keyword type in the database.
	Label = "keyword"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKeyword holds the string denoting the keyword field in the database.
	FieldKeyword = "keyword"
	// Table holds the table name of the keyword in the database.
	Table = "keywords"
)

// Columns holds all SQL columns for keyword fields.
var Columns = []string{
	FieldID,
	FieldKeyword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "keywords"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"book_cars",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	KeywordValidator func(string) error
)

// OrderOption defines the ordering options for the Keyword queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKeyword orders the results by the keyword field.
func ByKeyword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyword, opts...).ToFunc()
}

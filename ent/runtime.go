// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/marutaku/epub-index-creator/ent/book"
	"github.com/marutaku/epub-index-creator/ent/keyword"
	"github.com/marutaku/epub-index-creator/ent/page"
	"github.com/marutaku/epub-index-creator/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescIsbn is the schema descriptor for isbn field.
	bookDescIsbn := bookFields[0].Descriptor()
	// book.IsbnValidator is a validator for the "isbn" field. It is called by the builders before save.
	book.IsbnValidator = bookDescIsbn.Validators[0].(func(string) error)
	// bookDescTitle is the schema descriptor for title field.
	bookDescTitle := bookFields[1].Descriptor()
	// book.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	book.TitleValidator = bookDescTitle.Validators[0].(func(string) error)
	// bookDescLanguage is the schema descriptor for language field.
	bookDescLanguage := bookFields[2].Descriptor()
	// book.LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	book.LanguageValidator = bookDescLanguage.Validators[0].(func(string) error)
	// bookDescAuthor is the schema descriptor for author field.
	bookDescAuthor := bookFields[3].Descriptor()
	// book.AuthorValidator is a validator for the "author" field. It is called by the builders before save.
	book.AuthorValidator = bookDescAuthor.Validators[0].(func(string) error)
	// bookDescPublisher is the schema descriptor for publisher field.
	bookDescPublisher := bookFields[4].Descriptor()
	// book.PublisherValidator is a validator for the "publisher" field. It is called by the builders before save.
	book.PublisherValidator = bookDescPublisher.Validators[0].(func(string) error)
	keywordFields := schema.Keyword{}.Fields()
	_ = keywordFields
	// keywordDescKeyword is the schema descriptor for keyword field.
	keywordDescKeyword := keywordFields[1].Descriptor()
	// keyword.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	keyword.KeywordValidator = keywordDescKeyword.Validators[0].(func(string) error)
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescTitle is the schema descriptor for title field.
	pageDescTitle := pageFields[1].Descriptor()
	// page.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	page.TitleValidator = pageDescTitle.Validators[0].(func(string) error)
	// pageDescPath is the schema descriptor for path field.
	pageDescPath := pageFields[2].Descriptor()
	// page.PathValidator is a validator for the "path" field. It is called by the builders before save.
	page.PathValidator = pageDescPath.Validators[0].(func(string) error)
}

package book

import (
	"testing"
)

func TestNewBookFromOPF(t *testing.T) {
	metadata, err := NewBookFromOPF("testdata/testdata.opf")
	if err != nil {
		t.Errorf("Error parsing testdata.xml: %v", err)
	}
	if metadata.ISBN != "567812345678" {
		t.Errorf("ISBN is not correct")
	}
	if metadata.Title != "DummyTitle" {
		t.Errorf("Title is not correct")
	}
	if metadata.Language != "en" {
		t.Errorf("Language is not correct")
	}
	if metadata.Author != "DummyAuthor" {
		t.Errorf("Author is not correct")
	}
	if metadata.Publisher != "DummyPublisher" {
		t.Errorf("Publisher is not correct")
	}
	if len(metadata.Pages) != 3 {
		t.Errorf("Pages is not correct")
	}
}

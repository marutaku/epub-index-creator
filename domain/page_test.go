package domain

import (
	"fmt"
	"testing"
)

var TEST_DATA_PATH = "testdata/testpage.xhtml"

func TestNewPage(t *testing.T) {
	page := NewPage(1, TEST_DATA_PATH)
	if page.Id != 1 {
		t.Errorf("Id is not correct")
	}
	if page.Path != TEST_DATA_PATH {
		t.Errorf("Path is not correct")
	}
}

func TestContent(t *testing.T) {
	page := NewPage(1, TEST_DATA_PATH)
	content, err := page.Content()
	if err != nil {
		t.Errorf("Error reading testdata.html: %v", err)
	}
	if content != "<p>本文</p>" {
		fmt.Println(content)
		t.Errorf("Content is not correct")
	}
}

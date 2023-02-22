package internal

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func createTestData(t *testing.T) string {
	dir, err := os.MkdirTemp("", "hype-test-content")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(dir)

	testMarkdown1 := `
---
title: "Test Post 1"
author:
- name: "John Smith"
  email: "john@example.com"
  bio: "A writer and blogger based in New York City."
  url: "https://example.com/john-smith"
---
# Test Post 1

This is the first test post.
`
	// create temporary markdown files in the test directory
	err = os.WriteFile(filepath.Join(dir, "test1.md"), []byte(testMarkdown1), 0666)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	testMarkdown2 := `
---
title: "Test Post 1"
author:
- name: "John Smith"
  email: "john@example.com"
  bio: "A writer and blogger based in New York City."
  url: "https://example.com/john-smith"
---
# Test Post 1

This is the first test post.
`
	err = os.WriteFile(filepath.Join(dir, "test2.md"), []byte(testMarkdown2), 0666)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	return dir
}

func TestLoadAllContent(t *testing.T) {
	dir := createTestData(t)

	// load content from the test directory
	contents, err := LoadAllContent(dir)
	if err != nil {
		t.Fatalf("failed to load content: %v", err)
	}

	// ensure the correct number of content items were loaded
	if len(contents) != 2 {
		t.Errorf("expected 2 content items, got %d", len(contents))
	}
}

func TestLoadContent(t *testing.T) {
	tests := []struct {
		id     string
		exists bool
	}{
		{"example-post", true},
		{"invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			content, err := LoadContent(tt.id)
			if tt.exists && err != nil {
				t.Errorf("LoadContent() error: %s", err)
			} else if !tt.exists && err == nil {
				t.Error("LoadContent() should have returned an error, but did not")
			}
			if tt.exists && !reflect.DeepEqual(content.Metadata["title"], "Example Post") {
				t.Errorf("LoadContent() returned incorrect metadata: %v", content.Metadata)
			}
		})
	}
}

func TestSplitContent(t *testing.T) {
	data := []byte("title: Example Post\nauthor: me@example.com\n\nThis is the content.")
	meta, content := splitContent(data)

	if len(meta) != 2 {
		t.Errorf("splitContent() returned %d metadata items, expected 2", len(meta))
	}

	if string(content) != "This is the content." {
		t.Errorf("splitContent() returned content '%s', expected 'This is the content.'", string(content))
	}

	if meta["title"] != "Example Post" {
		t.Errorf("splitContent() returned title '%s', expected 'Example Post'", meta["title"])
	}

	if meta["author"] != "me@example.com" {
		t.Errorf("splitContent() returned author '%s', expected 'me@example.com'", meta["author"])
	}
}

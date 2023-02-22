package internal

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfgData := []byte(`
site:
  title: Example Site
  url: https://example.com
author:
  - name: John Doe
    email: john@example.com
    url: https://example.com/john
    bio: Web developer and tech enthusiast
  - name: Jane Smith
    email: jane@example.com
    url: https://example.com/jane
    bio: Designer and photographer
`)

	expectedConfig := &Config{
		Site: struct {
			Title string "yaml:\"title\""
			URL   string "yaml:\"url\""
		}{
			Title: "Example Site",
			URL:   "https://example.com",
		},
		Author: []struct {
			Name  string "yaml:\"name\""
			Email string "yaml:\"email\""
			URL   string "yaml:\"url\""
			Bio   string "yaml:\"bio\""
		}{
			{
				Name:  "John Doe",
				Email: "john@example.com",
				URL:   "https://example.com/john",
				Bio:   "Web developer and tech enthusiast",
			},
			{
				Name:  "Jane Smith",
				Email: "jane@example.com",
				URL:   "https://example.com/jane",
				Bio:   "Designer and photographer",
			},
		},
	}

	config, err := GetConfig(cfgData)
	if err != nil {
		t.Errorf("Error getting config: %s", err)
	}

	if config.Site.Title != expectedConfig.Site.Title {
		t.Errorf("Expected title %q, got %q", expectedConfig.Site.Title, config.Site.Title)
	}

	if config.Site.URL != expectedConfig.Site.URL {
		t.Errorf("Expected URL %q, got %q", expectedConfig.Site.URL, config.Site.URL)
	}

	if len(config.Author) != len(expectedConfig.Author) {
		t.Errorf("Expected %d authors, got %d", len(expectedConfig.Author), len(config.Author))
	}

	for i, author := range config.Author {
		if author.Name != expectedConfig.Author[i].Name {
			t.Errorf("Expected name %q for author %d, got %q", expectedConfig.Author[i].Name, i, author.Name)
		}

		if author.Email != expectedConfig.Author[i].Email {
			t.Errorf("Expected email %q for author %d, got %q", expectedConfig.Author[i].Email, i, author.Email)
		}

		if author.URL != expectedConfig.Author[i].URL {
			t.Errorf("Expected URL %q for author %d, got %q", expectedConfig.Author[i].URL, i, author.URL)
		}

		if author.Bio != expectedConfig.Author[i].Bio {
			t.Errorf("Expected bio %q for author %d, got %q", expectedConfig.Author[i].Bio, i, author.Bio)
		}
	}
}

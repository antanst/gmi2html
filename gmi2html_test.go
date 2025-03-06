package gmi2html

import (
	"strings"
	"testing"
)

func TestConvertGeminiContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple text line",
			input:    "This is a simple text line",
			expected: `<p class="gemini-textline">This is a simple text line</p>`,
		},
		{
			name:     "Heading line level 1",
			input:    "# Main heading",
			expected: `<h1 class="gemini-heading-1">Main heading</h1>`,
		},
		{
			name:     "Link line with description",
			input:    "=> https://example.com Example site",
			expected: `<div class="gemini-link-container"><a href="https://example.com">Example site</a></div>`,
		},
		{
			name:     "List item",
			input:    "* List item 1",
			expected: `<p class="gemini-list-item">• List item 1</p>`,
		},
		{
			name:     "Quote line",
			input:    "> This is a quote",
			expected: `<blockquote class="gemini-blockquote">This is a quote</blockquote>`,
		},
		{
			name:     "Preformatted text",
			input:    "```\ncode line 1\ncode line 2\n```",
			expected: `<pre class="gemini-preformatted">code line 1</pre><pre class="gemini-preformatted">code line 2</pre>`,
		},
		{
			name:     "Mixed content",
			input:    "# Title\n\nNormal paragraph\n\n=> https://example.com Link to example",
			expected: `<h1 class="gemini-heading-1">Title</h1><p class="gemini-textline"></p><p class="gemini-textline">Normal paragraph</p><p class="gemini-textline"></p><div class="gemini-link-container"><a href="https://example.com">Link to example</a></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertGeminiContent(tt.input)
			if result != tt.expected {
				t.Errorf("convertGeminiContent(%q):\ngot:  %s\nwant: %s",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestGmi2html(t *testing.T) {
	sample := "# Hello Gemini\n\nThis is a test document.\n\n=> https://gemini.circumlunar.space/ Project Gemini"
	result, _ := Gmi2html(sample, "Gemini Test", false)

	// Check that it contains the expected elements
	if !strings.Contains(result, "<title>Gemini Test</title>") {
		t.Error("Output HTML missing title")
	}

	if !strings.Contains(result, "<h1 class=\"gemini-heading-1\">Hello Gemini</h1>") {
		t.Error("Output HTML missing properly formatted heading")
	}

	if !strings.Contains(result, "<a href=\"https://gemini.circumlunar.space/\">Project Gemini</a>") {
		t.Error("Output HTML missing properly formatted link")
	}

	// Check that CSS is included
	if !strings.Contains(result, "<style>") {
		t.Error("Output HTML missing style section")
	}

	// Check that basic HTML structure is there
	if !strings.Contains(result, "<!DOCTYPE html>") {
		t.Error("Output HTML missing doctype declaration")
	}

	if !strings.Contains(result, "<div class=\"gemini-container\">") {
		t.Error("Output HTML missing container div")
	}
}

func TestParseGeminiLink(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedURL  string
		expectedDesc string
		expectError  bool
	}{
		{
			name:         "Valid link with description",
			input:        "=> https://example.com Example site",
			expectedURL:  "https://example.com",
			expectedDesc: "Example site",
			expectError:  false,
		},
		{
			name:         "Valid link without description",
			input:        "=> https://example.com",
			expectedURL:  "https://example.com",
			expectedDesc: "https://example.com",
			expectError:  false,
		},
		{
			name:         "Link with special characters",
			input:        "=> https://example.com/search?q=test Test search",
			expectedURL:  "https://example.com/search?q=test",
			expectedDesc: "Test search",
			expectError:  false,
		},
		{
			name:         "Malformed link line",
			input:        "=>",
			expectedURL:  "",
			expectedDesc: "",
			expectError:  true,
		},
		{
			name:         "Link with multiple spaces in description",
			input:        "=> https://example.com   Multiple    spaces",
			expectedURL:  "https://example.com",
			expectedDesc: "Multiple    spaces",
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, desc, err := parseGeminiLink(tt.input)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got error: %v", tt.expectError, err != nil)
			}

			// If we expect success, check the result
			if !tt.expectError {
				if url != tt.expectedURL {
					t.Errorf("Expected URL: %s, got: %s", tt.expectedURL, url)
				}
				if desc != tt.expectedDesc {
					t.Errorf("Expected description: %s, got: %s", tt.expectedDesc, desc)
				}
			}
		})
	}
}

package gmi2html

import (
	"testing"
)

func TestGmi2html(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		title           string
		contentOnly     bool
		replaceGmiExt   bool
		want            string
		wantErr         bool
	}{
		{
			name:            "Basic text",
			input:           "Hello world",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<p class=\"gemini-textline\">Hello world</p>",
			wantErr:         false,
		},
		{
			name:            "Headers",
			input:           "# Header 1\n## Header 2\n### Header 3",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<h1 class=\"gemini-heading-1\">Header 1</h1><h2 class=\"gemini-heading-2\">Header 2</h2><h3 class=\"gemini-heading-3\">Header 3</h3>",
			wantErr:         false,
		},
		{
			name:            "List items",
			input:           "* Item 1\n* Item 2",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<p class=\"gemini-list-item\">• Item 1</p><p class=\"gemini-list-item\">• Item 2</p>",
			wantErr:         false,
		},
		{
			name:            "Blockquote",
			input:           "> This is a quote",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<blockquote class=\"gemini-blockquote\">This is a quote</blockquote>",
			wantErr:         false,
		},
		{
			name:            "Link",
			input:           "=> https://example.com Example Link",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<div class=\"gemini-link-container\"><a href=\"https://example.com\">Example Link</a></div>",
			wantErr:         false,
		},
		{
			name:            "Link with gmi extension replacement",
			input:           "=> /path/file.gmi Example Link",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   true,
			want:            "<div class=\"gemini-link-container\"><a href=\"/path/file.html\">Example Link</a></div>",
			wantErr:         false,
		},
		{
			name:            "Preformatted text",
			input:           "```\nThis is preformatted\n```",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<pre class=\"gemini-preformatted\">\nThis is preformatted\n</pre>",
			wantErr:         false,
		},
		{
			name:            "Mixed content",
			input:           "# Title\nNormal text\n=> https://example.com Link\n```\nCode\n```\n* List item",
			title:           "Test",
			contentOnly:     true,
			replaceGmiExt:   false,
			want:            "<h1 class=\"gemini-heading-1\">Title</h1><p class=\"gemini-textline\">Normal text</p><div class=\"gemini-link-container\"><a href=\"https://example.com\">Link</a></div><pre class=\"gemini-preformatted\">\nCode\n</pre><p class=\"gemini-list-item\">• List item</p>",
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Gmi2html(tt.input, tt.title, tt.contentOnly, tt.replaceGmiExt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gmi2html() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gmi2html() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGeminiLink(t *testing.T) {
	tests := []struct {
		name           string
		linkLine       string
		replaceGmiExt  bool
		wantURL        string
		wantDesc       string
		wantErr        bool
	}{
		{
			name:           "Basic link",
			linkLine:       "=> https://example.com Example Link",
			replaceGmiExt:  false,
			wantURL:        "https://example.com",
			wantDesc:       "Example Link",
			wantErr:        false,
		},
		{
			name:           "Link without description",
			linkLine:       "=> https://example.com",
			replaceGmiExt:  false,
			wantURL:        "https://example.com",
			wantDesc:       "https://example.com",
			wantErr:        false,
		},
		{
			name:           "Invalid link format",
			linkLine:       "Invalid line",
			replaceGmiExt:  false,
			wantURL:        "",
			wantDesc:       "",
			wantErr:        true,
		},
		{
			name:           "Link with .gmi extension, no replacement",
			linkLine:       "=> /path/file.gmi Link to Gemini file",
			replaceGmiExt:  false,
			wantURL:        "/path/file.gmi",
			wantDesc:       "Link to Gemini file",
			wantErr:        false,
		},
		{
			name:           "Link with .gmi extension, with replacement",
			linkLine:       "=> /path/file.gmi Link to Gemini file",
			replaceGmiExt:  true,
			wantURL:        "/path/file.html",
			wantDesc:       "Link to Gemini file",
			wantErr:        false,
		},
		{
			name:           "Link without .gmi extension, with replacement",
			linkLine:       "=> /path/file.txt Link to text file",
			replaceGmiExt:  true,
			wantURL:        "/path/file.txt",
			wantDesc:       "Link to text file",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, gotDesc, err := parseGeminiLink(tt.linkLine, tt.replaceGmiExt)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGeminiLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("parseGeminiLink() gotURL = %v, want %v", gotURL, tt.wantURL)
			}
			if gotDesc != tt.wantDesc {
				t.Errorf("parseGeminiLink() gotDesc = %v, want %v", gotDesc, tt.wantDesc)
			}
		})
	}
}

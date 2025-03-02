package gmi2html

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"regexp"
	"strings"
)

// Based on https://geminiprotocol.net/docs/gemtext-specification.gmi

// Gmi2html converts Gemini text to HTML with proper escaping and wraps it in a container with typography-focused CSS
func Gmi2html(text string, title string) (string, error) {
	content := convertGeminiContent(text)

	// Handle any template errors with container
	var buffer bytes.Buffer
	err := containerTmpl.Execute(&buffer, struct {
		Title   string
		Content template.HTML
	}{
		Title:   title,
		Content: template.HTML(content), // Content already properly escaped in convertGeminiContent
	})
	if err != nil {
		fmt.Printf("Error executing container template: %s\n", err)
		return "", err
	}

	return buffer.String(), nil
}

// convertGeminiContent converts Gemini text to HTML with proper escaping
func convertGeminiContent(text string) string {
	lines := strings.Split(text, "\n")
	var buffer bytes.Buffer
	normalMode := true

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "=>"):
			handleLinkLine(&buffer, line)
		case strings.HasPrefix(line, "```"):
			normalMode = !normalMode
			// Don't output the ``` line
		case strings.HasPrefix(line, "###"):
			content := strings.TrimSpace(strings.TrimPrefix(line, "###"))
			err := h3Tmpl.Execute(&buffer, content)
			if err != nil {
				return ""
			}
		case strings.HasPrefix(line, "##"):
			content := strings.TrimSpace(strings.TrimPrefix(line, "##"))
			err := h2Tmpl.Execute(&buffer, content)
			if err != nil {
				return ""
			}
		case strings.HasPrefix(line, "#"):
			content := strings.TrimSpace(strings.TrimPrefix(line, "#"))
			err := h1Tmpl.Execute(&buffer, content)
			if err != nil {
				return ""
			}
		case strings.HasPrefix(line, "*"):
			content := strings.TrimSpace(strings.TrimPrefix(line, "*"))
			err := listItemTmpl.Execute(&buffer, content)
			if err != nil {
				return ""
			}
		case strings.HasPrefix(line, ">"):
			content := strings.TrimSpace(strings.TrimPrefix(line, ">"))
			err := blockquoteTmpl.Execute(&buffer, content)
			if err != nil {
				return ""
			}
		default:
			if normalMode {
				err := textLineTmpl.Execute(&buffer, line)
				if err != nil {
					return ""
				}
			} else {
				err := preformattedTmpl.Execute(&buffer, line)
				if err != nil {
					return ""
				}
			}
		}
	}

	return buffer.String()
}

// handleLinkLine parses and renders a link line
func handleLinkLine(buffer *bytes.Buffer, linkLine string) {
	url, description, err := parseGeminiLink(linkLine)
	if err != nil {
		fmt.Printf("Error parsing gemini link line: %s\n", err)
		return
	}

	err = linkTmpl.Execute(buffer, struct {
		URL, Description string
	}{url, description})
	if err != nil {
		return
	}
}

// parseGeminiLink extracts URL and description from a link line
func parseGeminiLink(linkLine string) (string, string, error) {
	re := regexp.MustCompile(`^=>[ \t]+(\S+)([ \t]+.*)?`)
	matches := re.FindStringSubmatch(linkLine)
	if len(matches) == 0 {
		return "", "", fmt.Errorf("error parsing link line: no regexp match for line %s", linkLine)
	}

	urlStr := matches[1]

	// Check: Unescape the URL if escaped
	_, err := url.QueryUnescape(urlStr)
	if err != nil {
		return "", "", fmt.Errorf("error parsing link line: %w input '%s'", err, linkLine)
	}

	// Set description to URL if not provided
	description := urlStr
	if len(matches) > 2 && strings.TrimSpace(matches[2]) != "" {
		description = strings.TrimSpace(matches[2])
	}

	return urlStr, description, nil
}

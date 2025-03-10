package gmi2html

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/url"
	"regexp"
	"strings"
)

// Based on https://geminiprotocol.net/docs/gemtext-specification.gmi

//go:embed assets/main.html
var rawTemplate string

// Gmi2html converts Gemini text to HTML with proper escaping and wraps it in a container with typography-focused CSS
func Gmi2html(text string, title string, contentOnly bool, replaceGmiExt bool) (string, error) {
	content := convertGeminiContent(text, replaceGmiExt)

	if contentOnly {
		return content, nil
	}

	tmpl := template.Must(template.New("gemini").Parse(rawTemplate))

	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, struct {
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
func convertGeminiContent(text string, replaceGmiExt bool) string {
	lines := strings.Split(text, "\n")
	var buffer bytes.Buffer
	normalMode := true

	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "=>"):
			handleLinkLine(&buffer, line, replaceGmiExt)
		case strings.HasPrefix(line, "```"):
			if normalMode {
				err := preformattedTmplStart.Execute(&buffer, line)
				if err != nil {
					return ""
				}
				buffer.WriteString("\n")
			} else {
				err := preformattedTmplEnd.Execute(&buffer, line)
				if err != nil {
					return ""
				}
			}
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
				buffer.WriteString(line)
				buffer.WriteString("\n")
			}
		}
	}

	return buffer.String()
}

// handleLinkLine parses and renders a link line
func handleLinkLine(buffer *bytes.Buffer, linkLine string, replaceGmiExt bool) {
	url, description, err := parseGeminiLink(linkLine, replaceGmiExt)
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
func parseGeminiLink(linkLine string, replaceGmiExt bool) (string, string, error) {
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

	// Replace .gmi extension with .html if requested
	if replaceGmiExt && strings.HasSuffix(urlStr, ".gmi") {
		urlStr = strings.TrimSuffix(urlStr, ".gmi") + ".html"
	}

	// Set description to URL if not provided
	description := urlStr
	if len(matches) > 2 && strings.TrimSpace(matches[2]) != "" {
		description = strings.TrimSpace(matches[2])
	}

	return urlStr, description, nil
}

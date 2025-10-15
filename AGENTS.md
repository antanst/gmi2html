# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Essential Commands

Build, test, and develop:
```shell
make              # Full workflow: format, lint, test, clean, build
make test         # Run tests only
make build        # Build binary to ./dist/gmi2html
make fmt          # Format code with gofumpt and gci
make lint         # Run linter after formatting
make lintfix      # Run linter with auto-fix
```

Running the tool:
```shell
./dist/gmi2html <input.gmi >output.html
./dist/gmi2html --no-container <input.gmi >content.html
./dist/gmi2html --replace-gmi-ext <input.gmi >output.html
```

## Architecture

This is a Go library and CLI tool that converts Gemini text format to HTML.

**Core Components:**
- `gmi2html.go`: Main conversion logic with `Gmi2html()` function and `convertGeminiContent()` for parsing
- `templates.go`: HTML templates for each Gemini element type (headings, links, lists, etc.)
- `cmd/gmi2html/main.go`: CLI entry point that reads from stdin and writes to stdout
- `assets/main.html`: Embedded HTML template with CSS for the full page container

**Key Architecture Patterns:**
- Uses Go's `html/template` package for safe HTML generation with automatic escaping
- Embeds the main HTML template using `//go:embed` directive
- Line-by-line parser that switches between normal and preformatted modes
- Two output modes: full HTML document or content-only for embedding
- Optional `.gmi` to `.html` link conversion for static site generation

**Gemini Format Support:**
- Headings (#, ##, ###), links (=>), lists (*), quotes (>), preformatted blocks (```)
- Proper handling of preformatted content with mode switching
- URL parsing and validation for links
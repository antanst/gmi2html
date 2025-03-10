# gmi2html

A small library (and CLI tool) that converts Gemini text to HTML.

To run tests and build:

```shell
make
```

Running:

```shell
./dist/gmi2html <gemtext.gmi >gemtext.html
```

Options:

- `--no-container`: Don't output container HTML
- `--replace-gmi-ext`: Replace .gmi extension with .html in links

Example:

```shell
# Convert Gemini text and replace all .gmi links with .html
./dist/gmi2html --replace-gmi-ext <input.gmi >output.html

# Convert only the content without wrapping it in the HTML container
./dist/gmi2html --no-container <input.gmi >output-content.html
```

Help:
```shell
./dist/gmi2html --help
```

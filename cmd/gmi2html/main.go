package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/antanst/gmi2html/pkg/gmi2html"
)

func main() {
	noContainer := flag.Bool("no-container", false, "Don't output container HTML")
	replaceGmiExt := flag.Bool("replace-gmi-ext", false, "Replace .gmi extension with .html in links")
	flag.Parse()
	err := runApp(*noContainer, *replaceGmiExt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runApp(noContainer bool, replaceGmiExt bool) error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	html, err := gmi2html.Gmi2html(string(data), "", noContainer, replaceGmiExt)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(os.Stdout, "%s", html)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/antanst/gmi2html"
)

func main() {
	err := runApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runApp() error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	html, err := gmi2html.Gmi2html(string(data), "", false)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(os.Stdout, "%s", html)
	if err != nil {
		return err
	}
	return nil
}

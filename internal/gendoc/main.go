package main

import (
	"log"

	"github.com/johejo/gf-cli/internal"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := internal.RootCmd()
	if err := doc.GenMarkdownTree(cmd, "./docs"); err != nil {
		log.Fatal(err)
	}
}

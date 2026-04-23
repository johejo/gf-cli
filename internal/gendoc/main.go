package main

import (
	"log"

	"github.com/johejo/gf-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := internal.RootCmd()
	// The body schema is carried on Annotations so --help can render it after
	// the Flags block. cobra/doc only renders Long, so splice it back in here.
	inlineBodySchema(cmd)
	if err := doc.GenMarkdownTree(cmd, "./docs"); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenManTree(cmd, nil, "./man"); err != nil {
		log.Fatal(err)
	}
}

func inlineBodySchema(cmd *cobra.Command) {
	if schema := cmd.Annotations["bodySchema"]; schema != "" {
		// cobra/doc renders Synopsis from Long. Mirror the pre-change shape
		// where Short was the lead paragraph so body-only commands keep their
		// one-line summary above the schema.
		switch {
		case cmd.Long != "":
			cmd.Long = cmd.Long + "\n\n" + schema
		case cmd.Short != "":
			cmd.Long = cmd.Short + "\n\n" + schema
		default:
			cmd.Long = schema
		}
	}
	for _, c := range cmd.Commands() {
		inlineBodySchema(c)
	}
}

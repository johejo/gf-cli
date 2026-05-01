package main

import (
	"log"
	"strings"

	"github.com/johejo/gf-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := internal.RootCmd()
	// Body and Response schemas are carried on Annotations so --help can
	// render them after the Flags block. cobra/doc only renders Long, so
	// splice them back in here.
	inlineSchemas(cmd)
	if err := doc.GenMarkdownTree(cmd, "./docs"); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenManTree(cmd, nil, "./man"); err != nil {
		log.Fatal(err)
	}
}

func inlineSchemas(cmd *cobra.Command) {
	body := cmd.Annotations["bodySchema"]
	response := cmd.Annotations["responseSchema"]
	if body != "" || response != "" {
		// cobra/doc renders Synopsis from Long. Mirror the pre-change shape
		// where Short was the lead paragraph so body-only commands keep their
		// one-line summary above the schema.
		var lead string
		switch {
		case cmd.Long != "":
			lead = cmd.Long
		case cmd.Short != "":
			lead = cmd.Short
		}
		parts := make([]string, 0, 3)
		if lead != "" {
			parts = append(parts, lead)
		}
		if body != "" {
			parts = append(parts, body)
		}
		if response != "" {
			parts = append(parts, response)
		}
		cmd.Long = strings.Join(parts, "\n\n")
	}
	for _, c := range cmd.Commands() {
		inlineSchemas(c)
	}
}

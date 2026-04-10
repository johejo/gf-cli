package main

import (
	"fmt"
	"go/parser"
)

type ResponseField struct {
	HasPayload bool
}

// ParseResponse finds the response struct by type name within the package directory
// and checks if it has a Payload field.
func ParseResponse(baseDir string, pkgName string, responseTypeName string) (*ResponseField, error) {
	pkgs, err := getParsedPkg(baseDir, pkgName, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			st := findStructType(f, responseTypeName)
			if st == nil {
				continue
			}
			for _, field := range st.Fields.List {
				if len(field.Names) > 0 && field.Names[0].Name == "Payload" {
					return &ResponseField{HasPayload: true}, nil
				}
			}
			return &ResponseField{HasPayload: false}, nil
		}
	}
	return nil, fmt.Errorf("response type %s not found in package %s", responseTypeName, pkgName)
}

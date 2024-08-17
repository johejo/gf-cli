package main

import (
	"log"
	"runtime/debug"

	"github.com/johejo/gf-cli/internal"
)

var (
	version = ""
)

func getVersion() string {
	if version != "" {
		return version
	}
	info, ok := debug.ReadBuildInfo()
	if ok {
		return info.Main.Version
	}
	return "(devel)"
}

func main() {
	cmd := internal.RootCmd()
	cmd.Version = getVersion()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

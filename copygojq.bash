#!/usr/bin/env bash

set -e -o pipefail

version="$(grep gojq go.mod | grep -o -E 'v.*$')"

function download() {
	rm -f "$2"

	cat <<EOF >"$2"
// Code copied from itchyny/gojq by copygojq.bash.
// See LICENSE.gojq for original license.
EOF

	curl -fsSL "https://raw.githubusercontent.com/itchyny/gojq/${version}/$1" >>"$2"
}

download cli/encoder.go ./internal/cli/encoder.go
download cli/color.go ./internal/cli/color.go
download LICENSE ./LICENSE.gojq

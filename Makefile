all:

OPENAPI_CLIENT_DIR := $(shell go list -m -json github.com/grafana/grafana-openapi-client-go | jq -r .Dir)

gen:
	go run ./internal/gen2 -base-dir=$(OPENAPI_CLIENT_DIR) -out-go=./internal/gen.go -out-help-json=./internal/help.json

gendoc:
	rm -rf docs man
	mkdir -p docs man
	go run ./internal/gendoc

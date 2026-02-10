all:

clean:
	rm -rf ./dist

gen:
	go run ./internal/gen > ./internal/gen.go

gendoc:
	rm -rf docs
	mkdir -p docs
	go run ./internal/gendoc

install: gen
	go install ./cmd/gf

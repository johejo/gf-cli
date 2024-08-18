all:

clean:
	rm -rf ./dist

gen:
	go run ./internal/gen > ./internal/gen.go

gendoc:
	go run ./internal/gendoc

install: gen
	go install ./cmd/gf

copygojq:
	./copygojq.bash

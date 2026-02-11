all:

gen:
	go run ./internal/gen > ./internal/gen.go

gendoc:
	rm -rf docs man
	mkdir -p docs man
	go run ./internal/gendoc

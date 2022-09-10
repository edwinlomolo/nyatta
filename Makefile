init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen

run:
	go run server.go

test:
#	go clean -testcache
	go test -v ./...

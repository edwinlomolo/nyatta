init:
	go get -v ./...
	go run github.com/99designs/gqlgen init --verbose

generate:
	go run github.com/99designs/gqlgen --verbose

migrate-db:
	sqlc generate

run:
	go run server.go

test:
#	go clean -testcache
	go test -v ./...

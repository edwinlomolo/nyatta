# init:
# 	go run github.com/99designs/gqlgen init --verbose
# 
# generate:
# 	go run github.com/99designs/gqlgen --verbose
# 
# migrate-db:
# 	sqlc generate
# 
test:
#	go clean -testcache
	go test -v ./...

run:
	go run server.go



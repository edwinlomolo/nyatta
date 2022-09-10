package nyatta_context

import "errors"

var (
	AlreadyExists = errors.New("AlreadExists")
	NotFound      = errors.New("NotFound")
	DatabaseError = errors.New("DatabaseError")
	ResolverError = errors.New("ResolverError")
)

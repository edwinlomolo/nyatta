package config

import "errors"

var (
	AlreadyExists        = errors.New("AlreadExists")
	NotFound             = errors.New("NotFound")
	DatabaseError        = errors.New("DatabaseError")
	ResolverError        = errors.New("ResolverError")
	CredentialsError     = errors.New("credentials error")
	MigrationErr         = errors.New("MigrationErr")
	MigrationDownErr     = errors.New("MigrationDownErr")
	MigrationUpErr       = errors.New("MigrationUpErr")
	MigrationInstanceErr = errors.New("MigrationInstanceErr")
	MigrationDriverErr   = errors.New("MigrationDriverErr")
)

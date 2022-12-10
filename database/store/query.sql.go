// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
	"time"
)

const createAmenity = `-- name: CreateAmenity :one
INSERT INTO amenities (
  name, provider, property_id
) VALUES (
  $1, $2, $3
)
RETURNING id, name, provider, created_at, updated_at, property_id
`

type CreateAmenityParams struct {
	Name       string `json:"name"`
	Provider   string `json:"provider"`
	PropertyID int64  `json:"property_id"`
}

func (q *Queries) CreateAmenity(ctx context.Context, arg CreateAmenityParams) (Amenity, error) {
	row := q.db.QueryRowContext(ctx, createAmenity, arg.Name, arg.Provider, arg.PropertyID)
	var i Amenity
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyID,
	)
	return i, err
}

const createProperty = `-- name: CreateProperty :one
INSERT INTO properties (
  name, town, postal_code, created_by
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, town, postal_code, created_at, updated_at, created_by
`

type CreatePropertyParams struct {
	Name       string `json:"name"`
	Town       string `json:"town"`
	PostalCode string `json:"postal_code"`
	CreatedBy  int64  `json:"created_by"`
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error) {
	row := q.db.QueryRowContext(ctx, createProperty,
		arg.Name,
		arg.Town,
		arg.PostalCode,
		arg.CreatedBy,
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Town,
		&i.PostalCode,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
	)
	return i, err
}

const createPropertyUnit = `-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms
) VALUES (
  $1, $2
)
RETURNING id, bathrooms, created_at, updated_at, property_id
`

type CreatePropertyUnitParams struct {
	PropertyID int64 `json:"property_id"`
	Bathrooms  int32 `json:"bathrooms"`
}

func (q *Queries) CreatePropertyUnit(ctx context.Context, arg CreatePropertyUnitParams) (PropertyUnit, error) {
	row := q.db.QueryRowContext(ctx, createPropertyUnit, arg.PropertyID, arg.Bathrooms)
	var i PropertyUnit
	err := row.Scan(
		&i.ID,
		&i.Bathrooms,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyID,
	)
	return i, err
}

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, end_date, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING id, start_date, end_date, created_at, updated_at, property_unit_id
`

type CreateTenantParams struct {
	StartDate      time.Time    `json:"start_date"`
	EndDate        sql.NullTime `json:"end_date"`
	PropertyUnitID int64        `json:"property_unit_id"`
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, createTenant, arg.StartDate, arg.EndDate, arg.PropertyUnitID)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
	)
	return i, err
}

const createUnitBedroom = `-- name: CreateUnitBedroom :one
INSERT INTO bedrooms (
  property_unit_id, bedroom_number, en_suite, master
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, bedroom_number, en_suite, master, created_at, updated_at, property_unit_id
`

type CreateUnitBedroomParams struct {
	PropertyUnitID int64 `json:"property_unit_id"`
	BedroomNumber  int32 `json:"bedroom_number"`
	EnSuite        bool  `json:"en_suite"`
	Master         bool  `json:"master"`
}

func (q *Queries) CreateUnitBedroom(ctx context.Context, arg CreateUnitBedroomParams) (Bedroom, error) {
	row := q.db.QueryRowContext(ctx, createUnitBedroom,
		arg.PropertyUnitID,
		arg.BedroomNumber,
		arg.EnSuite,
		arg.Master,
	)
	var i Bedroom
	err := row.Scan(
		&i.ID,
		&i.BedroomNumber,
		&i.EnSuite,
		&i.Master,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name
) VALUES (
  $1, $2, $3
)
RETURNING id, email, first_name, last_name, created_at, updated_at
`

type CreateUserParams struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.FirstName, arg.LastName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByEmail = `-- name: FindByEmail :one
SELECT id, email, first_name, last_name, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) FindByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProperty = `-- name: GetProperty :one
SELECT id, name, town, postal_code, created_at, updated_at, created_by FROM properties
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProperty(ctx context.Context, id int64) (Property, error) {
	row := q.db.QueryRowContext(ctx, getProperty, id)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Town,
		&i.PostalCode,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
	)
	return i, err
}

const getPropertyUnits = `-- name: GetPropertyUnits :many
SELECT id, bathrooms, created_at, updated_at, property_id FROM property_units
WHERE property_id = $1
`

func (q *Queries) GetPropertyUnits(ctx context.Context, propertyID int64) ([]PropertyUnit, error) {
	rows, err := q.db.QueryContext(ctx, getPropertyUnits, propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PropertyUnit
	for rows.Next() {
		var i PropertyUnit
		if err := rows.Scan(
			&i.ID,
			&i.Bathrooms,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitBedrooms = `-- name: GetUnitBedrooms :many
SELECT id, bedroom_number, en_suite, master, created_at, updated_at, property_unit_id FROM bedrooms
WHERE property_unit_id = $1
`

func (q *Queries) GetUnitBedrooms(ctx context.Context, propertyUnitID int64) ([]Bedroom, error) {
	rows, err := q.db.QueryContext(ctx, getUnitBedrooms, propertyUnitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bedroom
	for rows.Next() {
		var i Bedroom
		if err := rows.Scan(
			&i.ID,
			&i.BedroomNumber,
			&i.EnSuite,
			&i.Master,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyUnitID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitTenancy = `-- name: GetUnitTenancy :many
SELECT id, start_date, end_date, created_at, updated_at, property_unit_id FROM tenants
WHERE property_unit_id = $1
`

func (q *Queries) GetUnitTenancy(ctx context.Context, propertyUnitID int64) ([]Tenant, error) {
	rows, err := q.db.QueryContext(ctx, getUnitTenancy, propertyUnitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tenant
	for rows.Next() {
		var i Tenant
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyUnitID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, email, first_name, last_name, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const propertiesCreatedBy = `-- name: PropertiesCreatedBy :many
SELECT id, name, town, postal_code, created_at, updated_at, created_by FROM properties
WHERE created_by = $1
`

func (q *Queries) PropertiesCreatedBy(ctx context.Context, createdBy int64) ([]Property, error) {
	rows, err := q.db.QueryContext(ctx, propertiesCreatedBy, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Property
	for rows.Next() {
		var i Property
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Town,
			&i.PostalCode,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const propertyAmenities = `-- name: PropertyAmenities :many
SELECT id, name, provider, created_at, updated_at, property_id FROM amenities
WHERE property_id = $1
`

func (q *Queries) PropertyAmenities(ctx context.Context, propertyID int64) ([]Amenity, error) {
	rows, err := q.db.QueryContext(ctx, propertyAmenities, propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Amenity
	for rows.Next() {
		var i Amenity
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Provider,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

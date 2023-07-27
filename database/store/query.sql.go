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
  name, category, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING id, name, provider, created_at, category, updated_at, property_unit_id
`

type CreateAmenityParams struct {
	Name           string `json:"name"`
	Category       string `json:"category"`
	PropertyUnitID int64  `json:"property_unit_id"`
}

func (q *Queries) CreateAmenity(ctx context.Context, arg CreateAmenityParams) (Amenity, error) {
	row := q.db.QueryRowContext(ctx, createAmenity, arg.Name, arg.Category, arg.PropertyUnitID)
	var i Amenity
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.CreatedAt,
		&i.Category,
		&i.UpdatedAt,
		&i.PropertyUnitID,
	)
	return i, err
}

const createCaretaker = `-- name: CreateCaretaker :one
INSERT INTO caretakers (
  first_name, last_name, idVerification, country_code, phone, image
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, first_name, last_name, idverification, country_code, created_at, updated_at, phone, image, verified
`

type CreateCaretakerParams struct {
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Idverification string         `json:"idverification"`
	CountryCode    string         `json:"country_code"`
	Phone          sql.NullString `json:"phone"`
	Image          string         `json:"image"`
}

func (q *Queries) CreateCaretaker(ctx context.Context, arg CreateCaretakerParams) (Caretaker, error) {
	row := q.db.QueryRowContext(ctx, createCaretaker,
		arg.FirstName,
		arg.LastName,
		arg.Idverification,
		arg.CountryCode,
		arg.Phone,
		arg.Image,
	)
	var i Caretaker
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Idverification,
		&i.CountryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Phone,
		&i.Image,
		&i.Verified,
	)
	return i, err
}

const createProperty = `-- name: CreateProperty :one
INSERT INTO properties (
  name, town, postal_code, type, created_by, caretaker
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, name, town, postal_code, type, status, created_at, updated_at, created_by, caretaker
`

type CreatePropertyParams struct {
	Name       string        `json:"name"`
	Town       string        `json:"town"`
	PostalCode string        `json:"postal_code"`
	Type       string        `json:"type"`
	CreatedBy  int64         `json:"created_by"`
	Caretaker  sql.NullInt64 `json:"caretaker"`
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error) {
	row := q.db.QueryRowContext(ctx, createProperty,
		arg.Name,
		arg.Town,
		arg.PostalCode,
		arg.Type,
		arg.CreatedBy,
		arg.Caretaker,
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Town,
		&i.PostalCode,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.Caretaker,
	)
	return i, err
}

const createPropertyUnit = `-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms, name, type, price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, type, state, price, bathrooms, created_at, updated_at, property_id
`

type CreatePropertyUnitParams struct {
	PropertyID int64  `json:"property_id"`
	Bathrooms  int32  `json:"bathrooms"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Price      int32  `json:"price"`
}

func (q *Queries) CreatePropertyUnit(ctx context.Context, arg CreatePropertyUnitParams) (PropertyUnit, error) {
	row := q.db.QueryRowContext(ctx, createPropertyUnit,
		arg.PropertyID,
		arg.Bathrooms,
		arg.Name,
		arg.Type,
		arg.Price,
	)
	var i PropertyUnit
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.State,
		&i.Price,
		&i.Bathrooms,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyID,
	)
	return i, err
}

const createShootSchedule = `-- name: CreateShootSchedule :one
INSERT INTO shoots (
  shoot_date, property_id, property_unit_id, caretaker_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, shoot_date, property_id, property_unit_id, status, caretaker_id
`

type CreateShootScheduleParams struct {
	ShootDate      time.Time `json:"shoot_date"`
	PropertyID     int64     `json:"property_id"`
	PropertyUnitID int64     `json:"property_unit_id"`
	CaretakerID    int64     `json:"caretaker_id"`
}

func (q *Queries) CreateShootSchedule(ctx context.Context, arg CreateShootScheduleParams) (Shoot, error) {
	row := q.db.QueryRowContext(ctx, createShootSchedule,
		arg.ShootDate,
		arg.PropertyID,
		arg.PropertyUnitID,
		arg.CaretakerID,
	)
	var i Shoot
	err := row.Scan(
		&i.ID,
		&i.ShootDate,
		&i.PropertyID,
		&i.PropertyUnitID,
		&i.Status,
		&i.CaretakerID,
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
  email, first_name, last_name, avatar, phone
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at
`

type CreateUserParams struct {
	Email     sql.NullString `json:"email"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Avatar    sql.NullString `json:"avatar"`
	Phone     sql.NullString `json:"phone"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.Avatar,
		arg.Phone,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByEmail = `-- name: FindByEmail :one
SELECT id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) FindByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByPhone = `-- name: FindUserByPhone :one
SELECT id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at FROM users
WHERE phone = $1
`

func (q *Queries) FindUserByPhone(ctx context.Context, phone sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProperty = `-- name: GetProperty :one
SELECT id, name, town, postal_code, type, status, created_at, updated_at, created_by, caretaker FROM properties
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
		&i.Type,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.Caretaker,
	)
	return i, err
}

const getPropertyUnits = `-- name: GetPropertyUnits :many
SELECT id, name, type, state, price, bathrooms, created_at, updated_at, property_id FROM property_units
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
			&i.Name,
			&i.Type,
			&i.State,
			&i.Price,
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
SELECT id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at FROM users
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
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const mailingExists = `-- name: MailingExists :one
SELECT EXISTS(
  SELECT id, email FROM mailings
  WHERE email = $1
)
`

func (q *Queries) MailingExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, mailingExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const occupiedUnitsCount = `-- name: OccupiedUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'occupied'
`

func (q *Queries) OccupiedUnitsCount(ctx context.Context, propertyID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, occupiedUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const onboardUser = `-- name: OnboardUser :one
UPDATE users
SET onboarding = $1
WHERE email = $2
RETURNING id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at
`

type OnboardUserParams struct {
	Onboarding sql.NullBool   `json:"onboarding"`
	Email      sql.NullString `json:"email"`
}

func (q *Queries) OnboardUser(ctx context.Context, arg OnboardUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, onboardUser, arg.Onboarding, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const propertiesCreatedBy = `-- name: PropertiesCreatedBy :many
SELECT id, name, town, postal_code, type, status, created_at, updated_at, created_by, caretaker FROM properties
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
			&i.Type,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.Caretaker,
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

const propertyUnitsCount = `-- name: PropertyUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1
`

func (q *Queries) PropertyUnitsCount(ctx context.Context, propertyID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, propertyUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const saveMail = `-- name: SaveMail :one
INSERT INTO mailings (
  email
) VALUES (
  $1
)
RETURNING id, email
`

func (q *Queries) SaveMail(ctx context.Context, email string) (Mailing, error) {
	row := q.db.QueryRowContext(ctx, saveMail, email)
	var i Mailing
	err := row.Scan(&i.ID, &i.Email)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET avatar = $1, first_name = $2, last_name = $3, onboarding = $4
WHERE email = $5
RETURNING id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at
`

type UpdateUserParams struct {
	Avatar     sql.NullString `json:"avatar"`
	FirstName  sql.NullString `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Onboarding sql.NullBool   `json:"onboarding"`
	Email      sql.NullString `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Avatar,
		arg.FirstName,
		arg.LastName,
		arg.Onboarding,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPhone = `-- name: UpdateUserPhone :one
UPDATE users
SET phone = $1
WHERE email = $2
RETURNING id, email, first_name, last_name, phone, onboarding, avatar, created_at, updated_at
`

type UpdateUserPhoneParams struct {
	Phone sql.NullString `json:"phone"`
	Email sql.NullString `json:"email"`
}

func (q *Queries) UpdateUserPhone(ctx context.Context, arg UpdateUserPhoneParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPhone, arg.Phone, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const vacantUnitsCount = `-- name: VacantUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'vacant'
`

func (q *Queries) VacantUnitsCount(ctx context.Context, propertyID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, vacantUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name, avatar, phone
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  name, town, postal_code, type, created_by
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetProperty :one
SELECT * FROM properties
WHERE id = $1 LIMIT 1;

-- name: FindByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: PropertiesCreatedBy :many
SELECT * FROM properties
WHERE created_by = $1;

-- name: CreateAmenity :one
INSERT INTO amenities (
  name, category, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, end_date, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms, name, type, price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateUnitBedroom :one
INSERT INTO bedrooms (
  property_unit_id, bedroom_number, en_suite, master
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUnitBedrooms :many
SELECT * FROM bedrooms
WHERE property_unit_id = $1;

-- name: GetUnitTenancy :many
SELECT * FROM tenants
WHERE property_unit_id = $1;

-- name: GetPropertyUnits :many
SELECT * FROM property_units
WHERE property_id = $1;

-- name: UpdateUser :one
UPDATE users
SET avatar = $1, first_name = $2, last_name = $3, onboarding = $4
WHERE email = $5
RETURNING *;

-- name: FindUserByPhone :one
SELECT * FROM users
WHERE phone = $1;

-- name: CreateCaretaker :one
INSERT INTO caretakers (
  first_name, last_name, idVerification, country_code, phone
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateShootSchedule :one
INSERT INTO shoots (
  shoot_date, property_id, property_unit_id, caretaker_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

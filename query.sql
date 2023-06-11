-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name, avatar
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  name, town, postal_code, type, min_price, max_price, created_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
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
  name, provider, category, property_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: PropertyAmenities :many
SELECT * FROM amenities
WHERE property_id = $1;

-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, end_date, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms
) VALUES (
  $1, $2
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

-- name: GetListings :many
SELECT * FROM properties
WHERE town ILIKE $1 AND min_price >= $2 AND max_price <= $3;

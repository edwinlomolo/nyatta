-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  phone, next_renewal
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  name, type, created_by, caretaker_id, location
) VALUES (
  $1, $2, $3, $4, sqlc.arg(location)
)
RETURNING *;

-- name: GetProperty :one
SELECT id, name, type, ST_AsGeoJSON(location)::jsonb AS location, created_by, created_at, updated_at FROM properties
WHERE id = $1 LIMIT 1;

-- name: PropertiesCreatedBy :many
SELECT p.id, p.name, p.type, ST_AsGeoJSON(p.location) AS location, p.created_by, p.created_at, p.updated_at FROM properties p WHERE p.created_by = $1
UNION
SELECT u.id, u.name, u.type, ST_AsGeoJSON(u.location) AS location, u.created_by, u.created_at, u.updated_at FROM units u WHERE u.created_by = $1
ORDER BY updated_at;

-- name: CreateAmenity :one
INSERT INTO amenities (
  name, category, unit_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUnitAmenities :many
SELECT * FROM amenities
WHERE unit_id = $1;

-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, unit_id, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCurrentTenant :one
SELECT * FROM tenants
WHERE unit_id = $1 AND end_date IS NULL;

-- name: CreateUnit :one
INSERT INTO units (
  property_id, bathrooms, name, type, price, state
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: CreateOtherUnit :one
INSERT INTO units (
  property_id, bathrooms, name, type, price, state, caretaker_id, created_by, location
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, sqlc.arg(location)
)
RETURNING *;

-- name: GetUnit :one
SELECT id, name, type, state, price, bathrooms, ST_AsGeoJSON(location)::jsonb AS location, created_by, caretaker_id, property_id, created_at, updated_at FROM units
WHERE id = $1;

-- name: CreateUnitBedroom :one
INSERT INTO bedrooms (
  unit_id, bedroom_number, en_suite, master
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUnitBedrooms :many
SELECT * FROM bedrooms
WHERE unit_id = $1;

-- name: GetUnitTenancy :many
SELECT * FROM tenants
WHERE unit_id = $1;

-- name: GetUserTenancy :many
SELECT * FROM tenants
WHERE user_id = $1;

-- name: GetUnits :many
SELECT * FROM units
WHERE property_id = $1;

-- name: FindUserByPhone :one
SELECT * FROM users
WHERE phone = $1;

-- name: CreateCaretaker :one
INSERT INTO caretakers (
  first_name, last_name, phone, created_by
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetCaretakerByPhone :one
SELECT * FROM caretakers
WHERE phone = $1;

-- name: GetCaretakerById :one
SELECT * FROM caretakers
WHERE id = $1;

-- name: CreateShootSchedule :one
INSERT INTO shoots (
  shoot_date, property_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: SaveMail :one
INSERT INTO mailings (
  email
) VALUES (
  $1
)
RETURNING *;

-- name: MailingExists :one
SELECT EXISTS(
  SELECT * FROM mailings
  WHERE email = $1
);

-- name: UnitsCount :one
SELECT COUNT(*) FROM units
WHERE property_id = $1;

-- name: OccupiedUnitsCount :one
SELECT COUNT(*) FROM units
WHERE property_id = $1 AND state = 'occupied';

-- name: VacantUnitsCount :one
SELECT COUNT(*) FROM units
WHERE property_id = $1 AND state = 'vacant';

-- name: CreateInvoice :one
INSERT INTO invoices (
  reference, phone, msid
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateInvoiceForMpesa :one
UPDATE invoices
SET channel = $1, status = $2, amount = $3, currency = $4, bank = $5, auth_code = $6, country_code = $7, fees = $8, created_at = $9, updated_at = $10
WHERE reference = $11
RETURNING *;

-- name: UpdateLandlord :one
UPDATE users
SET next_renewal = $1
WHERE phone = $2
RETURNING *;

-- name: UpdateUserInfo :one
UPDATE users
SET first_name = $1, last_name = $2
WHERE phone = $3
RETURNING *;

-- name: CreateUserAvatar :one
INSERT INTO uploads (
  upload, category, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateCaretakerAvatar :one
INSERT INTO uploads (
  upload, category, caretaker_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserAvatar :one
SELECT id, upload, category FROM uploads
WHERE user_id = $1 AND category = $2 LIMIT 1;

-- name: GetCaretakerAvatar :one
SELECT id, upload, category FROM uploads
WHERE caretaker_id = $1 AND category = $2;

-- name: UpdateUpload :one
UPDATE uploads
SET upload = $1
WHERE id = $2
RETURNING *;

-- name: CreatePropertyThumbnail :one
INSERT INTO uploads (
  upload, category, property_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPropertyThumbnail :one
SELECT id, upload FROM uploads
WHERE property_id = $1 AND category = $2 LIMIT 1;

-- name: CreateUnitImage :one
INSERT INTO uploads (
  upload, category, label, unit_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUnitImages :many
SELECT id, upload, label FROM uploads
WHERE unit_id = $1 AND category = $2 LIMIT 1;

-- name: TrackSubscribeRetries :one
UPDATE users SET subscribe_retries = $1
WHERE phone = $2
RETURNING *;

-- name: GetNearByUnits :many
SELECT u.id, u.property_id, u.name, u.type, u.price, u.updated_at, ST_Distance(p.location, sqlc.arg(point)::geography) AS distance FROM properties p
JOIN units u
ON ST_DWithin(p.location, sqlc.arg(point)::geography, 10000) WHERE u.property_id = p.id AND u.state = 'VACANT'
UNION
SELECT u.id, u.property_id, u.name, u.type, u.price, u.updated_at, ST_Distance(u.location, sqlc.arg(point)::geography) AS distance FROM units u
WHERE ST_DWithin(u.location, sqlc.arg(point)::geography, 10000) AND u.state = 'VACANT'
ORDER BY updated_at;

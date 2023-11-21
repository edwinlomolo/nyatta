-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  phone
) VALUES (
  $1
)
RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  name, type, created_by, caretaker_id, location
) VALUES (
  $1, $2, $3, $4, ST_GeomFromText(sqlc.arg(location)::text)
)
RETURNING *;

-- name: GetProperty :one
SELECT * FROM properties
WHERE id = $1 LIMIT 1;

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
  start_date, property_unit_id, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms, name, type, price, state
) VALUES (
  $1, $2, $3, $4, $5, $6
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

-- name: PropertyUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1;

-- name: OccupiedUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'occupied';

-- name: VacantUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'vacant';

-- name: UnitAmenityCount :one
SELECT COUNT(*) from amenities
WHERE property_unit_id = $1;

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
  upload, category, label, property_unit_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUnitImages :many
SELECT id, upload, label FROM uploads
WHERE property_unit_id = $1 AND category = $2 LIMIT 1;

-- name: TrackSubscribeRetries :one
UPDATE users SET subscribe_retries = $1
WHERE phone = $2
RETURNING *;

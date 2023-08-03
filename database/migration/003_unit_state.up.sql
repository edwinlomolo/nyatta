CREATE TYPE IF NOT EXISTS unit_state AS ENUM ('vacant', 'unavailable', 'occupied');

ALTER TABLE property_units
ADD COLUMN IF NOT EXISTS state unit_state NOT NULL DEFAULT 'vacant';

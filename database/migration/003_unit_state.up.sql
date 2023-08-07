DO $$ BEGIN
  CREATE TYPE unit_state AS ENUM ('vacant', 'unavailable', 'occupied');
EXCEPTION
  WHEN duplicate_object THEN null;
END $$;

ALTER TABLE property_units
ADD COLUMN IF NOT EXISTS state unit_state NOT NULL DEFAULT 'vacant';

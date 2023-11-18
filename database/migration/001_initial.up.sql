CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  next_renewal TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  phone VARCHAR(15) UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS caretakers (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  phone VARCHAR(15) UNIQUE,
  verified BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS properties (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  location GEOMETRY(POINT, 4326) NOT NULL,
  type VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by BIGINT REFERENCES users ON DELETE SET NULL ON UPDATE CASCADE,
  caretaker_id BIGINT REFERENCES caretakers ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS property_units (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  type VARCHAR(100) NOT NULL,
  state VARCHAR(50) NOT NULL DEFAULT 'VACANT',
  price INTEGER NOT NULL,
  bathrooms INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_id BIGINT REFERENCES properties ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS amenities (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  provider VARCHAR(100) NOT NULL DEFAULT '',
  category VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT REFERENCES property_units ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS tenants (
  id BIGSERIAL PRIMARY KEY,
  start_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  end_date TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT REFERENCES property_units ON DELETE SET NULL,
  user_id BIGINT REFERENCES users ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS bedrooms (
  id BIGSERIAL PRIMARY KEY,
  bedroom_number INTEGER NOT NULL,
  en_suite BOOLEAN NOT NULL DEFAULT false,
  master BOOLEAN NOT NULL DEFAULT false,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE SET NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS uploads (
  id BIGSERIAL PRIMARY KEY,
  upload TEXT NOT NULL,
  category VARCHAR(100) NOT NULL,
  label VARCHAR(50),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT REFERENCES property_units ON DELETE CASCADE,
  property_id BIGINT REFERENCES properties ON DELETE CASCADE,
  user_id BIGINT REFERENCES users ON DELETE CASCADE,
  caretaker_id BIGINT REFERENCES caretakers ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS shoots (
  id BIGSERIAL PRIMARY KEY,
  shoot_date TIMESTAMP NOT NULL,
  property_id BIGINT NOT NULL REFERENCES properties ON DELETE CASCADE,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE CASCADE,
  status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
  caretaker_id BIGINT NOT NULL REFERENCES caretakers ON DELETE SET NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS mailings (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS invoices (
  id BIGSERIAL PRIMARY KEY,
  msid VARCHAR(15),
  channel VARCHAR(100),
  currency TEXT,
  bank TEXT,
  auth_code TEXT,
  country_code VARCHAR(5),
  fees money,
  amount money,
  phone VARCHAR(15) REFERENCES users(phone) ON DELETE CASCADE,
  status VARCHAR(50) NOT NULL DEFAULT 'PROCESSING',
  reference VARCHAR(20),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

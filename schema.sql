CREATE DATABASE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE,
  first_name TEXT,
  last_name TEXT,
  avatar TEXT,
  phone VARCHAR(15) NOT NULL,
  onboarding BOOLEAN DEFAULT true,
  is_landlord BOOLEAN DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS caretakers (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  image TEXT NOT NULL,
  phone VARCHAR(15),
  verified BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS properties (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  thumbnail TEXT NOT NULL,
  location GEOMETRY,
  type VARCHAR(20) NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by BIGINT REFERENCES users ON DELETE SET NULL ON UPDATE CASCADE,
  caretaker_id BIGINT REFERENCES caretakers ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TYPE unit_state AS ENUM ('vacant', 'unavailable', 'occupied');
CREATE TABLE IF NOT EXISTS property_units (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  type VARCHAR(100) NOT NULL,
  state unit_state NOT NULL DEFAULT 'vacant',
  location GEOMETRY,
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
  user_id BIGINT REFERENCES users on DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS bedrooms (
  id BIGSERIAL PRIMARY KEY,
  bedroom_number INTEGER NOT NULL,
  en_suite BOOLEAN NOT NULL DEFAULT false,
  master BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS uploads (
  id BIGSERIAL PRIMARY KEY,
  image TEXT NOT NULL,
  category VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS shoots (
  id BIGSERIAL PRIMARY KEY,
  shoot_date TIMESTAMP NOT NULL,
  property_id BIGINT NOT NULL REFERENCES properties ON DELETE CASCADE,
  property_unit_id BIGINT NOT NULL REFERENCES property_units ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  caretaker_id BIGINT NOT NULL REFERENCES caretakers ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS mailings (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL UNIQUE
);

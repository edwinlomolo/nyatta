CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email text NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS properties (
  id BIGSERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  town varchar(50) NOT NULL,
  postal_code varchar(20) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  created_by bigint NOT NULL REFERENCES users ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS amenities (
  id BIGSERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  provider varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  property_id bigint NOT NULL REFERENCES properties ON DELETE CASCADE
);

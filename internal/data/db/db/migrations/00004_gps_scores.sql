-- +goose Up
-- Add GPS latitude to the locations table
ALTER TABLE locations
ADD COLUMN gps_latitude REAL;

-- Add GPS longitude to the locations table
ALTER TABLE locations
ADD COLUMN gps_longitude REAL;

-- Add GPS altitude to the locations table
ALTER TABLE locations
ADD COLUMN gps_altitude REAL;

-- Add GPS URL to the locations table
ALTER TABLE locations
ADD COLUMN gps_url TEXT;

-- +goose Down
-- Revert GPS latitude addition
ALTER TABLE locations
DROP COLUMN gps_latitude;

-- Revert GPS longitude addition
ALTER TABLE locations
DROP COLUMN gps_longitude;

-- Revert GPS altitude addition
ALTER TABLE locations
DROP COLUMN gps_altitude;

-- Revert GPS URL addition
ALTER TABLE locations
DROP COLUMN gps_url;

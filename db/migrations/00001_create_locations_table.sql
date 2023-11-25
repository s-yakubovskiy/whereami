-- +goose Up
CREATE TABLE IF NOT EXISTS locations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ip TEXT,
    country TEXT,
    countryCode TEXT,
    region TEXT,
    regionCode TEXT,
    city TEXT,
    timezone TEXT,
    zip TEXT,
    flag TEXT,
    emojiFlag TEXT,
    isp TEXT,
    org TEXT,
    asn TEXT,
    latitude REAL,
    longitude REAL,
    date TEXT,
    vpn INTEGER
);

-- +goose Down
DROP TABLE IF EXISTS locations;


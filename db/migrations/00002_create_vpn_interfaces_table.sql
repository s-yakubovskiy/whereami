-- +goose Up
CREATE TABLE IF NOT EXISTS vpn_interfaces (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    interface_name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS vpn_interfaces;

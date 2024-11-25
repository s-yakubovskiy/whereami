-- +goose Up
CREATE TABLE IF NOT EXISTS location_scores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    location_id INTEGER,
    fraud_score INTEGER,
    host TEXT,
    proxy INTEGER,
    vpn INTEGER,
    tor INTEGER,
    is_crawler INTEGER,
    recent_abuse INTEGER,
    bot_status INTEGER,
    FOREIGN KEY (location_id) REFERENCES locations(id)
);

-- +goose Down
DROP TABLE IF EXISTS location_scores;

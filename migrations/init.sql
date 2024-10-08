CREATE DATABASE IF NOT EXISTS url_shortener;

USE url_shortener;

CREATE TABLE IF NOT EXISTS url_shortener (
    url_id VARCHAR(20) PRIMARY KEY,
    url_original TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_url_id ON url_shortener(url_id);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    author TEXT NOT NULL
);
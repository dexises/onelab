CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at date,
    deleted_at date,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE  if not exists books_readers (
    `id`          int PRIMARY KEY AUTO_INCREMENT,
    `book_id`     int,
    `reader_id`   int,
    `taken_at`    date NOT NULL,
    `returned_at` date ,
    `created_at`  timestamp NOT NULL DEFAULT (now()),
    `updated_at`  timestamp NOT NULL DEFAULT (now())
);

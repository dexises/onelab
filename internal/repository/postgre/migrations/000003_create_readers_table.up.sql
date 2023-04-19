CREATE TABLE IF NOT EXISTS books_readers (
    id SERIAL PRIMARY KEY,
    book_id INTEGER,
    reader_id INTEGER,
    taken_at DATE NOT NULL,
    returned_at DATE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

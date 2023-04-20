CREATE TABLE IF NOT EXISTS books_readers (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL,
    reader_id INTEGER NOT NULL,
    transaction_id INTEGER NOT NULL,
    title TEXT NOT NULL ,
    taken_at DATE NOT NULL,
    returned_at DATE,
    foreign key (reader_id) references users(id),
    foreign key (book_id) references books(id),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

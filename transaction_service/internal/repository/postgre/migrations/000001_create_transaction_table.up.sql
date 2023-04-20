create table if not exists transactions (
    id serial primary key,
    book_id INTEGER not null,
    user_id INTEGER not null,
    book_price INTEGER NOT NULL ,
    user_balance INTEGER NOT NULL ,
    borrow_date timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
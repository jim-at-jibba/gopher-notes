CREATE TABLE IF NOT EXISTS notes (
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    title text NOT NULL,
    text text NOT NULL,
    user_id  integer NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

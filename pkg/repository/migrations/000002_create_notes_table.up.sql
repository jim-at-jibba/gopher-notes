CREATE TABLE IF NOT EXISTS notes (
    id UUID PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    title text NOT NULL,
    text text NOT NULL,
    user_id  UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

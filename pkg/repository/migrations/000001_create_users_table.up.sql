CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    username text not null,
    password text not null
);

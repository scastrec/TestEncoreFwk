CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    pwd TEXT NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL
);

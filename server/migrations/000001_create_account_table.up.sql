CREATE TABLE IF NOT EXISTS account (
    id uuid PRIMARY KEY,
    email varchar NOT NULL UNIQUE,
    email_verified boolean NOT NULL DEFAULT false,
    username varchar NOT NULL UNIQUE,
    is_active boolean NOT NULL DEFAULT true,
    password bytea NOT NULL
);
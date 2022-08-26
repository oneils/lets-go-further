CREATE TABLE IF NOT EXISTS users
(
    id            bigserial PRIMARY KEY       NOT NULL,
    created_at    timestamp(0) with time zone NOT NULL DEFAULT now(),
    name          text                        NOT NULL,
    email         text UNIQUE               NOT NULL,
    password_hash bytea                       NOT NULL,
    activated     boolean                     NOT NULL DEFAULT false,
    version       integer                     NOT NULL DEFAULT 1
)
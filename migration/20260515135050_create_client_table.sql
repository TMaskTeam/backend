-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.client (
    client_id       SERIAL      PRIMARY KEY,
    first_name      TEXT                        NOT NULL,
    middle_name     TEXT                        NULL
    last_name       TEXT                        NOT NULL,
    phone_number    TEXT        UNIQUE          NOT NULL,
    email           TEXT        UNIQUE          NOT NULL,
    birthday        DATE                        NOT NULL,   
    password_hash   TEXT                        NOT NULL
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP 
);  

-- +goose Down
DROP TABLE IF EXISTS    mask.client;
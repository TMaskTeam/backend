-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.business_owner (
    owner_id        SERIAL      PRIMARY KEY,
    password_hash   TEXT                        NOT NULL,
    first_name      TEXT                        NOT NULL,
    middle_name     TEXT                        NULL,
    last_name       TEXT                        NOT NULL,
    inn             TEXT        UNIQUE          NOT NULL,
    phone_number    TEXT        UNIQUE          NOT NULL,
    email           TEXT        UNIQUE          NOT NULL,
    birthday        DATE                        NOT NULL,   
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP   
);   

-- +goose Down
DROP TABLE IF EXISTS    mask.business_owner;

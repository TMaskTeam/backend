-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.client_bonus_program (
    client_bonus_program  SERIAL      PRIMARY KEY,
    program_id            SERIAL,
    client_id             SERIAL,
    tokens                TEXT,
    created_at            TIMESTAMP,
    updated_at            TIMESTAMP   
);  

-- +goose Down
DROP TABLE IF EXISTS    mask.client_bonus_program;
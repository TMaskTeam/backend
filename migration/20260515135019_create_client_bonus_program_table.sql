-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.client_bonus_program (
    client_bonus_program  SERIAL      PRIMARY KEY,
    program_id            INT,
    client_id             INT,
    tokens_count          TEXT,
    created_at            TIMESTAMP,
    updated_at            TIMESTAMP   
);  

ALTER TABLE mask.client_bonus_program
ADD CONSTRAINT fk_program_id
FOREIGN KEY (program_id)
REFERENCES mask.bonus_program(program_id)
ON DELETE CASCADE;

ALTER TABLE mask.client
ADD CONSTRAINT fk_client_id
FOREIGN KEY (client_id)
REFERENCES mask.client(client_id)
ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS    mask.client_bonus_program;
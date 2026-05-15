-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.bonus_program (
    program_id   SERIAL      NOT NULL,
    business_id  INTEGER     NOT NULL,
    program_name TEXT        NOT NULL,
    token_name   TEXT        NOT NULL,
    created_at   TIMESTAMP,
    updated_at   TIMESTAMP,
    PRIMARY KEY (program_id)
);


ALTER TABLE mask.bonus_program
ADD CONSTRAINT fk_bonus_program
FOREIGN KEY (business_id)
REFERENCES mask.business(business_id)
ON DELETE CASCADE;


-- +goose Down
ALTER TABLE mask.bonus_program DROP CONSTRAINT IF EXISTS fk_bonus_program;
DROP TABLE IF EXISTS    mask.bonus_program;

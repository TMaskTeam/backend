-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.bonus_program (
    program_id      SERIAL          PRIMARY KEY,
    business_id     INTEGER         NOT NULL,
    program_name    TEXT            NOT NULL,
    token_name      TEXT            NOT NULL,
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP
);


ALTER TABLE mask.bonus_program
ADD CONSTRAINT fk_business
FOREIGN KEY (business_id)
REFERENCES mask.business(business_id)
ON DELETE CASCADE;


-- +goose Down
ALTER TABLE mask.bonus_program DROP CONSTRAINT IF EXISTS fk_business;
DROP TABLE IF EXISTS    mask.bonus_program;





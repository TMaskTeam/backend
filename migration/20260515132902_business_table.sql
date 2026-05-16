-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.business (
    business_id     SERIAL          PRIMARY KEY,
    owner_id        INTEGER         NOT NULL,
    name            TEXT            NOT NULL,
    address         TEXT            NOT NULL,
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP
);


ALTER TABLE mask.business
ADD CONSTRAINT fk_business_owner
FOREIGN KEY (owner_id)
REFERENCES mask.business_owner(owner_id)
ON DELETE CASCADE;


-- +goose Down
ALTER TABLE mask.business DROP CONSTRAINT IF EXISTS fk_business_owner;
DROP TABLE IF EXISTS    mask.business;




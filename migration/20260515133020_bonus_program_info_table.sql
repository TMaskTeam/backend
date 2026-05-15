-- +goose Up
CREATE SCHEMA IF NOT EXISTS mask;

CREATE TABLE IF NOT EXISTS mask.bonus_program_info (
    program_info_id             SERIAL          PRIMARY KEY,
    program_id                  INTEGER         NOT NULL,
    visit_tokens                INTEGER         NOT NULL,
    percentage_purchase_tokens  INTEGER     NOT NULL,
    register_tokens             INTEGER     NOT NULL,
    birthday_tokens             INTEGER     NOT NULL,
    friend_invite_tokens        INTEGER     NOT NULL,
    minimum_receipt_limit       INTEGER     NOT NULL,
    created_at                  TIMESTAMP,
    updated_at                  TIMESTAMP,
);


ALTER TABLE mask.bonus_program_info
ADD CONSTRAINT fk_bonus_program
FOREIGN KEY (program_id)
REFERENCES mask.bonus_program(program_id)
ON DELETE CASCADE;


-- +goose Down
ALTER TABLE mask.bonus_program_info DROP CONSTRAINT IF EXISTS fk_bonus_program;
DROP TABLE IF EXISTS mask.bonus_program_info;





-- +goose Up
CREATE TABLE IF NOT EXISTS mask.employee (
    employee_id         SERIAL          PRIMARY KEY,
    client_id           INTEGER         NOT NULL REFERENCES mask.client(client_id),
    business_id         INTEGER         NOT NULL REFERENCES mask.business(business_id),
    assigned_by_owner_id INTEGER        NOT NULL REFERENCES mask.business_owner(owner_id),
    created_at          TIMESTAMP       NOT NULL DEFAULT NOW(),
    UNIQUE (client_id, business_id)
);

-- +goose Down
DROP TABLE IF EXISTS mask.employee;

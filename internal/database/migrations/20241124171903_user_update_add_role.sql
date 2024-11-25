-- +goose Up
-- +goose StatementBegin
ALTER TYPE USER_ROLE
ADD
    VALUE 'NURSE';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS USER_ROLE;

CREATE TYPE USER_ROLE AS ENUM ('DOCTOR', 'RECEPTIONIST');

-- +goose StatementEnd
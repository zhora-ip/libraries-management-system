-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id BIGSERIAL primary key,
    login TEXT not null unique,
    encrypted_password TEXT not null,
    full_name TEXT not null,
    phone_number TEXT unique,
    email TEXT unique,
    role INT not null,
    created_at TIMESTAMP not null default now(),
    updated_at TIMESTAMP not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users cascade;
-- +goose StatementEnd

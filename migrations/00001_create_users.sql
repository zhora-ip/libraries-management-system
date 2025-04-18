-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id BIGSERIAL primary key,
    login TEXT not null unique,
    password_hash TEXT not null,
    full_name TEXT not null,
    phone_number TEXT,
    email TEXT unique,
    created_at TIMESTAMP not null default now(),
    updated_at TIMESTAMP not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users cascade;
-- +goose StatementEnd

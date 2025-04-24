-- +goose Up
-- +goose StatementBegin
create table if not exists tasks(
    id SERIAL primary key,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now(),
    finished_at TIMESTAMPTZ,
    status INT not null,
    attempt_count INT not null default 0,
    type INT not null,
    payload BYTEA not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd

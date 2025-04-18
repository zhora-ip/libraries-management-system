-- +goose Up
-- +goose StatementBegin
create table if not exists books(
    id BIGSERIAL primary key,
    title TEXT not null,
    author TEXT not null,
    genre TEXT,
    description TEXT,
    age_limit INTEGER,
    created_at TIMESTAMP not null default now(),
    updated_at TIMESTAMP not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists books;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
create table if not exists lib_cards(
    id BIGSERIAL primary key,
    code TEXT not null unique,
    user_id BIGINT REFERENCES users(id) unique,
    created_at TIMESTAMP not null default now(),
    expires_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists lib_cards;
-- +goose StatementEnd

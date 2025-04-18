-- +goose Up
-- +goose StatementBegin
create table if not exists lib_cards(
    id BIGSERIAL primary key,
    code TEXT not null unique,
    user_id BIGINT REFERENCES users(id),
    created_at TIMESTAMP not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists lib_cards;
-- +goose StatementEnd

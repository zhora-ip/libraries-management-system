-- +goose Up
-- +goose StatementBegin
create table if not exists libraries(
    id BIGSERIAL primary key,
    name TEXT not null,
    address TEXT not null,
    phone_number TEXT,
    latitude DOUBLE PRECISION not null,
    longitude DOUBLE PRECISION not null,
    created_at TIMESTAMP not null default now(),
    updated_at TIMESTAMP not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists libraries;
-- +goose StatementEnd

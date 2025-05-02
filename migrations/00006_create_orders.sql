-- +goose Up
-- +goose StatementBegin
create table if not exists orders (
    id BIGSERIAL primary key,
    library_id BIGINT not null REFERENCES libraries(id),
    physical_book_id BIGINT not null REFERENCES physical_books(id),
    user_id BIGINT not null REFERENCES users(id) ON DELETE CASCADE,
    status INT not null default 0,
    created_at TIMESTAMP not null default now(),
    updated_at TIMESTAMP not null default now(),
    expires_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists orders;
-- +goose StatementEnd

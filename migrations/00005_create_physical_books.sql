-- +goose Up
-- +goose StatementBegin
create table if not exists physical_books(
    id BIGSERIAL primary key,
    library_id BIGINT not null REFERENCES library(id),
    book_id BIGINT not null REFERENCES books(id),
    is_available BOOLEAN not null default true,
    created_at TIMESTAMP not null default now(),
    updated_at TIMESTAMP not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists physical_books;
-- +goose StatementEnd

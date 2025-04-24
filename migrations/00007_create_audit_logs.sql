-- +goose Up
-- +goose StatementBegin
create table audit_logs(
    id BIGSERIAL primary key,
    created_at TIMESTAMPTZ not null,
    log TEXT
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table audit_logs;
-- +goose StatementEnd

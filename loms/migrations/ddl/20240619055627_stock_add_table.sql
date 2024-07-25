-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stocks
(
    id bigserial primary key,
    sku bigint not null,
    total_count bigint not null,
    reserved int default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS stock;
-- +goose StatementEnd

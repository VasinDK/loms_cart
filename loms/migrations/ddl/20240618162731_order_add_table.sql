-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
    id bigserial primary key,
    user_id bigint not null,
    status varchar(50) default 'new'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS order;
-- +goose StatementEnd

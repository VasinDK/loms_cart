-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items_order
(
    id bigserial primary key,
    sku bigint not null,
    count int not null,
    order_id bigserial,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS items_order;
-- +goose StatementEnd

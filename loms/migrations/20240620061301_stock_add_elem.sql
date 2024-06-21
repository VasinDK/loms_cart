-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks (sku, total_count) VALUES (1148162, 260);
INSERT INTO stocks (sku, total_count) VALUES (1076963, 270);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DELETE 
-- FROM stocks
-- WHERE sku in (1148162, 1076963);
-- +goose StatementEnd

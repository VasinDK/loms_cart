-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks (sku, total_count) VALUES (1148162, 260);
INSERT INTO stocks (sku, total_count) VALUES (1076963, 270);
INSERT INTO stocks (sku, total_count) VALUES (773297411, 100);
INSERT INTO stocks (sku, total_count) VALUES (2618151, 100);
INSERT INTO stocks (sku, total_count) VALUES (2956315, 100);
INSERT INTO stocks (sku, total_count) VALUES (2958025, 100);
INSERT INTO stocks (sku, total_count) VALUES (3596599, 100);
INSERT INTO stocks (sku, total_count) VALUES (3618852, 100);
INSERT INTO stocks (sku, total_count) VALUES (4288068, 100);
INSERT INTO stocks (sku, total_count) VALUES (4465995, 100);
INSERT INTO stocks (sku, total_count) VALUES (4487693, 100);
INSERT INTO stocks (sku, total_count) VALUES (4669069, 100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DELETE 
-- FROM stocks
-- WHERE sku in (1148162, 1076963);
-- +goose StatementEnd

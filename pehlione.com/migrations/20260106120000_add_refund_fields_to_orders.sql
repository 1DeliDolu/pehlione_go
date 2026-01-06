-- +goose Up
use pehlione_go;
ALTER TABLE orders
  ADD COLUMN refunded_cents BIGINT NOT NULL DEFAULT 0,
  ADD COLUMN refunded_at DATETIME NULL;

CREATE INDEX idx_orders_refunded_at ON orders(refunded_at);

-- +goose Down
DROP INDEX idx_orders_refunded_at ON orders;

ALTER TABLE orders
  DROP COLUMN refunded_at,
  DROP COLUMN refunded_cents;

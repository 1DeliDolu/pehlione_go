-- +goose Up
CREATE TABLE IF NOT EXISTS fx_rates (
  currency CHAR(3) NOT NULL,
  rate DECIMAL(18,8) NOT NULL,
  source VARCHAR(32) NOT NULL,
  fetched_at DATETIME(3) NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (currency)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE orders
  ADD COLUMN base_currency CHAR(3) NOT NULL DEFAULT 'TRY' AFTER currency,
  ADD COLUMN base_subtotal_cents INT NOT NULL DEFAULT 0 AFTER base_currency,
  ADD COLUMN base_tax_cents INT NOT NULL DEFAULT 0 AFTER base_subtotal_cents,
  ADD COLUMN base_shipping_cents INT NOT NULL DEFAULT 0 AFTER base_tax_cents,
  ADD COLUMN base_discount_cents INT NOT NULL DEFAULT 0 AFTER base_shipping_cents,
  ADD COLUMN base_total_cents INT NOT NULL DEFAULT 0 AFTER base_discount_cents,
  ADD COLUMN display_currency CHAR(3) NOT NULL DEFAULT 'TRY' AFTER base_total_cents,
  ADD COLUMN fx_rate DECIMAL(18,8) NULL AFTER display_currency,
  ADD COLUMN fx_source VARCHAR(32) NULL AFTER fx_rate;

UPDATE orders
SET
  base_currency = currency,
  base_subtotal_cents = subtotal_cents,
  base_tax_cents = tax_cents,
  base_shipping_cents = shipping_cents,
  base_discount_cents = discount_cents,
  base_total_cents = total_cents,
  display_currency = currency,
  fx_rate = 1.0,
  fx_source = 'seed';

ALTER TABLE order_items
  ADD COLUMN base_currency CHAR(3) NOT NULL DEFAULT 'TRY' AFTER currency,
  ADD COLUMN base_unit_price_cents INT NOT NULL DEFAULT 0 AFTER base_currency,
  ADD COLUMN base_line_total_cents INT NOT NULL DEFAULT 0 AFTER base_unit_price_cents;

UPDATE order_items
SET
  base_currency = currency,
  base_unit_price_cents = unit_price_cents,
  base_line_total_cents = line_total_cents;

-- +goose Down
ALTER TABLE order_items
  DROP COLUMN base_line_total_cents,
  DROP COLUMN base_unit_price_cents,
  DROP COLUMN base_currency;

ALTER TABLE orders
  DROP COLUMN fx_source,
  DROP COLUMN fx_rate,
  DROP COLUMN display_currency,
  DROP COLUMN base_total_cents,
  DROP COLUMN base_discount_cents,
  DROP COLUMN base_shipping_cents,
  DROP COLUMN base_tax_cents,
  DROP COLUMN base_subtotal_cents,
  DROP COLUMN base_currency;

DROP TABLE IF EXISTS fx_rates;

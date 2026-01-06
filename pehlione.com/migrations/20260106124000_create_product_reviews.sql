-- +goose Up
CREATE TABLE IF NOT EXISTS product_reviews (
  id CHAR(36) NOT NULL,
  product_id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  order_id CHAR(36) NOT NULL,
  rating INT NOT NULL,
  body TEXT NULL,
  status VARCHAR(16) NOT NULL DEFAULT 'pending',
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  deleted_at DATETIME(3) NULL,
  PRIMARY KEY (id),
  UNIQUE KEY ux_reviews_product_user (product_id, user_id),
  KEY idx_reviews_product_status (product_id, status),
  KEY idx_reviews_user (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE products
  ADD COLUMN rating_avg DECIMAL(4,2) NOT NULL DEFAULT 0 AFTER description,
  ADD COLUMN rating_count INT NOT NULL DEFAULT 0 AFTER rating_avg;

-- +goose Down
ALTER TABLE products
  DROP COLUMN rating_count,
  DROP COLUMN rating_avg;

DROP TABLE IF EXISTS product_reviews;

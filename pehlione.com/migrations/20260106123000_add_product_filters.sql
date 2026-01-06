-- +goose Up
ALTER TABLE products
  ADD COLUMN category_name VARCHAR(255) NULL AFTER description,
  ADD COLUMN category_slug VARCHAR(255) NULL AFTER category_name;

CREATE INDEX idx_products_status ON products(status);
CREATE INDEX idx_products_name ON products(name);
CREATE INDEX idx_products_category_slug ON products(category_slug);
CREATE INDEX idx_product_variants_price_stock ON product_variants(price_cents, stock);
CREATE INDEX idx_product_variants_product_price ON product_variants(product_id, price_cents);
CREATE INDEX idx_product_variants_product_stock ON product_variants(product_id, stock);

-- Backfill category columns based on description pattern "Category: Name | ..."
UPDATE products
SET category_name = TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(description, '|', 1), ':', -1))
WHERE description LIKE 'Category:%'
  AND (category_name IS NULL OR category_name = '');

UPDATE products
SET category_slug = LOWER(
        REPLACE(
          REPLACE(
            REPLACE(TRIM(SUBSTRING_INDEX(SUBSTRING_INDEX(description, '|', 1), ':', -1)), ' ', '-'),
          '&', 'and'),
        '''', '')
     )
WHERE description LIKE 'Category:%'
  AND (category_slug IS NULL OR category_slug = '');

UPDATE products
SET category_name = COALESCE(NULLIF(category_name, ''), 'all'),
    category_slug = COALESCE(NULLIF(category_slug, ''), 'all');

-- +goose Down
ALTER TABLE products
  DROP COLUMN category_slug,
  DROP COLUMN category_name;

DROP INDEX idx_products_status ON products;
DROP INDEX idx_products_name ON products;
DROP INDEX idx_products_category_slug ON products;
DROP INDEX idx_product_variants_price_stock ON product_variants;
DROP INDEX idx_product_variants_product_price ON product_variants;
DROP INDEX idx_product_variants_product_stock ON product_variants;

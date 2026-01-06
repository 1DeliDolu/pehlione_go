-- +goose Up
ALTER TABLE users
  ADD COLUMN email_verified_at DATETIME(3) NULL AFTER email,
  ADD COLUMN status VARCHAR(16) NOT NULL DEFAULT 'pending' AFTER email_verified_at;

CREATE INDEX ix_users_status ON users(status);

-- +goose Down
ALTER TABLE users DROP INDEX ix_users_status;
ALTER TABLE users DROP COLUMN status;
ALTER TABLE users DROP COLUMN email_verified_at;

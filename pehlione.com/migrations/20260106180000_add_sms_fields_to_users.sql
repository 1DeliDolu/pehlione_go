-- +goose Up
ALTER TABLE users
  ADD COLUMN phone_e164 VARCHAR(32) NULL,
  ADD COLUMN phone_verified_at DATETIME NULL,
  ADD COLUMN sms_opt_in TINYINT(1) NOT NULL DEFAULT 0,
  ADD COLUMN sms_opt_out_at DATETIME NULL;

CREATE INDEX idx_users_phone_e164 ON users(phone_e164);

-- +goose Down
DROP INDEX idx_users_phone_e164 ON users;
ALTER TABLE users
  DROP COLUMN sms_opt_out_at,
  DROP COLUMN sms_opt_in,
  DROP COLUMN phone_verified_at,
  DROP COLUMN phone_e164;
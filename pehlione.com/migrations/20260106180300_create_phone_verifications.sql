-- +goose Up
CREATE TABLE IF NOT EXISTS phone_verifications (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  phone_e164 VARCHAR(32) NOT NULL,
  code_hash CHAR(64) NOT NULL,
  expires_at DATETIME NOT NULL,
  used_at DATETIME NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_phone_verif_user (user_id),
  INDEX idx_phone_verif_phone (phone_e164),
  INDEX idx_phone_verif_expires (expires_at)
);

-- +goose Down
DROP TABLE IF EXISTS phone_verifications;
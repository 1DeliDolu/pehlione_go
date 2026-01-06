-- +goose Up
CREATE TABLE IF NOT EXISTS email_verifications (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  token_hash CHAR(64) NOT NULL,
  expires_at DATETIME NOT NULL,
  used_at DATETIME NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uq_email_verifications_token_hash (token_hash),
  INDEX idx_email_verifications_user (user_id),
  INDEX idx_email_verifications_expires (expires_at)
);

-- +goose Down
DROP TABLE IF EXISTS email_verifications;

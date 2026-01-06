-- +goose Up
CREATE TABLE IF NOT EXISTS sms_consents (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  phone_e164 VARCHAR(32) NOT NULL,
  action VARCHAR(16) NOT NULL, -- opt_in | opt_out | system_opt_out
  source VARCHAR(32) NOT NULL, -- profile | checkout | inbound_stop | admin
  ip VARCHAR(64) NULL,
  user_agent VARCHAR(255) NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_sms_consents_user (user_id),
  INDEX idx_sms_consents_phone (phone_e164),
  INDEX idx_sms_consents_created (created_at)
);

-- +goose Down
DROP TABLE IF EXISTS sms_consents;
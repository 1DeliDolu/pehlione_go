-- +goose Up
CREATE TABLE IF NOT EXISTS sms_outbox (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  to_phone_e164 VARCHAR(32) NOT NULL,
  template VARCHAR(64) NOT NULL, -- shipped | delivered | otp
  payload JSON NOT NULL,
  status VARCHAR(16) NOT NULL DEFAULT 'pending', -- pending|processing|sent|failed
  attempt_count INT NOT NULL DEFAULT 0,
  last_error TEXT NULL,
  scheduled_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  locked_at DATETIME NULL,
  locked_by VARCHAR(128) NULL,
  sent_at DATETIME NULL,
  provider_message_id VARCHAR(128) NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_sms_outbox_pending (status, scheduled_at),
  INDEX idx_sms_outbox_phone (to_phone_e164)
);

-- +goose Down
DROP TABLE IF EXISTS sms_outbox;
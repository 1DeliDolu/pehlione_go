-- +goose Up
use pehlione_go;
CREATE TABLE IF NOT EXISTS email_outbox (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  to_email VARCHAR(320) NOT NULL,
  template VARCHAR(128) NOT NULL,
  payload JSON NOT NULL,
  status VARCHAR(16) NOT NULL DEFAULT 'pending',
  attempt_count INT NOT NULL DEFAULT 0,
  last_error TEXT NULL,
  scheduled_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  locked_at DATETIME(3) NULL,
  locked_by VARCHAR(128) NULL,
  sent_at DATETIME(3) NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  INDEX idx_email_outbox_pending (status, scheduled_at),
  INDEX idx_email_outbox_locked (locked_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS email_outbox;

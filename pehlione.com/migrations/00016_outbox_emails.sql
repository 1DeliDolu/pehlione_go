-- +goose Up
CREATE TABLE outbox_emails (
  id CHAR(36) NOT NULL,
  to_email VARCHAR(255) NOT NULL,
  subject VARCHAR(255) NOT NULL,

  body_text TEXT NULL,
  body_html MEDIUMTEXT NULL,

  status VARCHAR(16) NOT NULL DEFAULT 'pending',
  attempts INT NOT NULL DEFAULT 0,
  last_error VARCHAR(255) NULL,
  next_attempt_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  sent_at DATETIME(3) NULL,

  PRIMARY KEY (id),
  KEY ix_outbox_status_next (status, next_attempt_at),
  KEY ix_outbox_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS outbox_emails;

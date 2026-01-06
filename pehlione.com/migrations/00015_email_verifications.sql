-- +goose Up
CREATE TABLE email_verifications (
  id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,

  code_hash VARBINARY(32) NOT NULL,
  expires_at DATETIME(3) NOT NULL,
  used_at DATETIME(3) NULL,

  attempts INT NOT NULL DEFAULT 0,
  last_sent_at DATETIME(3) NULL,

  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

  PRIMARY KEY (id),
  UNIQUE KEY ux_email_verifications_user_active (user_id, used_at),
  KEY ix_email_verifications_expires (expires_at),
  CONSTRAINT fk_email_verifications_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS email_verifications;

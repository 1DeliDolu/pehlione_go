-- +goose Up
CREATE TABLE IF NOT EXISTS shipments (
  id CHAR(36) NOT NULL,
  order_id CHAR(36) NOT NULL,
  provider VARCHAR(32) NOT NULL,
  carrier VARCHAR(64) NOT NULL,
  service VARCHAR(64) NULL,
  note VARCHAR(255) NULL,
  tracking_no VARCHAR(128) NULL,
  tracking_url VARCHAR(512) NULL,
  label_url VARCHAR(512) NULL,
  status VARCHAR(32) NOT NULL,
  error_message TEXT NULL,
  shipped_at DATETIME(3) NULL,
  delivered_at DATETIME(3) NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_shipments_order_created (order_id, created_at),
  KEY idx_shipments_status (status),
  CONSTRAINT fk_shipments_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS shipment_jobs (
  id BIGINT NOT NULL AUTO_INCREMENT,
  shipment_id CHAR(36) NOT NULL,
  status VARCHAR(16) NOT NULL,
  payload JSON NOT NULL,
  attempt_count INT NOT NULL DEFAULT 0,
  last_error TEXT NULL,
  scheduled_at DATETIME(3) NOT NULL,
  locked_at DATETIME(3) NULL,
  locked_by VARCHAR(64) NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY ux_shipment_jobs_shipment (shipment_id),
  KEY idx_shipment_jobs_status_sched (status, scheduled_at),
  CONSTRAINT fk_shipment_jobs_shipment FOREIGN KEY (shipment_id) REFERENCES shipments(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS shipment_jobs;
DROP TABLE IF EXISTS shipments;

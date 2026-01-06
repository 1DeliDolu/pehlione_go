package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is required")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Create email_outbox table
	sql := `CREATE TABLE IF NOT EXISTS email_outbox (
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
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	if err := db.Exec(sql).Error; err != nil {
		log.Fatalf("failed to create email_outbox table: %v", err)
	}

	fmt.Println("✅ email_outbox table created successfully!")
}

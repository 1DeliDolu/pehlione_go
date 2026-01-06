package main

import (
	"log"
	"os"

	"log/slog"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pehlione.com/app/internal/config"
	apphttp "pehlione.com/app/internal/http"
)

func main() {
	// Load .env file (ignore error if not found - prod uses real env vars)
	_ = godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := apphttp.NewRouter(logger, db, cfg)
	_ = r.Run(":8080")
}

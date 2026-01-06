package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is required")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Get all SQL files from migrations directory
	migrationDir := "migrations"
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		log.Fatalf("failed to read migrations directory: %v", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}

	// Sort files by name (both old and new format)
	sort.Strings(sqlFiles)

	fmt.Printf("Found %d migration files\n\n", len(sqlFiles))

	// Execute each migration
	for _, filename := range sqlFiles {
		filePath := filepath.Join(migrationDir, filename)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("❌ Failed to read %s: %v\n", filename, err)
			continue
		}

		sqlContent := string(content)

		// Extract the UP section (between +goose Up and +goose Down)
		upPattern := regexp.MustCompile(`(?s)--\s*\+goose\s+Up(.*?)(?:--\s*\+goose\s+Down|$)`)
		matches := upPattern.FindStringSubmatch(sqlContent)

		if len(matches) < 2 {
			fmt.Printf("⚠️  Skipped %s (no +goose Up marker)\n", filename)
			continue
		}

		upSQL := strings.TrimSpace(matches[1])

		// Remove SQL comments
		upSQL = regexp.MustCompile(`--.*$`).ReplaceAllString(upSQL, "")
		upSQL = strings.TrimSpace(upSQL)

		if upSQL == "" {
			fmt.Printf("⚠️  Skipped %s (no SQL in UP section)\n", filename)
			continue
		}

		// Split by semicolon to handle multiple statements
		statements := strings.Split(upSQL, ";")

		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if err := db.Exec(stmt).Error; err != nil {
				fmt.Printf("❌ Failed to execute %s:\n   Error: %v\n", filename, err)
				break
			}
		}

		fmt.Printf("✅ %s\n", filename)
	}

	fmt.Println("\n✅ All migrations completed!")
}

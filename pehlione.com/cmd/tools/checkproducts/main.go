package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	ctx := context.Background()

	// Check categories - simple count
	type CategoryCheck struct {
		CategoryName string
		CategorySlug string
		Count        int
	}

	var categories []CategoryCheck
	if err := db.WithContext(ctx).
		Table("products").
		Select("category_name, category_slug, COUNT(*) as count").
		Where("status = ?", "active").
		Group("category_slug, category_name").
		Order("count DESC").
		Find(&categories).Error; err != nil {
		log.Fatalf("failed to query categories: %v", err)
	}

	fmt.Println("Categories in database (simple GROUP BY):")
	for i, c := range categories {
		fmt.Printf("  %d. Name: '%s', Slug: '%s', Count: %d\n", i+1, c.CategoryName, c.CategorySlug, c.Count)
	}

	// Test the exact query from repository
	fmt.Println("\nTest query (exact from repository):")
	var testCategories []struct {
		Slug  string
		Name  string
		Count int64
	}
	if err := db.WithContext(ctx).
		Table("products").
		Select("COALESCE(NULLIF(category_slug, ''), 'all') AS slug, COALESCE(NULLIF(category_name, ''), 'All products') AS name, COUNT(DISTINCT id) AS count").
		Where("status = ?", "active").
		Group("slug, name").
		Order("count DESC").
		Scan(&testCategories).Error; err != nil {
		log.Fatalf("failed to test query: %v", err)
	}

	fmt.Printf("Found %d result rows:\n", len(testCategories))
	for i, c := range testCategories {
		fmt.Printf("  %d. Slug: '%s', Name: '%s', Count: %d\n", i+1, c.Slug, c.Name, c.Count)
	}
}

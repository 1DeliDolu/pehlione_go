package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pehlione.com/app/internal/modules/products"
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
	repo := products.NewGormRepo(db)

	// Test ListFiltered with default filters
	filters := products.ListFilters{
		PageSize: 24,
		Page:     1,
		Sort:     "newest",
	}

	result, err := repo.ListFiltered(ctx, filters)
	if err != nil {
		log.Fatalf("ListFiltered failed: %v", err)
	}

	fmt.Printf("Total products: %d\n", result.Total)
	fmt.Printf("Returned: %d products\n", len(result.Items))
	fmt.Printf("Categories found: %d\n", len(result.Categories))
	fmt.Println("\nCategories:")
	for i, cat := range result.Categories {
		fmt.Printf("  %d. %s (%s): %d\n", i+1, cat.Name, cat.Slug, cat.Count)
	}
}

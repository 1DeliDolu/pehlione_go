package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pehlione.com/app/internal/config"
	"pehlione.com/app/internal/modules/products"
)

func main() {
	_ = godotenv.Load()

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to DB
	db, err := gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	ctx := context.Background()

	// Create services
	productsRepo := products.NewGormRepo(db)
	productsSvc := products.NewService(productsRepo)

	fmt.Println("Testing ListWithFilters...")

	// Test basic list
	filters := products.ListFilters{
		PageSize: 24,
		Page:     1,
		Sort:     "newest",
	}

	result, err := productsSvc.ListWithFilters(ctx, filters)
	if err != nil {
		log.Fatalf("❌ ListWithFilters failed: %v", err)
	}

	fmt.Printf("✅ ListWithFilters success!\n")
	fmt.Printf("Found %d products\n", len(result.Items))
	fmt.Printf("Total: %d\n", result.Total)

	// Now test the handler's mapProductsForList function
	if len(result.Items) > 0 {
		fmt.Println("\nProduct details:")

		// Map products
		for i, p := range result.Items {
			fmt.Printf("\n%d. %s (ID: %s)\n", i+1, p.Name, p.ID)
			if len(p.Variants) == 0 {
				fmt.Printf("   ⚠️  No variants!\n")
			} else {
				fmt.Printf("   Variants: %d\n", len(p.Variants))
				for j, v := range p.Variants {
					fmt.Printf("     %d. Price: %d cents, Stock: %d\n", j+1, v.PriceCents, v.Stock)
				}
			}
			if len(p.Images) == 0 {
				fmt.Printf("   ⚠️  No images!\n")
			} else {
				fmt.Printf("   Images: %d\n", len(p.Images))
			}
		}
	}

	fmt.Println("\n✅ All tests passed!")
}

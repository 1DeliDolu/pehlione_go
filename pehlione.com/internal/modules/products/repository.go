package products

import (
	"context"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	ListActive(ctx context.Context, limit, offset int) ([]Product, error)
	ListFiltered(ctx context.Context, filters ListFilters) (ListResult, error)
	GetBySlug(ctx context.Context, slug string) (Product, error)
	ListByIDs(ctx context.Context, ids []string) ([]Product, error)
}

type GormRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{db: db}
}

func (r *GormRepo) ListActive(ctx context.Context, limit, offset int) ([]Product, error) {
	if limit <= 0 || limit > 100 {
		limit = 24
	}
	var items []Product
	err := r.db.WithContext(ctx).
		Model(&Product{}).
		Where("status = ?", "active").
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order("position asc, id asc")
		}).
		Preload("Variants", func(db *gorm.DB) *gorm.DB {
			return db.Order("id asc")
		}).
		Order("id desc").
		Limit(limit).
		Offset(offset).
		Find(&items).Error
	return items, err
}

func (r *GormRepo) ListFiltered(ctx context.Context, filters ListFilters) (ListResult, error) {
	pageSize := filters.PageSize
	if pageSize <= 0 || pageSize > 60 {
		pageSize = 24
	}
	page := filters.Page
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query := r.db.WithContext(ctx).Table("products AS p").Where("p.status = ?", "active")

	if filters.Query != "" {
		like := "%" + escapeLike(filters.Query) + "%"
		query = query.Where("(p.name LIKE ? ESCAPE '\\\\' OR p.description LIKE ? ESCAPE '\\\\')", like, like)
	}
	if filters.Category != "" && filters.Category != "all" {
		query = query.Where("p.category_slug = ?", filters.Category)
	}
	if filters.MinPrice > 0 {
		query = query.Where("EXISTS (SELECT 1 FROM product_variants pv WHERE pv.product_id = p.id AND pv.price_cents >= ?)", filters.MinPrice)
	}
	if filters.MaxPrice > 0 {
		query = query.Where("EXISTS (SELECT 1 FROM product_variants pv WHERE pv.product_id = p.id AND pv.price_cents <= ?)", filters.MaxPrice)
	}
	if filters.InStock {
		query = query.Where("EXISTS (SELECT 1 FROM product_variants pv WHERE pv.product_id = p.id AND pv.stock > 0)")
	}

	priceSub := r.db.WithContext(ctx).
		Model(&Variant{}).
		Select("product_id, MIN(price_cents) AS min_price_cents").
		Group("product_id")
	query = query.Joins("LEFT JOIN (?) price_agg ON price_agg.product_id = p.id", priceSub)

	switch filters.Sort {
	case "price_asc":
		query = query.Order("price_agg.min_price_cents ASC").Order("p.created_at DESC")
	case "price_desc":
		query = query.Order("price_agg.min_price_cents DESC").Order("p.created_at DESC")
	default:
		query = query.Order("p.created_at DESC")
	}

	countQuery := query.Session(&gorm.Session{NewDB: true})
	var total int64
	if err := countQuery.Table("products AS p").Count(&total).Error; err != nil {
		return ListResult{}, err
	}

	rowsQuery := query.Select("p.id").
		Offset(offset).
		Limit(pageSize)

	var idRows []struct {
		ID string
	}
	if err := rowsQuery.Scan(&idRows).Error; err != nil {
		return ListResult{}, err
	}

	products := make([]Product, 0, len(idRows))
	if len(idRows) > 0 {
		ids := make([]string, 0, len(idRows))
		for _, row := range idRows {
			ids = append(ids, row.ID)
		}

		var dbProducts []Product
		if err := r.db.WithContext(ctx).
			Model(&Product{}).
			Where("id IN ?", ids).
			Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Order("position asc, id asc") }).
			Preload("Variants", func(db *gorm.DB) *gorm.DB { return db.Order("price_cents asc") }).
			Find(&dbProducts).Error; err != nil {
			return ListResult{}, err
		}

		orderMap := make(map[string]int, len(ids))
		for idx, id := range ids {
			orderMap[id] = idx
		}
		ordered := make([]Product, len(dbProducts))
		for _, p := range dbProducts {
			if idx, ok := orderMap[p.ID]; ok {
				ordered[idx] = p
			}
		}
		for _, p := range ordered {
			if p.ID != "" {
				products = append(products, p)
			}
		}
	}

	categories := make([]CategoryFacet, 0)
	// Use a subquery approach: get filtered product IDs, then count categories
	// This avoids GORM GROUP BY issues with NULL columns

	// Get distinct products matching filters
	var filteredProds []Product
	filterQ := r.db.WithContext(ctx).Table("products AS p").Where("p.status = ?", "active")

	if filters.Query != "" {
		like := "%" + escapeLike(filters.Query) + "%"
		filterQ = filterQ.Where("(p.name LIKE ? ESCAPE '\\\\' OR p.description LIKE ? ESCAPE '\\\\')", like, like)
	}
	if filters.Category != "" && filters.Category != "all" {
		filterQ = filterQ.Where("p.category_slug = ?", filters.Category)
	}
	if filters.MinPrice > 0 {
		filterQ = filterQ.Where("EXISTS (SELECT 1 FROM product_variants pv WHERE pv.product_id = p.id AND pv.price_cents >= ?)", filters.MinPrice)
	}
	if filters.MaxPrice > 0 {
		filterQ = filterQ.Where("EXISTS (SELECT 1 FROM product_variants pv WHERE pv.product_id = p.id AND pv.price_cents <= ?)", filters.MaxPrice)
	}
	if filters.InStock {
		filterQ = filterQ.Where("EXISTS (SELECT 1 FROM product_variants pv WHERE pv.product_id = p.id AND pv.stock > 0)")
	}

	if err := filterQ.Select("id, category_slug, category_name").Find(&filteredProds).Error; err == nil {
		// Count categories
		catCounts := make(map[string]struct {
			slug string
			name string
			cnt  int
		})
		for _, prod := range filteredProds {
			slug := prod.CategorySlug
			name := prod.CategoryName
			if slug == "" {
				slug = "all"
			}
			if name == "" {
				name = "All products"
			}
			key := slug + ":" + name
			entry := catCounts[key]
			entry.slug = slug
			entry.name = name
			entry.cnt++
			catCounts[key] = entry
		}

		// Build categories from map, maintaining count order
		type catCount struct {
			slug  string
			name  string
			count int
		}
		catList := make([]catCount, 0, len(catCounts))
		for _, entry := range catCounts {
			catList = append(catList, catCount{
				slug:  entry.slug,
				name:  entry.name,
				count: entry.cnt,
			})
		}

		// Sort by count descending
		for i := 0; i < len(catList); i++ {
			for j := i + 1; j < len(catList); j++ {
				if catList[j].count > catList[i].count {
					catList[i], catList[j] = catList[j], catList[i]
				}
			}
		}

		for _, cc := range catList {
			categories = append(categories, CategoryFacet{
				Slug:  cc.slug,
				Name:  cc.name,
				Count: int64(cc.count),
			})
		}
	}

	return ListResult{
		Items:      products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		Categories: categories,
	}, nil
}

func escapeLike(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "%", "\\%")
	s = strings.ReplaceAll(s, "_", "\\_")
	return s
}

func (r *GormRepo) GetBySlug(ctx context.Context, slug string) (Product, error) {
	var p Product
	err := r.db.WithContext(ctx).
		Model(&Product{}).
		Where("slug = ? AND status = ?", slug, "active").
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order("position asc, id asc")
		}).
		Preload("Variants", func(db *gorm.DB) *gorm.DB {
			return db.Order("id asc")
		}).
		First(&p).Error
	return p, err
}

func (r *GormRepo) ListByIDs(ctx context.Context, ids []string) ([]Product, error) {
	if len(ids) == 0 {
		return []Product{}, nil
	}
	var items []Product
	err := r.db.WithContext(ctx).
		Model(&Product{}).
		Where("id IN ?", ids).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Order("position asc, id asc") }).
		Preload("Variants", func(db *gorm.DB) *gorm.DB { return db.Order("price_cents asc") }).
		Find(&items).Error
	return items, err
}

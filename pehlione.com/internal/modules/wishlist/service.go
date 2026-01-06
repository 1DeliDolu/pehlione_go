package wishlist

import (
	"context"

	"gorm.io/gorm"
)

type Service struct {
	repo *Repo
	db   *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		repo: NewRepo(db),
		db:   db,
	}
}

func (s *Service) Add(ctx context.Context, userID, productID string) error {
	return s.repo.Add(ctx, userID, productID)
}

func (s *Service) Remove(ctx context.Context, userID, productID string) error {
	return s.repo.Remove(ctx, userID, productID)
}

func (s *Service) Items(ctx context.Context, userID string) ([]Item, error) {
	return s.repo.ListByUser(ctx, userID)
}

func (s *Service) Contains(ctx context.Context, userID, productID string) (bool, error) {
	return s.repo.Contains(ctx, userID, productID)
}

package service

import (
	"context"

	"github.com/Nuttachai-K/cafe-finder-api/internal/model"
	"github.com/Nuttachai-K/cafe-finder-api/internal/repository"
)

type CafeService struct {
	repo repository.CafeRepository
}

func NewCafeService(repo repository.CafeRepository) *CafeService {
	return &CafeService{
		repo: repo,
	}
}

func (s *CafeService) GetByID(ctx context.Context, id int) (*model.Cafe, error) {
	// TODO
	return s.repo.GetByID(ctx, id)
}

func (s *CafeService) Create(ctx context.Context, cafe *model.Cafe) error {
	// TODO
	return s.repo.Create(ctx, cafe)
}

func (s *CafeService) GetAll(ctx context.Context) ([]model.Cafe, error) {
	//TODO
	return s.repo.GetAll(ctx)
}

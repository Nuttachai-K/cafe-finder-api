package service

import (
	"github.com/Nuttachai-K/cafe-finder-api/internal/model"
	"github.com/Nuttachai-K/cafe-finder-api/internal/repository"
)

type CafeService struct {
	repo repository.CafeRespository
}

func (s *CafeService) GetByID(id int) (*model.Cafe, error) {
	// TODO
	return s.repo.GetByID(id)
}

func (s *CafeService) Create(cafe *model.Cafe) error {
	// TODO
	return s.repo.Create(cafe)
}

func (s *CafeService) GetAll() ([]model.Cafe, error) {
	//TODO
	return s.repo.GetAll()
}

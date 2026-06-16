package repository

import (
	"database/sql"

	"github.com/Nuttachai-K/cafe-finder-api/internal/model"
)

// CafeResoisitory defines database operation for cafe
type CafeRespository interface {
	GetByID(id int) (*model.Cafe, error)
	GetAll() ([]model.Cafe, error)
	Create(cafe *model.Cafe) error
}

type cafeRespository struct {
	db *sql.DB
}

func NewCafeRespository(db *sql.DB) CafeRespository {
	return &cafeRespository{
		db: db,
	}
}

func (c cafeRespository) GetByID(id int) (*model.Cafe, error) {
	return &model.Cafe{}, nil
}

func (c cafeRespository) GetAll() ([]model.Cafe, error) {
	return []model.Cafe{}, nil
}

func (c cafeRespository) Create(cafe *model.Cafe) error {
	return nil
}

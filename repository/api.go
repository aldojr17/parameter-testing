package repository

import (
	"parameter-testing/domain/entity"

	"gorm.io/gorm"
)

type APIRepository struct {
	db *gorm.DB
}

func NewAPIRepository(db *gorm.DB) *APIRepository {
	return &APIRepository{
		db: db,
	}
}

func (r *APIRepository) Create(payload *entity.API) error {
	return r.db.Create(payload).Error
}

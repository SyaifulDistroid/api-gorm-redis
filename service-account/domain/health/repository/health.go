package repository

import "gorm.io/gorm"

type HealthRepository interface {
}

type healthRepository struct {
	db *gorm.DB
}

func NewHealthRepository(db *gorm.DB) HealthRepository {
	return &healthRepository{
		db: db,
	}
}

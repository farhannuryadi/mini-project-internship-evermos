package repositories

import "gorm.io/gorm"

type AlamatRepo struct {
	db *gorm.DB
}

func NewAlamatRepo(db *gorm.DB) *AlamatRepo {
	return &AlamatRepo{db}
}
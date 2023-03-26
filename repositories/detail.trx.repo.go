package repositories

import "gorm.io/gorm"

type DetailTrxRepo struct {
	db *gorm.DB
}

func NewDetailTrxRepo(db *gorm.DB) *DetailTrxRepo {
	return &DetailTrxRepo{db}
}
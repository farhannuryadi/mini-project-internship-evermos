package repositories

import "gorm.io/gorm"

type LogProdukRepo struct {
	db *gorm.DB
}

func NewLogProdukRepo(db *gorm.DB) *LogProdukRepo {
	return &LogProdukRepo{db}
}
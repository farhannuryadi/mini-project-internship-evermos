package repositories

import "gorm.io/gorm"

type ProdukRepo struct {
	db *gorm.DB
}

func NewProdukRepo(db *gorm.DB) *ProdukRepo {
	return &ProdukRepo{db}
}
package repositories

import "gorm.io/gorm"

type FotoProdukRepo struct {
	db *gorm.DB
}

func NewFotoProdukRepo(db *gorm.DB) *FotoProdukRepo {
	return &FotoProdukRepo{db}
}
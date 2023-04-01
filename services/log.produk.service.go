package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type LogProdukService struct {
	repo *repositories.LogProdukRepo
}

func NewLogProdukService() *LogProdukService {
	return &LogProdukService{&repositories.LogProdukRepo{}}
}

func (s *LogProdukService) Create(produk entity.Produk) (entity.LogProduk, error) {
	return s.repo.Create(produk)
}

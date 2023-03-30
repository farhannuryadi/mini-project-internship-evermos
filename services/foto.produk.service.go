package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"

)

type FotoProdukService struct {
	repo *repositories.FotoProdukRepo
}

func NewFotoProdukService() *FotoProdukService {
	return &FotoProdukService{&repositories.FotoProdukRepo{}}
}

func (s *FotoProdukService) Create(fotoProduk entity.FotoProduk) error {
	return s.repo.Create(fotoProduk)
}

func (s *FotoProdukService) GetByProdukId(produkId string) ([]entity.FotoProduk, error) {
	return s.repo.FindByProdukId(produkId)
}

func (s *FotoProdukService) Update(fotoProduk entity.FotoProduk) (entity.FotoProduk, error) {
	return s.repo.Update(fotoProduk)
}

func (s *FotoProdukService) Delete(fotoProduk entity.FotoProduk) (bool, error) {
	return s.repo.Delete(fotoProduk)
}
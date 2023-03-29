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

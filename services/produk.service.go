package services

import "mini-project-internship/repositories"

type ProdukService struct {
	repo *repositories.ProdukRepo
}

func NewProdukService(repo *repositories.ProdukRepo) *ProdukService {
	return &ProdukService{repo}
}
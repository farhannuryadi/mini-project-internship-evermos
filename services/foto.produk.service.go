package services

import "mini-project-internship/repositories"

type FotoProdukServcie struct {
	repo *repositories.FotoProdukRepo
}

func NewFotoProdukService(repo *repositories.FotoProdukRepo) *FotoProdukServcie {
	return &FotoProdukServcie{repo}
}
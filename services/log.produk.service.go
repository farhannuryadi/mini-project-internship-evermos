package services

import "mini-project-internship/repositories"

type LogProdukService struct {
	repo *repositories.LogProdukRepo
}

func NewLogProdukService(repo *repositories.LogProdukRepo) *LogProdukService {
	return &LogProdukService{repo}
}
package services

import "mini-project-internship/repositories"

type DetailTrxService struct {
	repo *repositories.DetailTrxRepo
}

func NewDetailTrxService(repo *repositories.DetailTrxRepo) *DetailTrxService {
	return &DetailTrxService{repo}
}
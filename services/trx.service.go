package services

import "mini-project-internship/repositories"

type TrxService struct {
	repo *repositories.TrxRepo
}

func NewTrxService(repo *repositories.TrxRepo) *TrxService {
	return &TrxService{repo}
}
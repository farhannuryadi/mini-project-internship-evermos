package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type DetailTrxService struct {
	repo *repositories.DetailTrxRepo
}

func NewDetailTrxService() *DetailTrxService {
	return &DetailTrxService{&repositories.DetailTrxRepo{}}
}

func (s *DetailTrxService) Create(detailTrx entity.DetailTrx) (entity.DetailTrx, error) {
	return s.repo.Create(detailTrx)
}

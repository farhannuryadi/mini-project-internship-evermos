package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type KotaService struct {
	repo *repositories.KotaRepo
}

func NewKotaService() *KotaService {
	return &KotaService{&repositories.KotaRepo{}}
}

func (s *KotaService) GetById(kotaId string) (entity.Kota, error) {
	return s.repo.FindById(kotaId)
}

func (s *KotaService) GetByProvinsiId(provinsiId string) ([]entity.Kota, error) {
	return s.repo.FindByProvinsiId(provinsiId)
}

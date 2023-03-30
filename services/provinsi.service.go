package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type ProvinsiService struct {
	repo *repositories.ProvinsiRepo
}

func NewProvinsiService() *ProvinsiService {
	return &ProvinsiService{&repositories.ProvinsiRepo{}}
}

func (s *ProvinsiService) GetById(provinsiId string) (entity.Provinsi, error) {
	return s.repo.FindById(provinsiId)
}

func (s *ProvinsiService) GetAll() ([]entity.Provinsi, error) {
	return s.repo.FindAll()
}

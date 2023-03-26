package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/repositories"
)

type TokoService struct {
	repo *repositories.TokoRepo
}

func NewTokoService() *TokoService {
	return &TokoService{&repositories.TokoRepo{}}
}

func (s *TokoService) Create(tokoReq request.TokoReq) error {
	return s.repo.Create(tokoReq)
}

func (s *TokoService) GetById(tokoId string) (entity.Toko, error) {
	return s.repo.FindById(tokoId)
}

func (s *TokoService) Update(tokoId string, tokoUpdate request.TokoUpdateReq) (entity.Toko, error) {
	return s.repo.Update(tokoId, tokoUpdate)
}

func (s *TokoService) GetAll() ([]entity.Toko, error) {
	return s.repo.FindAll()
}

func (s *TokoService) Delete(tokoId string) (bool, error) {
	return s.repo.Delete(tokoId)
}
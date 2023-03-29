package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type AlamatService struct {
	repo *repositories.AlamatRepo
}

func NewAlamatService(repo *repositories.AlamatRepo) *AlamatService {
	return &AlamatService{&repositories.AlamatRepo{}}
}

func (s *AlamatService) Create(alamat entity.Alamat) error {
	return s.repo.Create(alamat)
}

func (s *AlamatService) GetById(alamatId string) (entity.Alamat, error) {
	return s.repo.FindById(alamatId)
}

func (s *AlamatService) GetByUserId(userId string) ([]entity.Alamat, error) {
	return s.repo.FindByUserId(userId)
}

func (s *AlamatService) Update(alamatId string, update entity.Alamat) (entity.Alamat, error) {
	return s.repo.Update(alamatId, update)
}

func (s *AlamatService) Delete(alamatId string) (bool, error) {
	return s.repo.Delete(alamatId)
}

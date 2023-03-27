package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/repositories"
)

type CategoryServcie struct {
	repo *repositories.CategoryRepo
}

func NewCategoryService() *CategoryServcie {
	return &CategoryServcie{&repositories.CategoryRepo{}}
}

func (s *CategoryServcie) Create(categoryReq request.CategoryReq) error {
	return s.repo.Create(categoryReq)
}

func (s *CategoryServcie) GetById(categoryId string) (entity.Category, error) {
	return s.repo.FindById(categoryId)
}

func (s *CategoryServcie) Update(categoryId string, categoryReq request.CategoryReq) (entity.Category, error) {
	return s.repo.Update(categoryId, categoryReq)
}

func (s *CategoryServcie) GetAll() ([]entity.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryServcie) Delete(categoryId string) (bool, error) {
	return s.repo.Delete(categoryId)
}
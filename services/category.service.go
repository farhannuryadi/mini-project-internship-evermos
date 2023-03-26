package services

import "mini-project-internship/repositories"

type CategoryServcie struct {
	repo *repositories.CategoryRepo
}

func NewCategoryService(repo *repositories.CategoryRepo) *CategoryServcie {
	return &CategoryServcie{repo}
}
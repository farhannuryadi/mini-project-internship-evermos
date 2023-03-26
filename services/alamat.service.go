package services

import "mini-project-internship/repositories"

type AlamatService struct {
	repo *repositories.AlamatRepo
}

func NewAlamatService(repo *repositories.AlamatRepo) *AlamatService {
	return &AlamatService{repo}
}
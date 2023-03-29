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

func (s *TokoService) Update(tokoId string, tokoUpdate entity.Toko) (entity.Toko, error) {
	return s.repo.Update(tokoId, tokoUpdate)
}

func (s *TokoService) GetAll() ([]entity.Toko, error) {
	return s.repo.FindAll()
}

func (s *TokoService) Delete(tokoId string) (bool, error) {
	return s.repo.Delete(tokoId)
}

func (s *TokoService) GetByUserId(userId string) (entity.Toko, error) {
	return s.repo.FindByUserId(userId)
}

func (s *TokoService) GetAllPage(page, perPage uint64, search string) (map[string]interface{}, error) {
	totalData, err := s.repo.Count()
	if err != nil {
		return nil, err
	}
	stores, err := s.repo.FindAllPage(page, perPage, search)
	if err != nil {
		return nil, err
	}
	totalPage := getTotalPage(uint64(totalData), perPage)
	return map[string]interface{}{
		"total_data":   totalData,
		"total_page":   totalPage,
		"current_page": page,
		"data":         stores,
	}, nil
}

func getTotalPage(count, perPage uint64) uint64 {
	totalPage := count / perPage
	if count%perPage != 0 {
		totalPage++
	}
	return totalPage
}

package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type TrxService struct {
	repo *repositories.TrxRepo
}

func NewTrxService() *TrxService {
	return &TrxService{&repositories.TrxRepo{}}
}

func (s *TrxService) Create(trx entity.Trx) (entity.Trx, error) {
	return s.repo.Create(trx)
}

func (s *TrxService) GetAllPage(userId, search string, page, limit uint64) (map[string]interface{}, error) {
	totalData, err := s.repo.Count(userId)
	if err != nil {
		return nil, err
	}
	stores, err := s.repo.FindAllPage(userId, search, page, limit)
	if err != nil {
		return nil, err
	}
	totalPage := getTotalPage(uint64(totalData), limit)
	return map[string]interface{}{
		"total_data":   totalData,
		"total_page":   totalPage,
		"current_page": page,
		"data":         stores,
	}, nil
}

func (s *TrxService) GetById(trxId string) (entity.Trx, error) {
	return s.repo.FindById(trxId)
}

package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/repositories"
)

type ProdukService struct {
	repo *repositories.ProdukRepo
}

func NewProdukService() *ProdukService {
	return &ProdukService{&repositories.ProdukRepo{}}
}

func (s *ProdukService) Create(produk entity.Produk) (entity.Produk, error) {
	return s.repo.Create(produk)
}

func (s *ProdukService) GetAllByFilter(
	namaProduk string, categoryID string, tokoID string,
	maxHarga string, minHarga string, limit uint, page uint) (map[string]interface{}, error) {

	totalData, err := s.repo.Count()
	if err != nil {
		return nil, err
	}
	produks, err := s.repo.FindAllPage(namaProduk, categoryID, tokoID, maxHarga, minHarga, limit, page)
	if err != nil {
		return nil, err
	}
	totalPage := getTotalPage(uint64(totalData), uint64(limit))
	return map[string]interface{}{
		"total_data": totalData,
		"total_page": totalPage,
		"page":       page,
		"limit":      limit,
		"data":       produks,
	}, nil
}

package controllers

import "mini-project-internship/services"

type ProdukController struct {
	service *services.ProdukService
}

func NewProdukController(service *services.ProdukService) *ProdukController {
	return &ProdukController{service}
}
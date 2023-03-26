package controllers

import "mini-project-internship/services"

type FotoProdukController struct {
	service *services.FotoProdukServcie
}

func NewFotoProdukController(service *services.FotoProdukServcie) *FotoProdukController {
	return &FotoProdukController{service}
}
package controllers

import "mini-project-internship/services"

type LogProdukController struct {
	service *services.LogProdukService
}

func NewLogProdukController(service *services.LogProdukService) *LogProdukController {
	return &LogProdukController{service}
}
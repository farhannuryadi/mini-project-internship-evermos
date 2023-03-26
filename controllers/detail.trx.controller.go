package controllers

import "mini-project-internship/services"

type DetailTrxController struct {
	service *services.DetailTrxService
}

func NewDetailTrxController(service *services.DetailTrxService) *DetailTrxController {
	return &DetailTrxController{service}
}
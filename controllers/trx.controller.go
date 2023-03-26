package controllers

import "mini-project-internship/services"

type TrxController struct {
	service *services.TrxService
}

func NewTrxController(service *services.TrxService) *TrxController {
	return &TrxController{service}
}
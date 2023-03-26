package controllers

import "mini-project-internship/services"

type AlamatController struct {
	service *services.AlamatService
}

func NewAlamatController(service *services.AlamatService) *AlamatController {
	return &AlamatController{service}
}
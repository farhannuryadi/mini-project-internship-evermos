package controllers

import "mini-project-internship/services"

type CategoryController struct {
	service *services.CategoryServcie
}

func NewCategoryController(service *services.CategoryServcie) *CategoryController {
	return &CategoryController{service}
}
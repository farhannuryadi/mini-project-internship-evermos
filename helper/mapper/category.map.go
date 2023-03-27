package mapper

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/models/response"
)

func CategoryReqToCategory(req request.CategoryReq) entity.Category {
	return entity.Category{
		NamaCategory: req.NamaCategory,
	}
}

func CategoryUpdateToCategory(req request.CategoryReq, category entity.Category) entity.Category {
	if req.NamaCategory != "" {
		category.NamaCategory = req.NamaCategory
	}
	return category
}

func CategoryToCategoryRes(category entity.Category) response.CategoryRes {
	return response.CategoryRes{
		ID: category.ID,
		NamaCategory: category.NamaCategory,
	}
}
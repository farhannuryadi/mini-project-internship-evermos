package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
)

type CategoryRepo struct {
}

func (r *CategoryRepo) Create(categoryReq request.CategoryReq) error {
	var category entity.Category = mapper.CategoryReqToCategory(categoryReq)
	return database.DB.Debug().Create(&category).Error
}

func (r *CategoryRepo) FindById(categoryId string) (entity.Category, error) {
	var category entity.Category
	if err := database.DB.Debug().First(&category, "id = ?", categoryId).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *CategoryRepo) Update(categoryId string, categoryReq request.CategoryReq) (entity.Category, error) {
	var category entity.Category
	category, err := r.FindById(categoryId)
	if err != nil {
		return category, err
	}
	
	category = mapper.CategoryUpdateToCategory(categoryReq, category)

	errUpdate := database.DB.Debug().Save(&category).Error
	if errUpdate != nil {
		return category, errUpdate
	}
	return category, nil
}

func (r *CategoryRepo) FindAll() ([]entity.Category, error) {
	var category []entity.Category
	err := database.DB.Debug().Find(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *CategoryRepo) Delete(categoryId string) (bool, error) {
	var category entity.Category
	category, err := r.FindById(categoryId)
	if err != nil {
		return false, err
	}
	errDel := database.DB.Debug().Delete(&category).Error
	if errDel != nil {
		return false, err
	}
	return true, nil
}
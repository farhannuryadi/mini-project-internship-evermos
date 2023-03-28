package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"
)

type ProvinsiRepo struct {
}

func (r *ProvinsiRepo) Create(provinsi entity.Provinsi) error {
	return database.DB.Debug().Create(&provinsi).Error
}

func (r *ProvinsiRepo) FindById(provinsiId string) (entity.Provinsi, error) {
	var provinsi entity.Provinsi
	if err := database.DB.Debug().First(&provinsi, "id = ?", provinsiId).Error; err != nil {
		return provinsi, err
	}
	return provinsi, nil
}

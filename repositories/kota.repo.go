package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"

)

type KotaRepo struct {
}

func (r *KotaRepo) Create(kota entity.Kota) error {
	return database.DB.Debug().Create(&kota).Error
}

func (r *KotaRepo) FindById(kotaId string) (entity.Kota, error) {
	var kota entity.Kota
	if err := database.DB.Debug().First(&kota, "id = ?", kotaId).Error; err != nil {
		return kota, err
	}
	return kota, nil
}
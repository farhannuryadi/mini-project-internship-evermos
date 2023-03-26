package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
)

type TokoRepo struct {
}

func (r *TokoRepo) Create(tokoReq request.TokoReq) error {
	var toko entity.Toko = mapper.TokoReqToToko(tokoReq)
	return database.DB.Debug().Create(&toko).Error
}

func (r *TokoRepo) FindById(tokoId string) (entity.Toko, error) {
	var toko entity.Toko
	if err := database.DB.Debug().First(&toko, "id = ?", tokoId).Error; err != nil {
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepo) Update(tokoId string, tokoUpdate request.TokoUpdateReq) (entity.Toko, error) {
	var toko entity.Toko
	toko, err := r.FindById(tokoId)
	if err != nil {
		return toko, err
	}

	toko = mapper.TokoUpdateToToko(tokoUpdate, toko)

	errUpdate := database.DB.Debug().Save(&toko).Error
	if errUpdate != nil {
		return toko, err
	}

	return toko, nil
}  

func (r *TokoRepo) FindAll() ([]entity.Toko, error) {
	var toko []entity.Toko
	err := database.DB.Debug().Find(&toko).Error
	if err != nil {
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepo) Delete(tokoId string) (bool, error) {
	var toko entity.Toko
	toko, err := r.FindById(tokoId)
	if err != nil {
		return false, err
	}
	
	errDel := database.DB.Debug().Delete(&toko)
	if errDel != nil {
		return false, err
	}
	return true, nil
}
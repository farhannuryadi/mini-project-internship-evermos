package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"

)

type AlamatRepo struct {
}

func (r *AlamatRepo) Create(alamat entity.Alamat) error {
	return database.DB.Debug().Create(&alamat).Error
}

func (r *AlamatRepo) FindById(alamatId string) (entity.Alamat, error) {
	var alamat entity.Alamat
	err := database.DB.Debug().First(&alamat, "id = ?", alamatId).Error
	if err != nil {
		return alamat, err
	}
	return alamat, nil
}

func (r *AlamatRepo) FindByUserId(userId string) ([]entity.Alamat, error) {
	var alamat []entity.Alamat
	err := database.DB.Debug().Raw("SELECT * FROM `alamats` WHERE user_id = ?", userId).Scan(&alamat).Error
	if err != nil {
		return alamat, err
	}
	return alamat, nil
}

func (r *AlamatRepo) Update(alamatId string, update entity.Alamat) (entity.Alamat, error) {
	var alamat entity.Alamat
	alamat, err := r.FindById(alamatId)
	if err != nil {
		return alamat, err
	}
	alamat = mapper.AlamatUpdateToAlamat(update, alamat)
	errUpdate := database.DB.Debug().Save(&alamat).Error
	if errUpdate != nil {
		return alamat, errUpdate
	}
	return alamat, nil
}

func (r *AlamatRepo) Delete(alamatId string) (bool, error) {
	alamat, err := r.FindById(alamatId)
	if err != nil {
		return false, err
	}

	errDelete := database.DB.Debug().Delete(&alamat).Error
	if errDelete != nil {
		return false, err
	}

	return true, nil
}

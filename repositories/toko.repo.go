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
	if err := database.DB.Debug().First(&toko, tokoId).Error; err != nil {
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepo) Update(tokoId string, tokoUpdate entity.Toko) (entity.Toko, error) {
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

func (r *TokoRepo) FindByUserId(userId string) (entity.Toko, error) {
	var toko entity.Toko
	err := database.DB.Debug().Raw("SELECT * FROM `tokos` WHERE user_id = ?", userId).Scan(&toko).Error
	if err != nil {
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepo) Count() (int64, error) {
	var toko entity.Toko
	var count int64
	err := database.DB.Model(&toko).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TokoRepo) FindAllPage(page, perPage uint64, search string) ([]entity.Toko, error) {
	var tokos []entity.Toko
	offset := (page - 1) * perPage
	err := database.DB.Raw("SELECT * FROM tokos WHERE nama_toko LIKE ? ORDER BY id DESC LIMIT ? OFFSET ?", "%"+search+"%", perPage, offset).Scan(&tokos).Error
	if err != nil {
		return nil, err
	}
	return tokos, nil
}

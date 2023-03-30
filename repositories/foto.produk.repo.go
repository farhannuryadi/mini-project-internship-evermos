package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"
)

type FotoProdukRepo struct {
}

func (r *FotoProdukRepo) Create(fotoProduk entity.FotoProduk) error {
	err := database.DB.Debug().Create(&fotoProduk).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *FotoProdukRepo) FindByProdukId(produkId string) ([]entity.FotoProduk, error) {
	var photos []entity.FotoProduk
	err := database.DB.Debug().Where("produk_id = ?", produkId).Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil
}

func (r *FotoProdukRepo) Update(fotoProduk entity.FotoProduk) (entity.FotoProduk, error) {
	err := database.DB.Debug().Save(&fotoProduk).Error
	if err != nil {
		return fotoProduk, err
	}
	return fotoProduk, nil
}

func (r *FotoProdukRepo) Delete(fotoProduk entity.FotoProduk) (bool, error) {
	err := database.DB.Debug().Delete(&fotoProduk).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

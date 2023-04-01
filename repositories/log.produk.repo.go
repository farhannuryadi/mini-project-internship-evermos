package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
)

type LogProdukRepo struct {
}

func (r *LogProdukRepo) Create(produk entity.Produk) (entity.LogProduk, error) {

	logProduk := mapper.ProdukToLog(produk)
	err := database.DB.Debug().Create(&logProduk).Error
	if err != nil {
		return logProduk, err
	}
	return logProduk, nil
}

// func FindById()  {

// }

// func Update()  {

// }

// func FindAll()  {

// }

// func Delete()  {

// }

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

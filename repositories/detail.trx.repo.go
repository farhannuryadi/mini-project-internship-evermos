package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"
)

type DetailTrxRepo struct {
}

func (r *DetailTrxRepo) Create(detailTrx entity.DetailTrx) (entity.DetailTrx, error) {
	err := database.DB.Debug().Create(&detailTrx).Error
	if err != nil {
		return detailTrx, err
	}
	return detailTrx, nil
}

// func FindById()  {

// }

// func Update()  {

// }

// func FindAll()  {

// }

// func Delete()  {

// }

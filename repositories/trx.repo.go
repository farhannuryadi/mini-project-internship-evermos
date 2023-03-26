package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/models/entity"
)

type TrxRepo struct {
}

func (r *TrxRepo) Create(trx entity.Trx) error {
	return database.DB.Debug().Create(trx).Error
}

// func FindById()  {
	
// }

// func Update()  {
	
// }

// func FindAll()  {
	
// }

// func Delete()  {
	
// }
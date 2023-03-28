package repositories

import (
	"mini-project-internship/database"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"

)

type UserRepo struct {
}

func (r *UserRepo) Create(user entity.User) (entity.User, error) {
	err := database.DB.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) FindById(userId string) (entity.User, error) {
	var user entity.User
	if err := database.DB.Debug().First(&user, "id = ?", userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) Update(userId string, userUpdate request.UserUpdateReq) (entity.User, error) {
	var user entity.User
	user, err := r.FindById(userId)
	if err != nil {
		return user, err
	}

	user = mapper.UserUpdateToUser(userUpdate, user)

	errUpdate := database.DB.Debug().Save(&user).Error
	if errUpdate != nil {
		return user, errUpdate
	}
	return user, nil
}

func (r *UserRepo) FindAll() ([]entity.User, error) {
	var user []entity.User
	err := database.DB.Debug().Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) Delete(userId string) (bool, error) {
	var user entity.User
	user, err := r.FindById(userId)
	if err != nil {
		return false, err
	}
	errDel := database.DB.Debug().Delete(&user).Error
	if errDel != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepo) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := database.DB.Debug().First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) FindByNoTelp(noTelp string) (entity.User, error) {
	var user entity.User
	if err := database.DB.Debug().First(&user, "no_telp = ?", noTelp).Error; err != nil {
		return user, err
	}
	return user, nil
}

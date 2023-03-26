package services

import (
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/repositories"
	"mini-project-internship/utils"

)

type UserService struct {
	repo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{&repositories.UserRepo{}}
}

func (s *UserService) Create(userReq request.UserReq) error {
	hashedPassword, err := utils.HashingPassword(userReq.KataSandi)
	if err != nil {
		return err
	}

	userReq.KataSandi = hashedPassword
	return s.repo.Create(userReq)
}

func (s *UserService) GetById(userId string) (entity.User, error) {
	return s.repo.FindById(userId)
}

func (s *UserService) Update(userId string, userUpdate request.UserUpdateReq) (entity.User, error) {
	return s.repo.Update(userId, userUpdate)
}

func (s *UserService) GetAll() ([]entity.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Delete(userId string) (bool, error) {
	return s.repo.Delete(userId)
}

func (s *UserService) GetByEmail(email string) (entity.User, error) {
	return s.repo.FindByEmail(email)
}
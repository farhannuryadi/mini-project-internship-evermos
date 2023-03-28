package services

import (
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/repositories"
	"mini-project-internship/utils"

)

type AuthService struct {
	repo *repositories.UserRepo
}

func NewAuthService() *UserService {
	return &UserService{&repositories.UserRepo{}}
}

func (s *AuthService) Create(regisReq request.RegisterReq) (entity.User, error) {
	var user entity.User
	hashedPassword, err := utils.HashingPassword(regisReq.KataSandi)
	if err != nil {
		return user, err
	}

	regisReq.KataSandi = hashedPassword
	user = mapper.RegisToUser(regisReq)
	return s.repo.Create(user)
}

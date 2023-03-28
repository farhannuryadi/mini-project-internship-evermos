package controllers

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/models/response"
	"mini-project-internship/services"

)

var userService services.UserService = *services.NewUserService()

func UserCreate(ctx *fiber.Ctx) error {
	var userReq request.UserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	_ ,err := userService.Create(userReq);
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, nil)
}

func UserGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, err := userService.GetById(userId)
	var user1 entity.User = user
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	userRes := mapper.UserToUserRes(user1)

	return helper.SuccessHelper(ctx, fiber.StatusOK, userRes)
}

func UserUpdate(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var userReq request.UserUpdateReq

	if err := ctx.BodyParser(&userReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	user, errUpdate := userService.Update(userId, userReq)
	if errUpdate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errUpdate)
	}

	userRes := mapper.UserToUserRes(user)

	return helper.SuccessHelper(ctx, fiber.StatusAccepted, userRes)
}

func UserGetAll(ctx *fiber.Ctx) error {
	var userRes []response.UserRes
	users, err := userService.GetAll()
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	for _, user := range users{
		userRes = append(userRes, mapper.UserToUserRes(user))
	}
	
	return helper.SuccessHelper(ctx, fiber.StatusOK, userRes)
}

func UserDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	ket, err := userService.Delete(userId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusAccepted, ket)
}
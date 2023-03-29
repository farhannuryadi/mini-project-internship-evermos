package controllers

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/models/entity"
	"mini-project-internship/services"
)

var alamatService services.AlamatService = services.AlamatService{}

func AlamatCreate(ctx *fiber.Ctx) error {
	var alamat entity.Alamat
	if err := ctx.BodyParser(&alamat); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	userInfo, isOk := helper.ClaimsUserInfo(ctx)
	if isOk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, isOk)
	}

	userId, _ := userInfo["id"].(float64)
	alamat.UserID = uint(userId)

	validate := validator.New()
	errValidate := validate.Struct(alamat)
	if errValidate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errValidate)
	}

	err := alamatService.Create(alamat)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, 1)
}

func AlamatGetById(ctx *fiber.Ctx) error {
	log.Println("AlamatGetById")
	alamatId := ctx.Params("id")

	alamat, err := alamatService.GetById(alamatId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, alamat)
}

func AlamatGetByUserId(ctx *fiber.Ctx) error {
	userInfo, isOk := helper.ClaimsUserInfo(ctx)
	if isOk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, isOk)
	}

	userId, _ := userInfo["id"].(float64)
	id := fmt.Sprintf("%.0f", userId)

	alamats, err := alamatService.GetByUserId(id)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, alamats)
}

func AlamatUpdate(ctx *fiber.Ctx) error {
	alamatId := ctx.Params("id")
	var alamatReq entity.Alamat

	if err := ctx.BodyParser(&alamatReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	validate := validator.New()
	errValidate := validate.Struct(alamatReq)
	if errValidate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errValidate)
	}

	_, errUpdate := alamatService.Update(alamatId, alamatReq)
	if errUpdate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errUpdate)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, "")
}

func AlamatDelete(ctx *fiber.Ctx) error {
	alamatId := ctx.Params("id")

	_, err := alamatService.Delete(alamatId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, "")
}

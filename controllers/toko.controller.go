package controllers

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/services"
)

var tokoService services.TokoService = *services.NewTokoService()

func TokoCreate(ctx *fiber.Ctx) error {
	var tokoReq request.TokoReq
	if err := ctx.BodyParser(&tokoReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	if err := tokoService.Create(tokoReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusCreated, nil)
}

func TokoGetById(ctx *fiber.Ctx) error {
	var tokoId = ctx.Params("id")
	
	toko, err := tokoService.GetById(tokoId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, toko)
}

func TokoUpdate(ctx *fiber.Ctx) error {
	var tokoId = ctx.Params("id")
	var tokoReq request.TokoUpdateReq

	if err := ctx.BodyParser(&tokoReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	toko, errUpdate := tokoService.Update(tokoId, tokoReq)
	if errUpdate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errUpdate)
	}

	return helper.SuccessHelper(ctx, fiber.StatusAccepted, toko)
}

func TokoGetAll(ctx *fiber.Ctx) error {
	var tokos []entity.Toko
	tokos, err := tokoService.GetAll()
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, tokos)
}

func TokoDelete(ctx *fiber.Ctx) error {
	var tokoId = ctx.Params("id")

	ket, err := tokoService.Delete(tokoId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, ket)
}
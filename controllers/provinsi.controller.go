package controllers

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
)

func ProvinsiGetAll(ctx *fiber.Ctx) error {

	provinces, err := provinsiService.GetAll()
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, provinces)
}

func ProvinsiGetById(ctx *fiber.Ctx) error {
	provinsiId := ctx.Params("prov_id")

	province, err := provinsiService.GetById(provinsiId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, province)
}

package controllers

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
)

func KotaGetByProvinsiId(ctx *fiber.Ctx) error {
	provinsiId := ctx.Params("prov_id")

	citys, err := kotaService.GetByProvinsiId(provinsiId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, citys)
}

func KotaGetById(ctx *fiber.Ctx) error {
	kotaId := ctx.Params("city_id")

	city, err := kotaService.GetById(kotaId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, city)
}

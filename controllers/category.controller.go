package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/models/request"
	"mini-project-internship/services"
)

var categoryServcie services.CategoryServcie = *services.NewCategoryService()

func CategoryCreate(ctx *fiber.Ctx) error {
	var categoryReq request.CategoryReq
	if err := ctx.BodyParser(&categoryReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	if err := categoryServcie.Create(categoryReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, 1)
}

func CategoryGetById(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("id")
	category, err := categoryServcie.GetById(categoryId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, fmt.Errorf("No Data Category"))
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, category)
}

func CategoryUpdate(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("id")
	var categoryReq request.CategoryReq
	if err := ctx.BodyParser(&categoryReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	category, errUpdate := categoryServcie.Update(categoryId, categoryReq)
	if errUpdate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errUpdate)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, category)
}

func CategoryGetAll(ctx *fiber.Ctx) error {
	category, err := categoryServcie.GetAll()
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, category)
}

func CategoryDelete(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("id")

	ket, err := categoryServcie.Delete(categoryId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, ket)
}

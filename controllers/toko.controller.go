package controllers

import (
	"fmt"
	"strconv"

	// "log"

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
	var tokoId = ctx.Params("id_toko")

	toko, err := tokoService.GetById(tokoId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, toko)
}

func TokoUpdate(ctx *fiber.Ctx) error {
	var tokoId = ctx.Params("id_toko")

	namaToko := ctx.FormValue("nama_toko")
	fileName := ctx.Locals("fileName").(string)

	tokoUpdate := entity.Toko{
		NamaToko: namaToko,
		UrlFoto:  fileName,
	}

	_, errUpdate := tokoService.Update(tokoId, tokoUpdate)
	if errUpdate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errUpdate)
	}

	return helper.SuccessHelper(ctx, fiber.StatusAccepted, "Update toko succeed")
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

func TokoGetMy(ctx *fiber.Ctx) error {
	userInfo, isOk := helper.ClaimsUserInfo(ctx)
	if isOk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, isOk)
	}

	userId, _ := userInfo["id"].(float64)
	id := fmt.Sprintf("%.0f", userId)

	toko, err := tokoService.GetByUserId(id)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, toko)
}

func TokoGetAllPage(ctx *fiber.Ctx) error {
	name := ctx.Query("nama")
	page, perPage := getPagination(ctx)

	result, err := tokoService.GetAllPage(page, perPage, name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	return helper.SuccessHelper(ctx, fiber.StatusOK, result)
}

func getPagination(ctx *fiber.Ctx) (page, perPage uint64) {
	page = 1
	perPage = 10
	if queryPage := ctx.Query("page"); queryPage != "" {
		page, _ = strconv.ParseUint(queryPage, 10, 64)
	}
	if queryPerPage := ctx.Query("limit"); queryPerPage != "" {
		perPage, _ = strconv.ParseUint(queryPerPage, 10, 64)
	}
	return page, perPage
}

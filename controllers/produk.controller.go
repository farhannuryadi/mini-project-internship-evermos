package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/services"
	"mini-project-internship/utils"
)

var fotoProdukService services.FotoProdukService = *services.NewFotoProdukService()
var produkService services.ProdukService = *services.NewProdukService()

func ProdukCreate(ctx *fiber.Ctx) error {
	var produkReq request.ProdukReq
	if err := ctx.BodyParser(&produkReq); err != nil {
		log.Println("data form tidak masuk")
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

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

	produkCreate := mapper.ProdukReqToProduk(produkReq)
	produkCreate.TokoID = toko.ID

	produk, errCreateProduk := produkService.Create(produkCreate)
	if errCreateProduk != nil {
		log.Println("Error create produk")
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errCreateProduk)
	}

	fileNames := ctx.Locals("fileNames").([]string)
	for i, fileName := range fileNames {
		log.Println("photo index ke-", i, ", nama photonya :", fileName)
		fotoProdukCreate := entity.FotoProduk{
			Url:      fileName,
			ProdukID: produk.ID,
		}
		errCreateFotoProduk := fotoProdukService.Create(fotoProdukCreate)
		if errCreateFotoProduk != nil {
			log.Println("Error create foto produk")
			return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errCreateFotoProduk)
		}
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, 1)
}

func ProdukGetAllPage(ctx *fiber.Ctx) error {
	namaProduk := ctx.Query("nama_produk")
	categoryId := ctx.Query("category_id")
	tokoId := ctx.Query("toko_id")
	maxHarga := ctx.Query("max_harga")
	minHarga := ctx.Query("min_harga")
	page, limit := getPagination(ctx)

	result, err := produkService.GetAllByFilter(namaProduk,
		categoryId, tokoId, maxHarga, minHarga, uint(limit), uint(page))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	return helper.SuccessHelper(ctx, fiber.StatusOK, result)
}

func ProdukGetById(ctx *fiber.Ctx) error {
	produkId := ctx.Params("id")

	produk, err := produkService.GetById(produkId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	tokoId := strconv.Itoa(int(produk.TokoID))
	categoryId := strconv.Itoa(int(produk.CategoryID))

	toko, _ := tokoService.GetById(tokoId)
	category, _ := categoryServcie.GetById(categoryId)
	photos, _ := fotoProdukService.GetByProdukId(produkId)

	produkRes := mapper.ProdukToProdukRes(produk, toko, category, photos)

	return helper.SuccessHelper(ctx, fiber.StatusOK, produkRes)
}

func ProdukUpdate(ctx *fiber.Ctx) error {
	produkId := ctx.Params("id")
	var produkReq request.ProdukReq
	if err := ctx.BodyParser(&produkReq); err != nil {
		log.Println("data form tidak masuk")
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

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

	produk, errProduk := produkService.GetById(produkId)
	if errProduk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	if produk.TokoID != toko.ID {
		return helper.ErrorHelper(ctx, fiber.StatusForbidden, err)
	}

	photos, _ := fotoProdukService.GetByProdukId(produkId)

	fileNames := ctx.Locals("fileNames").([]string)
	for i, fileName := range fileNames {
		log.Println("photo index ke-", i, ", nama photonya :", fileName)

		fotoProdukSave := photos[i]
		fotoProdukSave.Url = fileName
		_, errCreateFotoProduk := fotoProdukService.Update(fotoProdukSave)
		if errCreateFotoProduk != nil {
			log.Println("Error create foto produk")
			return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errCreateFotoProduk)
		}
	}

	var del int = len(fileNames)

	for _, photo := range photos {
		if del > 0 {
			errDelFoto := utils.HandleRemoveFile(photo.Url)
			if errDelFoto != nil {
				log.Println("something wrong to delete", photo.Url)
			}
			del--
		}
	}

	produkSave := mapper.ProdukUpdateToProduk(produkReq, produk)

	_, errProdukUpdate := produkService.Update(produkSave)
	if errProdukUpdate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errProdukUpdate)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, "")
}

func ProdukDelete(ctx *fiber.Ctx) error {
	produkId := ctx.Params("id")

	produk, err := produkService.GetById(produkId)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, err)
	}

	photos, _ := fotoProdukService.GetByProdukId(produkId)
	for _, photo := range photos {
		_, errPhotoDBDel := fotoProdukService.Delete(photo)
		errDelFoto := utils.HandleRemoveFile(photo.Url)
		if errDelFoto != nil || errPhotoDBDel != nil {
			log.Println("something wrong to delete", photo.Url)
		}
	}

	_, errDelProduk := produkService.Delete(produk)
	if errDelProduk != nil {
		log.Println("something wrong to delete", produk.NamaProduk)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, "")
}

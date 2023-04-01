package controllers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/models/entity"
	"mini-project-internship/models/request"
	"mini-project-internship/services"
)

var trxService services.TrxService = *services.NewTrxService()
var detailTrxService services.DetailTrxService = *services.NewDetailTrxService()
var logProdukService services.LogProdukService = *services.NewLogProdukService()

func TrxCreate(ctx *fiber.Ctx) error {
	var trxReq request.TrxReq

	if err := ctx.BodyParser(&trxReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	userInfo, isOk := helper.ClaimsUserInfo(ctx)
	if isOk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, isOk)
	}

	userId, _ := userInfo["id"].(float64)
	id := fmt.Sprintf("%.0f", userId)

	var logProduks []entity.LogProduk
	var prodHargaTotal []int
	var trxHargaTotal int = 0

	for i, produk := range trxReq.DetailTrx {
		prod, _ := produkService.GetById(strconv.Itoa(produk.ProductID))
		logProduk, errLogCreate := logProdukService.Create(prod)
		if errLogCreate != nil {
			return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errLogCreate)
		}

		logProduks = append(logProduks, logProduk)
		num, _ := strconv.Atoi(logProduk.HargaKonsumen)
		prodHargaTotal = append(prodHargaTotal, (num * produk.Kuantitas))
		trxHargaTotal = trxHargaTotal + prodHargaTotal[i]
	}

	IdUser, _ := strconv.Atoi(id)

	kodeInvoice := generateKodeInvoice(trxReq.MethodBayar, id, trxReq.AlamatKirim)

	trxCreate := entity.Trx{
		HargaTotal:  uint(trxHargaTotal),
		KodeInvoice: &kodeInvoice,
		MethodBayar: &trxReq.MethodBayar,
		AlamatID:    uint(trxReq.AlamatKirim),
		UserID:      uint(IdUser),
	}
	trx, errTrxCreate := trxService.Create(trxCreate)
	if errTrxCreate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errTrxCreate)
	}

	var detailTrxs []entity.DetailTrx
	for i, logProduk := range logProduks {
		detailTrxCreate := entity.DetailTrx{
			LogProdukID: logProduk.ID,
			Kuantitas:   trxReq.DetailTrx[i].Kuantitas,
			HargaTotal:  prodHargaTotal[i],
			TrxID:       trx.ID,
			TokoID:      logProduk.TokoID,
		}

		detailTrx, errDetailTrxCreate := detailTrxService.Create(detailTrxCreate)
		if errDetailTrxCreate != nil {
			return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, errDetailTrxCreate)
		}
		detailTrxs = append(detailTrxs, detailTrx)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, detailTrxs)
}

func TrxGetAllPage(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	page, perPage := getPagination(ctx)

	userInfo, isOk := helper.ClaimsUserInfo(ctx)
	if isOk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, isOk)
	}

	userId, _ := userInfo["id"].(float64)
	id := fmt.Sprintf("%.0f", userId)

	result, err := trxService.GetAllPage(id, search, page, perPage)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusInternalServerError, err)
	}
	return helper.SuccessHelper(ctx, fiber.StatusOK, result)
}

func generateKodeInvoice(methodBayar, userId string, alamatKirim int) string {
	t := time.Now()
	formattedTime := t.Format("20060102")
	rand.Seed(t.UnixNano())
	randomNumber := rand.Intn(100)
	kodeInvoice := fmt.Sprintf("%s%d%s%s%d",
		methodBayar, randomNumber, formattedTime, userId, alamatKirim)
	return kodeInvoice
}

func TrxGetById(ctx *fiber.Ctx) error {
	trxId := ctx.Params("id")

	userInfo, isOk := helper.ClaimsUserInfo(ctx)
	if isOk != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, isOk)
	}

	userId, _ := userInfo["id"].(float64)
	id := fmt.Sprintf("%.0f", userId)

	trx, errGet := trxService.GetById(trxId)
	if errGet != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, errGet)
	}
	userIdGet := strconv.Itoa(int(trx.UserID))

	if id != userIdGet {
		return helper.ErrorHelper(ctx, fiber.StatusForbidden, fmt.Errorf("forbidden"))
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, trx)
}

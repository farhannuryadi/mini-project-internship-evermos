package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/helper/mapper"
	"mini-project-internship/models/request"
	"mini-project-internship/services"
	"mini-project-internship/utils"
)

var authService services.AuthService = services.AuthService{}
var provinsiService services.ProvinsiService = services.ProvinsiService{}
var kotaService services.KotaService = services.KotaService{}

func AuthLogin(ctx *fiber.Ctx) error {
	loginReq := new(request.LoginReq)
	if err := ctx.BodyParser(&loginReq); err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, err)
	}

	validate := validator.New()
	errValidate := validate.Struct(loginReq)
	if errValidate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errValidate)
	}

	user, err := userService.GetByNoTelp(loginReq.NoTelp)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, fmt.Errorf("No Telp atau kata sandi salah"))
	}

	provinsi, errProvinsi := provinsiService.GetById(*user.IdProvinsi)
	if errProvinsi != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, fmt.Errorf("provinsi user tidak ditemukan"))
	}

	kota, errkota := kotaService.GetById(*user.IdKota)
	if errkota != nil {
		return helper.ErrorHelper(ctx, fiber.StatusNotFound, fmt.Errorf("kota user tidak ditemukan"))
	}

	isValid := utils.CheckPasswordHas(loginReq.Password, user.KataSandi)

	if !isValid {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, fmt.Errorf("No Telp atau kata sandi salah"))
	}

	claims := utils.GenerateClaims(user)
	token, errGenerateToken := utils.GenerateTokenJwt(&claims)

	if errGenerateToken != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, fmt.Errorf("wrong credential"))
	}

	loginRes := mapper.UserToLoginRes(user, provinsi, kota, token)

	return helper.SuccessHelper(ctx, fiber.StatusOK, loginRes)
}

func AuthRegis(ctx *fiber.Ctx) error {
	regisReq := new(request.RegisterReq)
	if err := ctx.BodyParser(&regisReq); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(regisReq)
	if errValidate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errValidate)
	}

	_, errRegis := authService.Create(*regisReq)
	if errRegis != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errRegis)
	}

	user, errGet := userService.GetByNoTelp(regisReq.NoTelpon)
	if errGet != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errGet)
	}

	var userId string = strconv.Itoa(int(user.ID))
	log.Println("ini string")
	log.Println(userId)
	var toko = request.TokoReq{
		NamaToko: regisReq.Nama,
		UserID:   strconv.Itoa(int(user.ID)),
	}

	errToko := tokoService.Create(toko)
	if errToko != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errToko)
	}

	return helper.SuccessHelper(ctx, fiber.StatusOK, "Register Succeed")
}

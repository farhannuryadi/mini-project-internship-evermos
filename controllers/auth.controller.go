package controllers

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/helper"
	"mini-project-internship/models/request"
	"mini-project-internship/utils"
)

func AuthLogin(ctx *fiber.Ctx) error {
	loginReq := new(request.LoginReq)
	if err := ctx.BodyParser(&loginReq); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginReq)
	if errValidate != nil {
		return helper.ErrorHelper(ctx, fiber.StatusBadRequest, errValidate)
	}

	user, err := userService.GetByEmail(loginReq.Email)
	if err != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, fmt.Errorf("wrong credential"))
	}

	isValid := utils.CheckPasswordHas(loginReq.Password, user.KataSandi)
	
	if !isValid {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, fmt.Errorf("wrong credential"))
	}

	claims := utils.GenerateClaims(user)
	token, errGenerateToken := utils.GenerateTokenJwt(&claims)
	
	if errGenerateToken != nil {
		return helper.ErrorHelper(ctx, fiber.StatusUnauthorized, fmt.Errorf("wrong credential"))
	}

	log.Println(ctx.Method())
	return helper.SuccessHelper(ctx, fiber.StatusAccepted, fiber.Map{"token": token})
}
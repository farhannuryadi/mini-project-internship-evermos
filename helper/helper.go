package helper

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func ErrorHelper(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(
		response(false, fmt.Sprintf("Failed to %v", ctx.Method()), err, nil),
	)
}

func SuccessHelper(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(
		response(true, fmt.Sprintf("Succeed to %v", ctx.Method()), nil, data),
	)
}

func response(status bool, message string, err error, data interface{}) fiber.Map {
	var errRes interface{}
	if err != nil {
		errRes = err.Error()
	}
	return fiber.Map{
		"status":  status,
		"message": message,
		"error":   errRes,
		"data":    data,
	}
}

func ClaimsUserInfo(ctx *fiber.Ctx) (jwt.MapClaims, error) {
	userInfo, ok := ctx.Locals("userInfo").(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Error Claims User Info")
	}
	return userInfo, nil
}

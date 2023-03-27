package helper

import (
	"fmt"

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
		"status": status,
		"message": message,
		"error": errRes,
		"data": data,
	}
}
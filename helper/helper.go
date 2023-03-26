package helper

import "github.com/gofiber/fiber/v2"

func ErrorHelper(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status": "false",
		"code": status,
		"message": message,
	})
}

func SuccessHelper(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status": "true",
		"code": status,
		"message": "success",
		"data": data,
	})
}
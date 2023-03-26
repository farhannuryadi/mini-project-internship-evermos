package middleware

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("X-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	_, err := utils.VerifyToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
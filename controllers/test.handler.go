package controllers

import (
	"github.com/gofiber/fiber/v2"

)

func TestHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
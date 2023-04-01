package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// if claims["is_admin"] == false {
	// 	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"message": "forbidden access",
	// 	})
	// }

	ctx.Locals("userInfo", claims)

	return ctx.Next()
}

func AuthAdmin(ctx *fiber.Ctx) error {
	auth := Auth(ctx)
	userInfo, ok := ctx.Locals("userInfo").(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	isAdmin := userInfo["is_admin"].(bool)
	// log.Println("data user info ::", userInfo)
	// isAdmin := bool(userInfo["is_admin"])
	if !isAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}
	return auth
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}

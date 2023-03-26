package middleware

import "github.com/gofiber/fiber/v2/middleware/cors"

func CorsConfig() cors.Config{
	return cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}
}
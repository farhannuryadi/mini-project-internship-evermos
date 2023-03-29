package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"mini-project-internship/database"
	"mini-project-internship/middleware"
	"mini-project-internship/routes"
	// "mini-project-internship/utils"
)

func main() {
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")

	database.DatabaseInit()
	database.RunMigration()

	// utils.FetchData()

	app := fiber.New()
	app.Use(cors.New(middleware.CorsConfig()))
	routes.RoutesInit(app)
	app.Listen(port)
}

package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"mini-project-internship/database"
	"mini-project-internship/middleware"
	"mini-project-internship/routes"
)

func main() {
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")

	database.DatabaseInit()
	database.RunMigration()
	
	app := fiber.New()
	app.Use(cors.New(middleware.CorsConfig()))
	routes.RoutesInit(app)
	app.Listen(port)
}
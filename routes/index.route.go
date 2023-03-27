package routes

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/controllers"
	"mini-project-internship/middleware"
)

func RoutesInit(r *fiber.App) {
	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	user := v1.Group("/user")
	toko := v1.Group("/toko")
	category := v1.Group("/category")
	// produk := v1.Group("/produk")
	// trx := v1.Group("/trx")
	test := v1.Group("/test")

	test.Get("/", controllers.AmbilDataProvinsi)
	

	auth.Post("/login", controllers.AuthLogin)

	user.Post("/", controllers.UserCreate)
	user.Get("/:id", controllers.UserGetById)
	user.Put("/:id", controllers.UserUpdate)
	user.Get("/", middleware.AuthAdmin, controllers.UserGetAll)
	user.Delete("/:id", controllers.UserDelete) 

	toko.Post("/", controllers.TokoCreate)
	toko.Get("/:id", controllers.TokoGetById)
	toko.Put("/:id", controllers.TokoUpdate)
	toko.Get("/", controllers.TokoGetAll)
	toko.Delete("/:id", controllers.TokoDelete)

	category.Post("/", middleware.AuthAdmin ,controllers.CategoryCreate)
	category.Get("/:id", middleware.AuthAdmin, controllers.CategoryGetById)
	category.Put("/:id", middleware.AuthAdmin, controllers.CategoryUpdate)
	category.Get("/", middleware.AuthAdmin, controllers.CategoryGetAll)
	category.Delete("/:id", middleware.AuthAdmin, controllers.CategoryDelete)
}
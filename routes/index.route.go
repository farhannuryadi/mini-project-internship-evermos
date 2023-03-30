package routes

import (
	"github.com/gofiber/fiber/v2"

	"mini-project-internship/controllers"
	"mini-project-internship/middleware"
	"mini-project-internship/utils"

)

func RoutesInit(r *fiber.App) {
	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	user := v1.Group("/user")
	toko := v1.Group("/toko")
	alamat := v1.Group("/user/alamat")
	category := v1.Group("/category")
	produk := v1.Group("/product")
	// trx := v1.Group("/trx")
	test := v1.Group("/test")

	test.Get("/", controllers.AmbilDataProvinsi)
	test.Post("/", controllers.TestProdukFoto)
	test.Post("/multiple", utils.HandleMultipleFile, controllers.TestUploadMultiple)

	produk.Post("/", utils.HandleMultipleFile, middleware.Auth, controllers.ProdukCreate)
	produk.Get("/", middleware.Auth, controllers.ProdukGetAllPage)
	produk.Get("/:id", middleware.Auth, controllers.ProdukGetById)
	produk.Put("/:id", utils.HandleMultipleFile, middleware.Auth, controllers.ProdukGetById)
	produk.Delete("/:id", middleware.Auth, controllers.ProdukGetById)

	auth.Post("/login", controllers.AuthLogin)
	auth.Post("/register", controllers.AuthRegis)

	user.Post("/", controllers.UserCreate)
	user.Put("/", middleware.Auth, controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDelete)
	// user.Get("/list", middleware.AuthAdmin, controllers.UserGetAll)
	user.Get("/:id_user", controllers.UserGetById)
	user.Get("/", middleware.Auth, controllers.UserMyProfile)

	alamat.Post("/", middleware.Auth, controllers.AlamatCreate)
	alamat.Get("/:id", middleware.Auth, controllers.AlamatGetById)
	alamat.Get("/", middleware.Auth, controllers.AlamatGetByUserId)
	alamat.Put("/:id", middleware.Auth, controllers.AlamatUpdate)
	alamat.Delete("/:id", middleware.Auth, controllers.AlamatDelete)

	toko.Post("/", controllers.TokoCreate)
	toko.Put("/:id", middleware.Auth, controllers.TokoUpdate)
	toko.Get("/", middleware.Auth, controllers.TokoGetAllPage)
	toko.Delete("/:id", controllers.TokoDelete)
	toko.Get("/my", middleware.Auth, controllers.TokoGetMy)
	toko.Get("/:id_toko", middleware.Auth, controllers.TokoGetById)

	category.Post("/", middleware.AuthAdmin, controllers.CategoryCreate)
	category.Get("/:id", middleware.AuthAdmin, controllers.CategoryGetById)
	category.Put("/:id", middleware.AuthAdmin, controllers.CategoryUpdate)
	category.Get("/", middleware.AuthAdmin, controllers.CategoryGetAll)
	category.Delete("/:id", middleware.AuthAdmin, controllers.CategoryDelete)
}

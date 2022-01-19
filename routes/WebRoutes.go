package routes

import (
	"golangtest/controllers"
	"golangtest/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(App *fiber.App) {

	api := App.Group("/api", middleware.MiddlewareApi)
	api.Get("/", controllers.Welcome)
	api.Get("/allproduk", controllers.GetAllProduk)
	api.Post("/createproduk", controllers.AddProduk)
	api.Patch("/updateproduk/:id", controllers.UpdateProduk)
	api.Delete("/deleteproduk/:id", controllers.DeleteProduk)
	api.Get("/produk/:id", controllers.GetById)
	api.Post("/uploadimage", controllers.UploadPhoto)

	App.Post("/register", controllers.RegisterUser)
	App.Post("/login", controllers.Login)
}

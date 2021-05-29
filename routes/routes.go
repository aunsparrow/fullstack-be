package routes

import (
	controllers "be_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	apiRoutes := app.Group("/api") // /api CreateProduct

	shopApi := apiRoutes.Group("shop")
	shopApi.Post("/create", controllers.CreateShop)
	shopApi.Post("/update", controllers.UpdateShop)
	shopApi.Post("/delete", controllers.DeleteShop)
	shopApi.Get("/get/:shop_id", controllers.GetShopById)
	shopApi.Get("/getall", controllers.GetAllShop)

	productApi := apiRoutes.Group("product")
	productApi.Post("/create", controllers.CreateProduct)
	productApi.Post("/update", controllers.UpdateProduct)
	productApi.Post("/delete", controllers.DeleteProduct)
	productApi.Get("/get/:product_id", controllers.GetProductById)
	productApi.Get("/getall", controllers.GetAllProduct)
	productApi.Get("/get/shop/:shop_id", controllers.GetProductByShopId)

	categoryApi := apiRoutes.Group("category")
	categoryApi.Get("getall", controllers.GetAllCategory)

}

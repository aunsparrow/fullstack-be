package main

import (
	"be_api/db"
	"be_api/dbmodels"
	"be_api/env"
	repo "be_api/repositories"
	"be_api/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	switch os.Getenv("ENV") {
	default:
		env.SetDefaultEnv()
		fmt.Println(os.Getenv("ENV"))
	}

	db.Connect()
	dbs := db.OpenDB()
	dbs.AutoMigrate(&dbmodels.Shops{}, &dbmodels.Categories{}, &dbmodels.Products{})

	category := []dbmodels.Categories{
		dbmodels.Categories{
			CategoryName:   "เครื่องดื่มอัดลมและน้ำหวาน",
			CategoryDetail: "test",
		},
		dbmodels.Categories{
			CategoryName:   "ช่วยดับกระหาย เพิ่มความสดชื่น",
			CategoryDetail: "test",
		},
	}
	dbs.Create(&category)

	app := fiber.New(fiber.Config{
		BodyLimit: 25 * 1024 * 1024,
		// Prefork:       true,
		// CaseSensitive: true,
		// StrictRouting: true,
		// ServerHeader:  "Fiber",
	})

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format:     "${blue}${time} ${yellow}${status} - ${red}${latency} ${cyan}${method} ${path} ${green} ${ip} ${ua} ${reset}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     os.Stdout,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		var shopRepo repo.ShopRepository
		shopRepo = &repo.Repository{}
		var model dbmodels.Shops
		model.Tel = "000000"
		fmt.Println("safsdfs")
		ss, _ := shopRepo.SaveShop(model)
		_ = ss

		return c.SendString("AED Chatbot API v1 " + os.Getenv("ENV"))
	})
	routes.SetUpRoutes(app)
	app.Listen(":7000")
}

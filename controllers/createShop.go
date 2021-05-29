package controllers

import (
	"fmt"
	//service "be_api/services"
	reponseModels "be_api/reponseModels"
	requestModels "be_api/requestModels"
	service "be_api/services"

	"github.com/gofiber/fiber/v2"
)

func CreateShop(c *fiber.Ctx) error {
	reqBody := new(requestModels.CreateShop)
	if err := c.BodyParser(reqBody); err != nil {
		fmt.Println("bad req", err)
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var shopService service.ShopService
	shopService = &service.Services{}
	createShop, err := shopService.CreateShop(reqBody)
	if err != nil {
		fmt.Println("shopService", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, createShop}
	return c.Status(fiber.StatusOK).JSON(res)
}

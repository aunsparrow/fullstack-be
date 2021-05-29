package controllers

import (
	"fmt"

	//service "be_api/services"
	"be_api/helpers"
	reponseModels "be_api/reponseModels"
	"be_api/requestModels"
	service "be_api/services"

	"github.com/gofiber/fiber/v2"
)

func GetShopById(c *fiber.Ctx) error {
	shopId := c.Params("shop_id")
	if helpers.StringIsNullOrEmpty(shopId) {
		fmt.Println("bad req")
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var shopService service.ShopService
	shopService = &service.Services{}
	getShop, err := shopService.GetShopById(shopId)
	if err != nil {
		fmt.Println("shopService", err)
		status := reponseModels.StatusModel{fiber.StatusOK, "Shop not found."}
		res := reponseModels.ResponseModel{status, getShop}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	if helpers.StringIsNullOrEmpty(getShop.ShopId) {
		status := reponseModels.StatusModel{fiber.StatusOK, "Shop not found."}
		res := reponseModels.ResponseModel{status, getShop}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, getShop}
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetAllShop(c *fiber.Ctx) error {
	reqBody := new(requestModels.Pagination)
	if err := c.BodyParser(reqBody); err != nil {
	}

	var shopService service.ShopService
	shopService = &service.Services{}
	getShop, err := shopService.GetAllShop(*reqBody)
	if err != nil {
		fmt.Println("shopService", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, getShop}
	return c.Status(fiber.StatusOK).JSON(res)
}

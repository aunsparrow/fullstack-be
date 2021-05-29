package controllers

import (
	"fmt"
	//service "be_api/services"
	"be_api/helpers"
	reponseModels "be_api/reponseModels"
	requestModels "be_api/requestModels"
	service "be_api/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteShop(c *fiber.Ctx) error {
	reqBody := new(requestModels.DeleteShop)
	if err := c.BodyParser(reqBody); err != nil || helpers.StringIsNullOrEmpty(reqBody.ShopId) {
		fmt.Println("bad req", err)
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var shopService service.ShopService
	shopService = &service.Services{}
	deleteShop, err := shopService.DeleteShop(reqBody.ShopId)
	if err != nil || !deleteShop {
		fmt.Println("deleteShop", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, nil}
	return c.Status(fiber.StatusOK).JSON(res)
}

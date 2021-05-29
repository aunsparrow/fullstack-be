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

func DeleteProduct(c *fiber.Ctx) error {
	reqBody := new(requestModels.DeleteProduct)
	if err := c.BodyParser(reqBody); err != nil || helpers.StringIsNullOrEmpty(reqBody.ProductId) {
		fmt.Println("bad req", err)
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var productService service.ProductService
	productService = &service.Services{}
	deleteProduct, err := productService.DeleteProduct(reqBody.ProductId)
	if err != nil || !deleteProduct {
		fmt.Println("deleteProduct", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, nil}
	return c.Status(fiber.StatusOK).JSON(res)
}

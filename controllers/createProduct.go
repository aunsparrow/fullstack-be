package controllers

import (
	"fmt"
	//service "be_api/services"
	reponseModels "be_api/reponseModels"
	requestModels "be_api/requestModels"
	service "be_api/services"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	reqBody := new(requestModels.CreateProduct)
	if err := c.BodyParser(reqBody); err != nil {
		fmt.Println("bad req", err)
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var productService service.ProductService
	productService = &service.Services{}

	createProduct, err := productService.CreateProduct(reqBody)
	if err != nil {
		fmt.Println("createProduct", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	if err == nil && createProduct == nil {
		fmt.Println("bad request", err)
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, createProduct}
	return c.Status(fiber.StatusOK).JSON(res)
}

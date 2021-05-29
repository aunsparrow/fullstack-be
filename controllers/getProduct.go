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

func GetProductById(c *fiber.Ctx) error {
	productId := c.Params("product_id")
	if helpers.StringIsNullOrEmpty(productId) {
		fmt.Println("bad req")
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var productService service.ProductService
	productService = &service.Services{}
	getProduct, err := productService.GetProductById(productId)
	if err != nil || helpers.StringIsNullOrEmpty(getProduct.ProductId) {
		fmt.Println("bad req")
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, getProduct}
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetAllProduct(c *fiber.Ctx) error {
	reqBody := new(requestModels.Pagination)
	if err := c.BodyParser(reqBody); err != nil {
	}

	var productService service.ProductService
	productService = &service.Services{}
	getAll, err := productService.GetAllProduct(*reqBody)
	if err != nil {
		fmt.Println("getAll", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, getAll}
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetProductByShopId(c *fiber.Ctx) error {
	shopId := c.Params("shop_id")
	if helpers.StringIsNullOrEmpty(shopId) {
		fmt.Println("bad req")
		status := reponseModels.StatusModel{fiber.StatusBadRequest, "Bad request."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	var productService service.ProductService
	productService = &service.Services{}
	getAll, err := productService.GetProductByShopId(shopId)
	if err != nil {
		fmt.Println("bad req")
		status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
		res := reponseModels.ResponseModel{status, getAll}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, getAll}
	return c.Status(fiber.StatusOK).JSON(res)
}

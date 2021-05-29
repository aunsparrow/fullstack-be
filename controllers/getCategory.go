package controllers

import (
	"fmt"

	//service "be_api/services"

	reponseModels "be_api/reponseModels"
	repo "be_api/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(c *fiber.Ctx) error {
	var categoryService repo.CategoryRepository
	categoryService = &repo.Repository{}
	getcategory, err := categoryService.GetAllCategory()
	if err != nil {
		fmt.Println("getcategory", err)
		status := reponseModels.StatusModel{fiber.StatusInternalServerError, "Internal Server Error."}
		res := reponseModels.ResponseModel{status, nil}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	status := reponseModels.StatusModel{fiber.StatusOK, "Success."}
	res := reponseModels.ResponseModel{status, getcategory}
	return c.Status(fiber.StatusOK).JSON(res)
}

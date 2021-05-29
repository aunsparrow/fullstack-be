package services

import (
	dbmodels "be_api/dbmodels"
	"be_api/helpers"
	responseModels "be_api/reponseModels"
	repo "be_api/repositories"
	requestModels "be_api/requestModels"
	"fmt"
	"time"
)

type ProductService interface {
	CreateProduct(product *requestModels.CreateProduct) (*dbmodels.Products, error)
	UpdateProduct(product *requestModels.UpdateProduct) (*responseModels.ProductModel, error)
	DeleteProduct(productId string) (bool, error)
	GetProductById(productId string) (*responseModels.ProductModel, error)
	GetAllProduct(page requestModels.Pagination) ([]responseModels.ProductModel, error)
	GetProductByShopId(shopId string) ([]responseModels.ProductModel, error)
}

func (service *Services) CreateProduct(product *requestModels.CreateProduct) (*dbmodels.Products, error) {
	model := dbmodels.Products{
		ProductName:   product.ProductName,
		ProductDetail: product.ProductDetail,
		ProductPrice:  product.ProductPrice,
		ProductUnit:   product.ProductUnit,
		CategoryId:    product.CategoryId,
		ShopId:        product.ShopId,
	}

	var categoryRepo repo.CategoryRepository
	categoryRepo = &repo.Repository{}
	category, err := categoryRepo.GetCategoryById(product.CategoryId)
	if helpers.StringIsNullOrEmpty(category.CategoryId) || err != nil {
		fmt.Println("asdsadasdsa")
		return nil, nil
	}

	var shopRepo repo.ShopRepository
	shopRepo = &repo.Repository{}
	shop, err := shopRepo.GetShopById(product.ShopId)
	if helpers.StringIsNullOrEmpty(shop.ShopId) || err != nil {
		fmt.Println("asdsadasdsa")
		return nil, nil
	}

	var productRepo repo.ProductRepository
	productRepo = &repo.Repository{}

	createProd, err := productRepo.SaveProduct(model)
	if err != nil || helpers.StringIsNullOrEmpty(createProd.ProductId) {
		fmt.Println("createProd", err)
		return nil, err
	}
	return createProd, nil
}

func (service *Services) UpdateProduct(product *requestModels.UpdateProduct) (*responseModels.ProductModel, error) {
	var response responseModels.ProductModel
	at := time.Now()
	model := dbmodels.Products{
		ProductId:     product.ProductId,
		ProductName:   product.ProductName,
		ProductDetail: product.ProductDetail,
		ProductPrice:  product.ProductPrice,
		ProductUnit:   product.ProductUnit,
		UpdateAt:      &at,
	}

	var categoryRepo repo.CategoryRepository
	categoryRepo = &repo.Repository{}
	catTmp, _ := categoryRepo.GetCategoryById(product.CategoryId)
	if !helpers.StringIsNullOrEmpty(catTmp.CategoryId) {
		model.CategoryId = catTmp.CategoryId
	}

	var productRepo repo.ProductRepository
	productRepo = &repo.Repository{}

	getProduct, _ := productRepo.GetProductById(product.ProductId)
	if helpers.StringIsNullOrEmpty(getProduct.ProductId) {
		return nil, nil
	}

	updateProd, err := productRepo.UpdateProduct(model)
	if err != nil || helpers.StringIsNullOrEmpty(updateProd.ProductId) {
		fmt.Println("createProd", err)
		return nil, err
	}

	if helpers.StringIsNullOrEmpty(catTmp.CategoryId) {
		catTmp, _ = categoryRepo.GetCategoryById(updateProd.CategoryId)
	}

	response.CreateAt = updateProd.CreateAt
	response.ProductDetail = updateProd.ProductDetail
	response.ProductId = updateProd.ProductId
	response.ProductName = updateProd.ProductName
	response.ProductPrice = updateProd.ProductPrice
	response.ProductUnit = updateProd.ProductUnit
	response.ShopId = updateProd.ShopId
	response.UpdateAt = updateProd.UpdateAt
	response.Category = *catTmp
	return &response, nil
}

func (service *Services) DeleteProduct(productId string) (bool, error) {
	var productRepo repo.ProductRepository
	productRepo = &repo.Repository{}
	deleteProduct, err := productRepo.DeleteProduct(productId)
	if err != nil || !deleteProduct {
		fmt.Println("deleteProduct", err)
		return false, err
	}
	return true, nil
}

func (service *Services) GetProductById(productId string) (*responseModels.ProductModel, error) {
	var response responseModels.ProductModel
	var productRepo repo.ProductRepository
	productRepo = &repo.Repository{}
	getById, err := productRepo.GetProductById(productId)
	if err != nil {
		fmt.Println("getById", err)
		return nil, err
	}

	var categoryRepo repo.CategoryRepository
	categoryRepo = &repo.Repository{}
	getCategory, err := categoryRepo.GetCategoryById(getById.CategoryId)
	if err != nil {
		fmt.Println("getCategory", err)
		return nil, err
	}

	response.CreateAt = getById.CreateAt
	response.ProductDetail = getById.ProductDetail
	response.ProductId = getById.ProductId
	response.ProductName = getById.ProductName
	response.ProductPrice = getById.ProductPrice
	response.ProductUnit = getById.ProductUnit
	response.ShopId = getById.ShopId
	response.UpdateAt = getById.UpdateAt
	response.Category = *getCategory

	return &response, nil
}

func (service *Services) GetAllProduct(page requestModels.Pagination) ([]responseModels.ProductModel, error) {
	var response []responseModels.ProductModel = make([]responseModels.ProductModel, 0)
	var productRepo repo.ProductRepository
	productRepo = &repo.Repository{}
	getAll, err := productRepo.GetAllProduct(page)
	if err != nil {
		fmt.Println("getAll", err)
		return nil, err
	}

	var categoryRepo repo.CategoryRepository
	categoryRepo = &repo.Repository{}

	for _, prod := range getAll {
		tempCategory, _ := categoryRepo.GetCategoryById(prod.CategoryId)
		var prodSet responseModels.ProductModel
		prodSet.CreateAt = prod.CreateAt
		prodSet.ProductDetail = prod.ProductDetail
		prodSet.ProductId = prod.ProductId
		prodSet.ProductName = prod.ProductName
		prodSet.ProductPrice = prod.ProductPrice
		prodSet.ProductUnit = prod.ProductUnit
		prodSet.ShopId = prod.ShopId
		prodSet.UpdateAt = prod.UpdateAt
		prodSet.Category = *tempCategory
		response = append(response, prodSet)
	}

	return response, nil
}

func (service *Services) GetProductByShopId(shopId string) ([]responseModels.ProductModel, error) {
	var response []responseModels.ProductModel = make([]responseModels.ProductModel, 0)
	var productRepo repo.ProductRepository
	productRepo = &repo.Repository{}
	getAll, err := productRepo.GetProductByShopId(shopId)
	if err != nil {
		fmt.Println("getAll", err)
		return nil, err
	}

	var categoryRepo repo.CategoryRepository
	categoryRepo = &repo.Repository{}

	for _, prod := range getAll {
		tempCategory, _ := categoryRepo.GetCategoryById(prod.CategoryId)
		var prodSet responseModels.ProductModel
		prodSet.CreateAt = prod.CreateAt
		prodSet.ProductDetail = prod.ProductDetail
		prodSet.ProductId = prod.ProductId
		prodSet.ProductName = prod.ProductName
		prodSet.ProductPrice = prod.ProductPrice
		prodSet.ProductUnit = prod.ProductUnit
		prodSet.ShopId = prod.ShopId
		prodSet.UpdateAt = prod.UpdateAt
		prodSet.Category = *tempCategory
		response = append(response, prodSet)
	}

	return response, nil
}

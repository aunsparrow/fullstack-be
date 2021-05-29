package services

import (
	dbmodels "be_api/dbmodels"
	repo "be_api/repositories"
	requestModels "be_api/requestModels"
	"fmt"
	"time"
)

type ShopService interface {
	CreateShop(shop *requestModels.CreateShop) (*dbmodels.Shops, error)
	UpdateShop(shop *requestModels.UpdateShop) (*dbmodels.Shops, error)
	DeleteShop(shopId string) (bool, error)
	GetShopById(shopId string) (*dbmodels.Shops, error)
	GetAllShop(page requestModels.Pagination) ([]dbmodels.Shops, error)
}

func (service *Services) CreateShop(shop *requestModels.CreateShop) (*dbmodels.Shops, error) {
	model := dbmodels.Shops{
		ShopName: shop.ShopName,
		Tel:      shop.Tel,
		Address:  shop.Address}

	var shopRepo repo.ShopRepository
	shopRepo = &repo.Repository{}
	createShop, err := shopRepo.SaveShop(model)
	if err != nil {
		fmt.Println("createShop", err)
		return nil, err
	}
	return createShop, nil
}

func (service *Services) UpdateShop(shop *requestModels.UpdateShop) (*dbmodels.Shops, error) {
	at := time.Now()
	model := dbmodels.Shops{
		ShopId:   shop.ShopId,
		ShopName: shop.ShopName,
		Tel:      shop.Tel,
		Address:  shop.Address,
		UpdateAt: &at}

	var shopRepo repo.ShopRepository
	shopRepo = &repo.Repository{}
	updateShop, err := shopRepo.UpdateShop(model)
	if err != nil {
		fmt.Println("UpdateShop", err)
		return nil, err
	}
	return updateShop, nil
}

func (service *Services) DeleteShop(shopId string) (bool, error) {
	var shopRepo repo.ShopRepository
	shopRepo = &repo.Repository{}
	deleteShop, err := shopRepo.DeleteShop(shopId)
	if err != nil || !deleteShop {
		fmt.Println("DeleteShop", err)
		return false, err
	}
	return true, nil
}

func (service *Services) GetShopById(shopId string) (*dbmodels.Shops, error) {
	var shopRepo repo.ShopRepository
	shopRepo = &repo.Repository{}
	getById, err := shopRepo.GetShopById(shopId)
	if err != nil {
		fmt.Println("DeleteShop", err)
		return nil, err
	}
	return getById, nil
}

func (service *Services) GetAllShop(page requestModels.Pagination) ([]dbmodels.Shops, error) {
	var shopRepo repo.ShopRepository
	shopRepo = &repo.Repository{}
	getAll, err := shopRepo.GetAllShop(page)
	if err != nil {
		fmt.Println("DeleteShop", err)
		return nil, err
	}
	return getAll, nil
}

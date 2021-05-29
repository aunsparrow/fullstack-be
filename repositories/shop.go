package repositories

import (
	"be_api/db"
	"be_api/dbmodels"
	requestModels "be_api/requestModels"
	"fmt"
)

type ShopRepository interface {
	SaveShop(shop dbmodels.Shops) (*dbmodels.Shops, error)
	SaveManyShop(shops []dbmodels.Shops) ([]dbmodels.Shops, error)
	UpdateShop(shop dbmodels.Shops) (*dbmodels.Shops, error)
	DeleteShop(shopid string) (bool, error)
	GetShopById(shopid string) (*dbmodels.Shops, error)
	GetAllShop(page requestModels.Pagination) ([]dbmodels.Shops, error)
}

func (repo *Repository) SaveShop(shop dbmodels.Shops) (*dbmodels.Shops, error) {
	dB := db.Db
	if err := dB.Create(&shop).Error; err != nil {
		fmt.Println("Create", err)
		return nil, err
	}
	return &shop, nil
}

func (repo *Repository) SaveManyShop(shop []dbmodels.Shops) ([]dbmodels.Shops, error) {
	dB := db.Db
	if err := dB.Create(&shop).Error; err != nil {
		fmt.Println("Creates", err)
		return nil, err
	}
	return shop, nil
}

func (repo *Repository) UpdateShop(shop dbmodels.Shops) (*dbmodels.Shops, error) {
	dB := db.Db
	if err := dB.Model(&dbmodels.Shops{}).Where("shop_id = ?", shop.ShopId).Updates(&shop).Error; err != nil {
		fmt.Println("Update", err)
		return nil, err
	}

	if err := dB.Model(&dbmodels.Shops{}).Where("shop_id = ?", shop.ShopId).First(&shop).Error; err != nil {
		fmt.Println("Update", err)
		return nil, err
	}

	return &shop, nil
}

func (repo *Repository) DeleteShop(shopId string) (bool, error) {
	dB := db.Db
	if err := dB.Where("shop_id = ?", shopId).Delete(&dbmodels.Shops{}).Error; err != nil {
		fmt.Println("Delete", err)
		return false, err
	}
	return true, nil
}

func (repo *Repository) GetShopById(shopId string) (*dbmodels.Shops, error) {
	var model dbmodels.Shops
	dB := db.Db
	dB.Model(&dbmodels.Shops{}).Where("shop_id = ?", shopId).First(&model)
	return &model, nil
}

func (repo *Repository) GetAllShop(page requestModels.Pagination) ([]dbmodels.Shops, error) {
	offSet := 0
	if offSet > 0 {
		offSet = page.PageSize * (page.PageNumber - 1)
	}

	var model []dbmodels.Shops = make([]dbmodels.Shops, 0)

	dB := db.Db
	dB.Model(&dbmodels.Shops{}).Where("shop_name LIKE ?", "%"+page.FilterText+"%").Limit(page.PageSize).Offset(offSet).Find(&model)
	return model, nil
}

package repositories

import (
	"be_api/db"
	"be_api/dbmodels"
	requestModels "be_api/requestModels"
	"fmt"
)

type ProductRepository interface {
	SaveProduct(product dbmodels.Products) (*dbmodels.Products, error)
	UpdateProduct(product dbmodels.Products) (*dbmodels.Products, error)
	DeleteProduct(product string) (bool, error)
	GetProductById(product string) (*dbmodels.Products, error)
	GetAllProduct(page requestModels.Pagination) ([]dbmodels.Products, error)
	GetProductByShopId(shopid string) ([]dbmodels.Products, error)
}

func (repo *Repository) SaveProduct(product dbmodels.Products) (*dbmodels.Products, error) {
	dB := db.Db
	if err := dB.Create(&product).Error; err != nil {
		fmt.Println("Create", err)
		return nil, err
	}
	return &product, nil
}

func (repo *Repository) UpdateProduct(product dbmodels.Products) (*dbmodels.Products, error) {
	dB := db.Db
	if err := dB.Model(&dbmodels.Products{}).Where("product_id = ?", product.ProductId).Updates(&product).Error; err != nil {
		fmt.Println("Update", err)
		return nil, err
	}

	if err := dB.Model(&dbmodels.Products{}).Where("product_id = ?", product.ProductId).First(&product).Error; err != nil {
		fmt.Println("Update", err)
		return nil, err
	}

	return &product, nil
}

func (repo *Repository) DeleteProduct(productId string) (bool, error) {
	dB := db.Db
	if err := dB.Where("product_id = ?", productId).Delete(&dbmodels.Products{}).Error; err != nil {
		fmt.Println("Delete", err)
		return false, err
	}
	return true, nil
}

func (repo *Repository) GetProductById(productId string) (*dbmodels.Products, error) {
	var model dbmodels.Products
	dB := db.Db
	dB.Model(&dbmodels.Products{}).Where("product_id = ?", productId).First(&model)
	return &model, nil
}

func (repo *Repository) GetAllProduct(page requestModels.Pagination) ([]dbmodels.Products, error) {
	offSet := 0
	if offSet > 0 {
		offSet = page.PageSize * (page.PageNumber - 1)
	}

	var model []dbmodels.Products = make([]dbmodels.Products, 0)

	dB := db.Db
	dB.Model(&dbmodels.Products{}).Where("product_name LIKE ?", "%"+page.FilterText+"%").Limit(page.PageSize).Offset(offSet).Find(&model)
	return model, nil
}

func (repo *Repository) GetProductByShopId(shopId string) ([]dbmodels.Products, error) {
	var model []dbmodels.Products = make([]dbmodels.Products, 0)

	dB := db.Db
	dB.Model(&dbmodels.Products{}).Where("shop_id = ?", shopId).Find(&model)
	fmt.Println(shopId)
	fmt.Println(model)
	return model, nil
}

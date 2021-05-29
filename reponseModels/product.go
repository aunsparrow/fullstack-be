package reponseModels

import (
	"be_api/dbmodels"
	"time"
)

type ProductModel struct {
	ProductId     string              `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"product_id"`
	ProductName   string              `gorm:"type:text;" json:"product_name"`
	ProductDetail string              `gorm:"type:text;" json:"product_detail"`
	ProductUnit   string              `gorm:"type:text;" json:"product_unit"`
	ProductPrice  float32             `gorm:"type:float4;" json:"product_price"`
	CreateAt      *time.Time          `gorm:"type:timestamp;default:now()" json:"create_at"`
	UpdateAt      *time.Time          `gorm:"type:timestamp" json:"update_at"`
	ShopId        string              `gorm:"type:uuid" json:"shop_id"`
	Category      dbmodels.Categories `gorm:"type:uuid" json:"category"`
}

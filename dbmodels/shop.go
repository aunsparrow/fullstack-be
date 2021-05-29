package dbmodels

import "time"

type Shops struct {
	ShopId   string     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"shop_id"`
	ShopName string     `gorm:"type:text;" json:"shop_name"`
	Tel      string     `gorm:"type:text;" json:"tel"`
	Address  string     `gorm:"type:text;" json:"address"`
	CreateAt *time.Time `gorm:"type:timestamp;default:now()" json:"create_at"`
	UpdateAt *time.Time `gorm:"type:timestamp" json:"update_at"`
	Products []Products `gorm:"foreignKey:shop_id;references:shop_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"products"`
}

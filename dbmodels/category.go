package dbmodels

import "time"

type Categories struct {
	CategoryId     string     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"category_id"`
	CategoryName   string     `gorm:"type:text;" json:"category_name"`
	CategoryDetail string     `gorm:"type:text;" json:"category_detail"`
	CreateAt       *time.Time `gorm:"type:timestamp;default:now()" json:"create_at"`
	UpdateAt       *time.Time `gorm:"type:timestamp" json:"update_at"`
	Products       []Products `gorm:"foreignKey:category_id;references:category_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"products"`
}

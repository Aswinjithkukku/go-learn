package models

import "time"

type Order struct {
	ID             uint    `json:"id" gorm:"primaryKey"`
	ProductReferer int     `json:"productId"`
	Product        Product `gorm:"foreignKey:ProductReferer"`
	UserReferer    int     `json:"userId"`
	User           User    `gorm:"foreignKey:UserReferer"`
	CreatedAt      time.Time
}

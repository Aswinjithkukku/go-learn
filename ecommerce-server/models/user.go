package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique:true; not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time
}

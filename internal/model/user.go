package model

import "gorm.io/gorm"

// swagger:model
type User struct {
	gorm.Model
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// example: johndoe
	Username string `gorm:"size:255;uniqueIndex;not null" json:"username"`
	// example: foo@bar.com
	Email string `json:"email"`
	Bio   string `gorm:"size:100" json:"bio"`
}

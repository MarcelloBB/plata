package model

// swagger:model
type User struct {
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// example: johndoe
	Username string `gorm:"size:255;uniqueIndex;not null" json:"username"`
	// example: foo@bar.com
	Email string `json:"email"`
	// example: Just a regular user.
	Bio string `gorm:"size:100" json:"bio"`
}

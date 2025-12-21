package model

type Budget struct {
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// example: 1
	UserID int  `gorm:"not null;index" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
	// example: 1500.00
	Amount float64 `json:"amount"`
}

package model

type Category struct {
	// example: 1
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// example: Groceries
	Name string `gorm:"size:50;not null" json:"name"`
	// example: 1
	UserID int  `gorm:"not null;index" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`

	Budget Budget `gorm:"foreignKey:CategoryBudget" json:"budget"`
}

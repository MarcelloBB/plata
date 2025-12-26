package repository

import (
	"fmt"

	"github.com/MarcelloBB/plata/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{
		connection: db,
	}
}

func (p *CategoryRepository) GetCategories() ([]model.Category, error) {
	categories := []model.Category{}
	err := p.connection.Find(&categories).Error
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Category{}, err
	}

	return categories, nil
}

func (p *CategoryRepository) CreateCategory(category model.Category) (model.Category, error) {
	err := p.connection.Create(&category).Error
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

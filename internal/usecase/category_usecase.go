package usecase

import (
	"github.com/MarcelloBB/plata/internal/dto"
	"github.com/MarcelloBB/plata/internal/model"
	"github.com/MarcelloBB/plata/internal/repository"
)

type CategoryUseCase struct {
	repository repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) CategoryUseCase {
	return CategoryUseCase{
		repository: repo,
	}
}

func (p *CategoryUseCase) GetCategories() ([]model.Category, error) {
	return p.repository.GetCategories()
}

func (p *CategoryUseCase) CreateCategory(category dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	newCategory := model.Category{
		Name:   category.Name,
		UserID: category.UserID,
	}

	createdCategory, err := p.repository.CreateCategory(newCategory)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	response := dto.CategoryResponse{
		ID:     createdCategory.ID,
		Name:   createdCategory.Name,
		UserID: createdCategory.UserID,
	}

	return response, nil
}

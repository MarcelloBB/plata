package controller

import (
	"net/http"

	"github.com/MarcelloBB/plata/internal/dto"
	"github.com/MarcelloBB/plata/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(usecase usecase.CategoryUseCase) CategoryController {
	return CategoryController{
		categoryUseCase: usecase,
	}
}

// GetCategories godoc
// @Summary      List categories
// @Description  Returns a list of categories
// @Tags         category
// @Produce      json
// @Success      200  {array}   model.Category
// @Failure      500  {object}  map[string]string
// @Router       /category [get]
func (uc *CategoryController) GetCategories(c *gin.Context) {
	categories, err := uc.categoryUseCase.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// PostCategories godoc
// @Summary      Post categories
// @Description  Post an category
// @Tags         category
// @Produce      json
// @Success      200  {object}   model.Category
// @Failure      500  {object}  map[string]string
// @Router       /category [post]
func (uc *CategoryController) PostCategories(c *gin.Context) {
	var newCategory dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := uc.categoryUseCase.CreateCategory(newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to post category"})
		return
	}

	c.JSON(http.StatusCreated, category)

}

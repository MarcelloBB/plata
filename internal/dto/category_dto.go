package dto

type CreateCategoryRequest struct {
	Name   string `json:"name" binding:"required,max=50"`
	UserID int    `json:"user_id" binding:"required"`
}

type CategoryResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

package usecase

import (
	"github.com/MarcelloBB/plata/internal/dto"
	"github.com/MarcelloBB/plata/internal/model"
	"github.com/MarcelloBB/plata/internal/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (p *UserUseCase) GetUsers() ([]model.User, error) {
	return p.repository.GetUsers()
}

func (p *UserUseCase) CreateUser(user dto.CreateUserRequest) (dto.UserResponse, error) {
	userModel := model.User{
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
	}

	newUser, err := p.repository.CreateUser(userModel)
	if err != nil {
		return dto.UserResponse{}, err
	}

	response := dto.UserResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Bio:      newUser.Bio,
	}

	return response, nil
}

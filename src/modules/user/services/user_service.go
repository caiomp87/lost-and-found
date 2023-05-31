package services

import (
	"context"

	"github.com/caiomp87/lost-and-found/src/modules/user/dto"
	"github.com/caiomp87/lost-and-found/src/modules/user/entities"
	"github.com/caiomp87/lost-and-found/src/modules/user/repositories"
)

type IUserService interface {
	Create(ctx context.Context, user *dto.CreateUserRequest) error
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Create(ctx context.Context, userDto *dto.CreateUserRequest) error {
	user := entities.NewUser(userDto.Name, userDto.Email, userDto.Password)

	return s.repository.Create(ctx, &user)
}

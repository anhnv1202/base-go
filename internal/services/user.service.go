package services

import (
	"github.com/anhnv1202/base-go/internal/repositories"
)

type UserService struct{
	repo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepo(),
	}
}

func (uc UserService) GetOne() string{
    return uc.repo.GetInfoUser()
}

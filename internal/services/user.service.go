package services

import "github.com/anhnv1202/base-go/internal/repositories"

// type UserService struct{
// 	repo *repositories.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		repo: repositories.NewUserRepo(),
// 	}
// }

// func (uc UserService) GetOne() string{
//     return uc.repo.GetInfoUser()
// }

//INTERFACE_VERSION

type IUserService interface{
	GetOne() string
    Register(email string, purpose string) int
}

type userService struct{
    repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) IUserService {
    return &userService{
        repo: repo,
    }
}

func (us *userService) GetOne() string {
    return us.repo.GetInfoUser()
}

func (us *userService) Register(email string, purpose string) int {
    exists := us.repo.GetUserByEmail(email)
    if exists {
        return 409
    }
    return 201
}

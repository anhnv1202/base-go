package repositories

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "vanhne hehe"
// }


type IUserRepository interface{
    GetInfoUser() string
    GetUserByEmail(email string) bool
}

type userRepository struct{
}

func NewUserRepository() IUserRepository {
    return &userRepository{}
}

func (ur *userRepository) GetInfoUser() string {
    return "vanhne hehe"
}

func (ur *userRepository) GetUserByEmail(email string) bool {
    return false
}
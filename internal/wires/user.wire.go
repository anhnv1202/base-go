//go:build wireinject

package wires

import (
	"github.com/anhnv1202/base-go/internal/controllers"
	"github.com/anhnv1202/base-go/internal/repositories"
	"github.com/anhnv1202/base-go/internal/services"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		services.NewUserService,
		repositories.NewUserRepository,
		controllers.NewUserController,
	)
     return new(controllers.UserController), nil
}
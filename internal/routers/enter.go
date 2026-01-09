package routers

import (
	"github.com/anhnv1202/base-go/internal/routers/manager"
	"github.com/anhnv1202/base-go/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouteGroup
	Manager manager.ManagerRouteGroup
}
var RouterGroupApp = new(RouterGroup)
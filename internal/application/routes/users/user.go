package routerUsers

import (
	controller "srwilliamg/app/v1/internal/application/controller"
	"srwilliamg/app/v1/internal/application/request"

	"github.com/go-chi/chi"
)

type User struct {
	router         *chi.Mux
	userController controller.UserController
}

func NewUserRouter(router *chi.Mux, controller controller.UserController) *User {
	userRouter := &User{
		router:         router,
		userController: controller,
	}

	userRouter.initRoutes()
	return userRouter
}

func (u *User) initRoutes() {

	u.router.Mount("/v1/users", u.UsersRoute())
}

func (u *User) UsersRoute() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", request.WithReqHandlerWrapper((u.userController).GetUsers))
	router.Post("/", request.WithReqHandlerWrapper((u.userController).CreateUser))

	return router
}

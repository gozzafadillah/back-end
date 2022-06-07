package routes

import (
	"ppob/helper/valid"
	handler_users "ppob/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler   handler_users.UsersHandler
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// product public
	e.POST("/login", cl.UserHandler.Authorization)
	e.POST("/register", cl.UserHandler.Register)
	authUser := e.Group("user")
	authUser.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("customer", cl.UserHandler))
	authUser.POST("/account", cl.UserHandler.InsertAccount)
}

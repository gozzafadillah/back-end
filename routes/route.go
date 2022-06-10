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

const server = "masukan-server-disini"

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// product public
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	// product public
	e.POST("/login", cl.UserHandler.Authorization)
	e.POST("/register", cl.UserHandler.Register)
	e.GET("/admin/users", cl.UserHandler.GetUsers)
	e.GET("/admin/user/:phone", cl.UserHandler.GetUserForAdmin)
	// product public
	authUser := e.Group("user")
	authUser.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	authUser.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("customer", cl.UserHandler))
	authUser.POST("/account", cl.UserHandler.InsertAccount)
	authUser.GET("/profile", cl.UserHandler.GetUserForCustomer)
	authUser.POST("/profile", cl.UserHandler.UpdateProfile)
}

package routes

import (
	handler_products "ppob/products/handler"
	handler_users "ppob/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserHandler     handler_users.UsersHandler
	ProductsHandler handler_products.ProductsHandler
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// public
	e.POST("/login", cl.UserHandler.Authorization)
	e.POST("/register", cl.UserHandler.Register)
	// manage product endpoint
	e.POST("/product", cl.ProductsHandler.InsertProduct)
	e.GET("/products", cl.ProductsHandler.GetAllProduct)
	e.GET("/product/:id", cl.ProductsHandler.GetProduct)
}

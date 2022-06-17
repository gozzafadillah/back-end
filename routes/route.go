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

	// product public
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"server-front-end"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))

	e.POST("/login", cl.UserHandler.Authorization)
	e.POST("/register", cl.UserHandler.Register)
	// manage product endpoint
	e.POST("/product", cl.ProductsHandler.InsertProduct)
	e.GET("/products", cl.ProductsHandler.GetAllProduct)
	e.GET("/products/:category_id", cl.ProductsHandler.GetProductByCategory)
	e.GET("/product/:id", cl.ProductsHandler.GetProduct)
	e.PUT("/product/:id", cl.ProductsHandler.EditProduct)
	e.DELETE("/product/:id", cl.ProductsHandler.DestroyProduct)
	// manage detail product
	e.GET("/product/detail/:code", cl.ProductsHandler.GetDetailsProduct)
	e.POST("/product/detail/:code", cl.ProductsHandler.InsertDetail)
	e.PUT("/detail/:getID", cl.ProductsHandler.EditDetail)
	e.DELETE("/detail/:getID", cl.ProductsHandler.DestroyDetail)
	// manage category
	e.GET("/category", cl.ProductsHandler.GetCategories)
	e.POST("/category", cl.ProductsHandler.InsertCategory)
	e.PUT("/category/:id", cl.ProductsHandler.EditCategory)
	e.DELETE("/category/:id", cl.ProductsHandler.DestroyCategory)
}

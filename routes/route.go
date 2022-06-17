package routes

import (

	"ppob/helper/valid"

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

const server = "http://localhost:3000"

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
	e.GET("/admin/users/:phone", cl.UserHandler.GetUserForAdmin)
	// validasi
	e.POST("/validation", cl.UserHandler.VerifUser)
	// product public
	authUser := e.Group("users")
	authUser.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	authUser.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("customer", cl.UserHandler))
	// buat pin
	authUser.POST("/pin", cl.UserHandler.InsertAccount)
	authUser.GET("/profile", cl.UserHandler.GetUserSession)
	authUser.POST("/profile", cl.UserHandler.UpdateProfile)
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

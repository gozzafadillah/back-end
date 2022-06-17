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

	// access public
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	e.POST("/login", cl.UserHandler.Authorization)
	e.POST("/register", cl.UserHandler.Register)
	e.GET("/admin/users", cl.UserHandler.GetUsers)
	e.GET("/admin/users/:phone", cl.UserHandler.GetUserForAdmin)
	// validasi
	e.POST("/validation", cl.UserHandler.VerifUser)
	// access public product
	e.GET("/products", cl.ProductsHandler.GetAllProduct)
	e.GET("/products/:category_id", cl.ProductsHandler.GetProductByCategory)
	e.GET("/product/:id", cl.ProductsHandler.GetProduct)

	// access customer
	authUser := e.Group("users")
	authUser.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	authUser.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("customer", cl.UserHandler))
	// make pin
	authUser.POST("/pin", cl.UserHandler.InsertAccount)
	authUser.GET("/profile", cl.UserHandler.GetUserSession)
	authUser.POST("/profile", cl.UserHandler.UpdateProfile)

	// manage product endpoint (admin)
	authAdmin := e.Group("products")
	authAdmin.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("admin", cl.UserHandler))
	authAdmin.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	authAdmin.POST("/", cl.ProductsHandler.InsertProduct)
	authAdmin.PUT("/:id", cl.ProductsHandler.EditProduct)
	authAdmin.DELETE("/:id", cl.ProductsHandler.DestroyProduct)
	// manage detail product (admin)
	authAdmin.GET("/detail/:code", cl.ProductsHandler.GetDetailsProduct)
	authAdmin.POST("/detail/:code", cl.ProductsHandler.InsertDetail)
	authAdmin.PUT("/detail/:getID", cl.ProductsHandler.EditDetail)
	authAdmin.DELETE("/detail/:getID", cl.ProductsHandler.DestroyDetail)
	// manage category (admin)
	e.GET("/category", cl.ProductsHandler.GetCategories)
	authAdmin.POST("/category", cl.ProductsHandler.InsertCategory)
	authAdmin.PUT("/category/:id", cl.ProductsHandler.EditCategory)
	authAdmin.DELETE("/category/:id", cl.ProductsHandler.DestroyCategory)

}

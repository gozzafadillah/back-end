package routes

import (
	"ppob/app/middlewares"
	"ppob/helper/valid"
	handler_products "ppob/products/handler"
	handler_transaction "ppob/transaction/handler"
	handler_users "ppob/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserHandler        handler_users.UsersHandler
	ProductsHandler    handler_products.ProductsHandler
	TransactionHandler handler_transaction.TransactionHandler
}

const server = "http://localhost:3000"

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// log
	middlewares.LogMiddleware(e)

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
	e.GET("/products/category/:category_id", cl.ProductsHandler.GetProductByCategory)
	e.GET("/products/:id", cl.ProductsHandler.GetProduct)
	e.GET("/detail/:product_slug", cl.ProductsHandler.GetDetailsProduct)
	e.GET("/category", cl.ProductsHandler.GetCategories)
	e.POST("/transaction/callback_invoice", cl.TransactionHandler.Callback_Invoice)

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
	authUser.POST("/pin", cl.UserHandler.MakePin)
	authUser.GET("/session", cl.UserHandler.GetUserSession)
	authUser.POST("/profile", cl.UserHandler.UpdateProfile)
	// transaction
	authUser.POST("/checkout/:detail_slug", cl.TransactionHandler.Checkout)
	authUser.GET("/history", cl.TransactionHandler.GetHistoryTransaction)
	authUser.GET("/favorite", cl.TransactionHandler.FavoriteUser)

	// manage product endpoint (admin)
	authAdmin := e.Group("admin")
	authAdmin.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("admin", cl.UserHandler))
	authAdmin.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{server},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	authAdmin.POST("/products/:category_id", cl.ProductsHandler.InsertProduct)
	authAdmin.PUT("/products/edit/:id", cl.ProductsHandler.EditProduct)
	authAdmin.DELETE("/products/delete/:id", cl.ProductsHandler.DestroyProduct)
	// manage detail product (admin)

	authAdmin.POST("/detail/:product_slug", cl.ProductsHandler.InsertDetail)
	authAdmin.PUT("/detail/edit/:getID", cl.ProductsHandler.EditDetail)
	authAdmin.DELETE("/detail/delete/:getID", cl.ProductsHandler.DestroyDetail)
	// manage category (admin)

	authAdmin.POST("/category", cl.ProductsHandler.InsertCategory)
	authAdmin.PUT("/category/edit/:id", cl.ProductsHandler.EditCategory)
	authAdmin.DELETE("/category/delete/:id", cl.ProductsHandler.DestroyCategory)

}

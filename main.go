package main

import (
	"ppob/app/config"
	"ppob/app/middlewares"
	migrate "ppob/migrator"
	handler_products "ppob/products/handler"
	mysql_products "ppob/products/repository/mysql"
	service_products "ppob/products/service"
	"ppob/routes"
	handler_users "ppob/users/handler"
	mysql_users "ppob/users/repository/mysql"
	service_users "ppob/users/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	migrate.AutoMigrate(db)

	configJWT := middlewares.ConfigJwt{
		SecretJWT: config.Conf.JWTSecret,
	}

	e := echo.New()

	//Factory
	// Users
	userRepo := mysql_users.NewUsersRepo(db)
	userServ := service_users.NewUsersService(userRepo, &configJWT)
	UserHandler := handler_users.NewUsersHandler(userServ)
	// Products
	productRepo := mysql_products.NewProductsRepository(db)
	productServ := service_products.NewProductsService(productRepo)
	productsHandler := handler_products.NewProductsHandler(productServ)

	// Route
	routeInit := routes.ControllerList{
		JWTMiddleware:   configJWT.Init(),
		UserHandler:     UserHandler,
		ProductsHandler: productsHandler,
	}

	routeInit.RouteRegister(e)
	// start server
	e.Logger.Fatal(e.Start(":8080"))

}

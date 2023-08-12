package main

import (
	handler_admin "ppob/admin/handler"
	"ppob/app/config"
	"ppob/app/middlewares"
	migrate "ppob/migrator"
	handler_products "ppob/products/handler"
	mysql_products "ppob/products/repository/mysql"
	service_products "ppob/products/service"
	"ppob/routes"
	handler_transaction "ppob/transaction/handler"
	mysql_transaction "ppob/transaction/repository/mysql"
	service_transaction "ppob/transaction/service"
	handler_users "ppob/users/handler"
	mysql_users "ppob/users/repository/mysql"
	service_users "ppob/users/service"
	"os"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
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
	// Transaction
	transactionRepo := mysql_transaction.NewTransactionRepo(db)
	transactionServ := service_transaction.NewTransactionService(transactionRepo)
	transactionHandler := handler_transaction.NewTransactionHandler(transactionServ, productServ, userServ)
	// Admin
	adminHandler := handler_admin.NewAdminHandler(transactionServ, productServ, userServ)

	// Route
	routeInit := routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserHandler:        UserHandler,
		ProductsHandler:    productsHandler,
		TransactionHandler: transactionHandler,
		AdminHandler:       adminHandler,
	}

	routeInit.RouteRegister(e)
	// start server
	e.Logger.Fatal(e.Start(":"+port))

}

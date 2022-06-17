package migrate

import (
	mysql_products "ppob/products/repository/mysql"
	mysql_users "ppob/users/repository/mysql"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&mysql_users.Users{}, &mysql_products.Products{}, &mysql_products.Category_Product{}, &mysql_products.Detail_Product{})
}

package migrate

import (
	mysql_users "ppob/users/repository/mysql"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&mysql_users.Users{}, &mysql_users.Account{}, &mysql_users.UserVerif{})
}

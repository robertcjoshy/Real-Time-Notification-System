package utils

import (
	"fmt"

	"github.com/robert/notification/admin/admin_model"
	"github.com/robert/notification/user/user_model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Opendatabaseconnection() {

	var err error

	username := "postgres"
	password := "1234"
	databaseName := "startup"
	ssl := "disable"
	zone := "Asia/Shanghai"

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode = %s TimeZone=%s", username, password, databaseName, ssl, zone)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}

	Migrate()
}

func Migrate() {
	Database.AutoMigrate(user_model.User{}, admin_model.Admin{})
}

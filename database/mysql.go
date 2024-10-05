package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MYSQL *gorm.DB

func DatabaseInitMysql() {
	var err error
	dsn := "web66_65011212007:65011212007@csmsu@tcp(202.28.34.197)/web66_65011212007?charset=utf8mb4&parseTime=True&loc=Local"

	MYSQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("\n\n Cannot connect to database")
	}
	fmt.Println("\n\n Connect to database")
}

package database

import (
	"fmt"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Connects to the MySQL database using GORM.
func Connection() (db *gorm.DB, err error) {
	dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	Db = db

	db.AutoMigrate(models.User{})

	// Ensure that the database connection is closed.

	return db, nil
}

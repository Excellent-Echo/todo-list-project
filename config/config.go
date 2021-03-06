package config

import (
	"fmt"
	"os"
	"todoAPIGolang/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	// err := godotenv.Load()

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// dsn := "Q6sEcExPoi:z0afMi60oe@tcp(remotemysql.com:3306)/Q6sEcExPoi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&migration.User{})
	db.AutoMigrate(&migration.Category{})
	db.AutoMigrate(&migration.Todo{})
	db.AutoMigrate(&migration.UserDetail{})
	db.AutoMigrate(&migration.UserProfile{})

	return db
}

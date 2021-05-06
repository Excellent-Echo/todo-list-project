package config

import (
	"todoAPIGolang/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "root:@tcp(localhost)/todo_db_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Todo{})
	db.AutoMigrate(&entity.UserDetail{})

	return db
}

package main

import (
	"github.com/kabos0809/redis_jwt_gin/Models"
	"github.com/kabos0809/redis_jwt_gin/Routes"
	"github.com/kabos0809/redis_jwt_gin/Config"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	dsn := Config.DBUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&Models.User{})

	ctrl := Models.Model{Db: db}

	router := Routes.SetRoutes(ctrl)

	router.Run(":8080")
}
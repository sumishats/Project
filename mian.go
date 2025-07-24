package main

import (
	"sample/model"
	"sample/route"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=sumisha@2006 dbname=testpsql port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.User{})

	r := gin.Default()

	route.SetupRoutes(r, db)

	r.Run(":8080")

}

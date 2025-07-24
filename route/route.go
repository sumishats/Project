package route

import (
	"sample/handler"
	"sample/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetupRoutes(r *gin.Engine,db *gorm.DB){
	userHandler:=handler.NewUserHandler(db)
	adminHandler:=handler.NewAdminHandler(db)

	r.POST("/signup",userHandler.Signup) 
	r.POST("/login",userHandler.Login)
	r.POST("/admin/login",adminHandler.AdminLogin)
	protected:=r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	adminroutes:=r.Group("/admin")
	adminroutes.Use(middleware.AuthMiddleware(),middleware.AdminMiddleware())
}
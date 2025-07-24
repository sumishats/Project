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
	{
		protected.GET("/home",adminHandler.Home)
	}
	adminRoutes:=r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(),middleware.AdminMiddleware())
	
	{
		adminRoutes.GET("/users", adminHandler.GetUsers)
		adminRoutes.POST("/users", adminHandler.CreateUser)
		adminRoutes.GET("/users/:id", adminHandler.GetUser)
		adminRoutes.PUT("/users/:id", adminHandler.UpdateUser)
		adminRoutes.DELETE("/users/:id", adminHandler.DeleteUser)
		
	}
}
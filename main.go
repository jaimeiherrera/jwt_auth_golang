package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaimeiherrera/jwt_auth_golang/controllers"
	"github.com/jaimeiherrera/jwt_auth_golang/initializers"
	"github.com/jaimeiherrera/jwt_auth_golang/middleware"
)

func init() {
	initializers.Load_env()
	initializers.Connect_db()
	initializers.Migrate_db()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Sign_up)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Required_auth, controllers.Validate)
	r.Run()
}

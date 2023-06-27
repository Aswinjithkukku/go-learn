package main

import (
	"fmt"
	"os"

	"github.com/aswinjithkukku/jwt-auth/controllers"
	"github.com/aswinjithkukku/jwt-auth/initializer"
	"github.com/aswinjithkukku/jwt-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
	initializer.SyncDatabse()
}

func main() {
	fmt.Println("Hello world!!")

	r := gin.Default()

	r.POST("/user/signup", controllers.SignUp)
	r.POST("/user/signin", controllers.SignIn)
	r.GET("/user/me", middlewares.RequireAuth, controllers.Validate)

	r.Run(":" + os.Getenv("PORT"))
}

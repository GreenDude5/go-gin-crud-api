package main

import (
	"go-gin-crud/controllers"
	"go-gin-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)

	router.Run()
}

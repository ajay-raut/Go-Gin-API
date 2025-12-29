package main

import (
	initializers "github.com/ajay-raut/gocrud/Initializers"
	"github.com/ajay-raut/gocrud/controllers"
	"github.com/ajay-raut/gocrud/migrations"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	migrations.MigrateModels()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.CreatePost)

	router.GET("/posts", controllers.PostsIndex)

	router.GET("/posts/:id", controllers.PostShow)

	router.PUT("/posts/:id", controllers.PostUpdate)

	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run() // listen and serve on 0.0.0.0:3001
}

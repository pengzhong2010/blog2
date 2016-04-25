package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"blog2/controllers"
	// "net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.Home)
	router.GET("/category", controllers.Category)
	router.GET("/blog", controllers.Blog)
	router.GET("/topic", controllers.Topic)

	router.Run(":8080")

}

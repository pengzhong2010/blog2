package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"blog2/admin"
	"blog2/controllers"
	// "net/http"
)

func main() {

	router := gin.Default()
	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.Home)
	router.GET("/category", controllers.Category)
	router.GET("/blog", controllers.Blog)
	router.GET("/topic", controllers.Topic)

	v1 := router.Group("/admin")
	{
		v2 := v1.Group("/user")
		{
			v2.GET("/list", admin.UserList)
			v2.GET("/add", admin.UserAdd)
			v2.POST("/add", admin.UserAddDo)
			v2.GET("/edit/:id", admin.UserEdit)
			v2.POST("/edit/:id", admin.UserEditDo)
			v2.GET("/del/:id", admin.UserDel)
		}
		v3 := v1.Group("/category")
		{
			v3.GET("/list", admin.CategoryList)
			v3.GET("/add", admin.CategoryAdd)
			v3.POST("/add", admin.CategoryAddDo)
			v3.GET("/edit/:id", admin.CategoryEdit)
			v3.POST("/edit/:id", admin.CategoryEditDo)
			v3.GET("/del/:id", admin.CategoryDel)
		}
		v4 := v1.Group("/topic")
		{
			v4.GET("/list", admin.TopicList)
			v4.GET("/add", admin.TopicAdd)
			v4.POST("/add", admin.TopicAddDo)
			v4.GET("/edit/:id", admin.TopicEdit)
			v4.POST("/edit/:id", admin.TopicEditDo)
			v4.GET("/del/:id", admin.TopicDel)
		}
		v1.GET("/base", admin.Base)
	}

	router.Run(":8080")

}

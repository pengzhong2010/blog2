package controllers

import (
	_ "blog2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
	})
}

func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "category.html", gin.H{
		"title": "Home",
	})
}

func Blog(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title": "Home",
	})
}

func Topic(c *gin.Context) {
	c.HTML(http.StatusOK, "topic.html", gin.H{
		"title": "Home",
	})
}

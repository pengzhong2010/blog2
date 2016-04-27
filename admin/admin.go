package admin

import (
	"blog2/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	// "os"
	// "text/template"
)

var i *gorm.DB = models.I

func Base(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{})
}
func UserList(c *gin.Context) {
	us := []models.User{}
	i.Find(&us)
	log.Println(us)
	// c.String(http.StatusOK, "ok1")
	// user := models.User{}
	// user.List()
	c.HTML(http.StatusOK, "userlist.html", gin.H{
		"users": us,
	})

}

func UserAdd(c *gin.Context) {

	c.HTML(http.StatusOK, "useradd.html", gin.H{})

}

type UserAddForm struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"pwd" binding:"required"`
}

func UserAddDo(c *gin.Context) {
	var form UserAddForm
	if c.Bind(&form) == nil {
		u := models.User{}

		i.FirstOrCreate(&u, models.User{Email: form.Email, Pwd: form.Password})
		c.Redirect(http.StatusMovedPermanently, "/admin/user/list")
	}

}

func UserEdit(c *gin.Context) {
	id := c.Param("id")
	u := models.User{}
	i.First(&u, id)
	if u.ID == 0 {
		c.String(http.StatusOK, "not found")
	} else {
		c.HTML(http.StatusOK, "useradd.html", gin.H{
			"user": u,
		})
	}

}

func UserEditDo(c *gin.Context) {
	id := c.Param("id")
	var form UserAddForm
	if c.Bind(&form) == nil {
		u := models.User{}

		i.First(&u, id)
		u.Pwd = form.Password
		i.Save(&u)
		c.Redirect(http.StatusMovedPermanently, "/admin/user/list")
	}

}

func UserDel(c *gin.Context) {
	// c.String(http.StatusOK, "ok111")
	id := c.Param("id")
	u := models.User{}
	i.First(&u, id)
	log.Println(u)
	log.Println(u.ID)
	if u.ID != 0 {
		i.Delete(&u)
	}

	c.Redirect(http.StatusMovedPermanently, "/admin/user/list")

}

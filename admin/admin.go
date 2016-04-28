package admin

import (
	"blog2/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
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

func CategoryList(c *gin.Context) {
	us := []models.Category{}
	i.Find(&us)
	log.Println(us)
	// c.String(http.StatusOK, "ok1")
	// user := models.User{}
	// user.List()
	c.HTML(http.StatusOK, "categorylist.html", gin.H{
		"category": us,
	})

}

func CategoryAdd(c *gin.Context) {

	c.HTML(http.StatusOK, "categoryadd.html", gin.H{})

}

type CategoryAddForm struct {
	Name string `form:"name" binding:"required"`
}

func CategoryAddDo(c *gin.Context) {
	var form CategoryAddForm
	if c.Bind(&form) == nil {
		u := models.Category{}

		i.FirstOrCreate(&u, models.Category{Name: form.Name})
		c.Redirect(http.StatusMovedPermanently, "/admin/category/list")
	}
}

func CategoryEdit(c *gin.Context) {
	id := c.Param("id")
	u := models.Category{}
	i.First(&u, id)
	if u.ID == 0 {
		c.String(http.StatusOK, "not found")
	} else {
		c.HTML(http.StatusOK, "categoryadd.html", gin.H{
			"category": u,
		})
	}

}

func CategoryEditDo(c *gin.Context) {
	id := c.Param("id")
	var form CategoryAddForm
	if c.Bind(&form) == nil {
		u := models.Category{}

		i.First(&u, id)
		u.Name = form.Name
		i.Save(&u)
		c.Redirect(http.StatusMovedPermanently, "/admin/category/list")
	}
}

func CategoryDel(c *gin.Context) {
	// c.String(http.StatusOK, "ok111")
	id := c.Param("id")
	u := models.Category{}
	i.First(&u, id)
	log.Println(u)
	log.Println(u.ID)
	if u.ID != 0 {
		i.Delete(&u)
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/category/list")
}

func TopicList(c *gin.Context) {
	us := []models.Topic{}
	i.Find(&us)
	topics := make(map[int]map[string]string)
	for j := 0; j < len(us); j++ {
		u := us[j]
		log.Println(u.UserID)
		tu := models.User{}
		i.First(&tu, u.UserID)
		log.Println(tu)
		topics[j] = make(map[string]string)
		uid := strconv.Itoa(int(u.ID))
		topics[j]["ID"] = uid
		topics[j]["Name"] = u.Name
		topics[j]["Email"] = tu.Email
	}
	log.Println(topics)
	// rows, err := i.Table("topics").Select("topics.name,users.email").Joins("left join users on topics.user_id = users.id").Rows()
	// i.Find(&us)
	// log.Println(rows)
	// for rows.Next() {
	// 	log.Println(rows.Columns())
	// }
	// log.Println(err)
	// c.String(http.StatusOK, "ok1")
	// user := models.User{}
	// user.List()
	c.HTML(http.StatusOK, "topiclist.html", gin.H{
		"topics": topics,
	})

}

func TopicAdd(c *gin.Context) {

	us := []models.User{}
	i.Find(&us)
	c.HTML(http.StatusOK, "topicadd.html", gin.H{
		"users": us,
	})

}

type TopicAddForm struct {
	Name    string `form:"name" binding:"required"`
	Content string `form:"content"`
	UserID  string `form:"userid" binding:"required"`
}

func TopicAddDo(c *gin.Context) {
	var form TopicAddForm
	if c.Bind(&form) == nil {
		u := models.Topic{}
		uid1 := form.UserID
		uid, err := strconv.Atoi(uid1)
		if err != nil {
			c.String(http.StatusOK, "error")
		} else {
			i.FirstOrCreate(&u, models.Topic{Name: form.Name, Content: form.Content, UserID: uid})
			c.Redirect(http.StatusMovedPermanently, "/admin/topic/list")
		}
	}
}

func TopicEdit(c *gin.Context) {
	id := c.Param("id")
	u := models.Topic{}
	i.First(&u, id)
	log.Println(u)
	if u.ID == 0 {
		c.String(http.StatusOK, "not found")
	} else {
		us := models.User{}
		i.First(&us, u.UserID)
		log.Println(us)

		if us.ID == 0 {
			c.String(http.StatusOK, "user not found")
		} else {
			uss := []models.User{}
			i.Find(&uss)
			log.Println(uss)
			c.HTML(http.StatusOK, "topicadd.html", gin.H{
				"users": uss,
				"topic": u,
				"uid":   us.ID,
			})
		}

	}

}

func TopicEditDo(c *gin.Context) {
	id := c.Param("id")
	var form TopicAddForm
	if c.Bind(&form) == nil {
		u := models.Topic{}

		i.First(&u, id)
		u.Name = form.Name
		u.Content = form.Content
		uid1 := form.UserID
		uid, _ := strconv.Atoi(uid1)
		u.UserID = uid
		i.Save(&u)
		c.Redirect(http.StatusMovedPermanently, "/admin/topic/list")
	}
}

func TopicDel(c *gin.Context) {
	// c.String(http.StatusOK, "ok111")
	id := c.Param("id")
	u := models.Topic{}
	i.First(&u, id)
	log.Println(u)
	log.Println(u.ID)
	if u.ID != 0 {
		i.Delete(&u)
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/topic/list")
}

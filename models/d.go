package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	gorm.Model
	Email  string `gorm:"size:50;unique"`
	Pwd    string
	Topics []Topic
	Entrys []Entry
}

type Category struct {
	gorm.Model
	Name   string
	Topics []Topic `gorm:"many2many:category_topic;"`
}

type Topic struct {
	gorm.Model
	Name    string
	Content string
	UserID  int
	Entrys  []Entry
}
type Entry struct {
	gorm.Model
	EntryID int
	UserID  int
	Content string
}

var I *gorm.DB

func init() {
	var err error
	I, err = gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("%v", err)
	}
	I.LogMode(true)
	I.AutoMigrate(&User{}, &Category{}, &Topic{}, &Entry{})

}

// func (u *User) List() {
// 	a := i.Find(&u)
// 	log.Println(a)
// 	i.First(&u)
// 	log.Println(u)
// }

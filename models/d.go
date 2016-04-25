package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Email  string `gorm:"size:50"`
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

func init() {
	db, _ := gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{}, &Category{}, &Topic{}, &Entry{})

}

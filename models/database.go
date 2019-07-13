package models

import (
	"github.com/double1996/smart-evernote-blog/config"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB(config config.Configuration) (*gorm.DB, error) {
	var db = DB

	driver := config.DataBase.Driver
	database := config.DataBase.DbName
	username := config.DataBase.Username
	password := config.DataBase.Password
	host := config.DataBase.Host
	port := config.DataBase.Port

	if driver == "sqlite3" {
		db, err = gorm.Open("sqlite3", "./blog.db")
	} else if driver == "mysql" {
		db, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		return nil, err
	}
	//db.AutoMigrate(&Post{}, &Tag{}, &Comment{})
	//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
	return db, err
}

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

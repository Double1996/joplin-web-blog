package models

import (
	"github.com/double1996/joplin-web-blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB(config config.Configuration) (*gorm.DB, error) {

	driver := config.DataBase.Driver
	database := config.DataBase.DbName
	username := config.DataBase.Username
	password := config.DataBase.Password
	host := config.DataBase.Host
	port := config.DataBase.Port

	if driver == "sqlite3" {
		DB, err = gorm.Open("sqlite3", "./blog.db")
	} else if driver == "mysql" {
		DB, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		return nil, err
	}
	DB.AutoMigrate(&Post{}, &Tag{}, &Comment{})
	//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
	//DB.LogMode(true)
	return DB, err
}

type BaseModel struct {
	ID          uint  `gorm:"primary_key"`
	CreatedTime int64 `json:"created_time"`
	UpdatedTime int64 `json:"updated_time"`
}

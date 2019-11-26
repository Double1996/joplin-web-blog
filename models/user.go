package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email         string `gorm:"unique_index;default:null"` //邮箱
	Telephone     string `gorm:"unique_index;default:null"` //手机号码
	Password      string `gorm:"default:null"`              //密码
	VerifyState   string `gorm:"default:'0'"`               //邮箱验证状态
	SecretKey     string `gorm:"default:null"`              //密钥
	OutTime       time.Time                                 //过期时间
	GithubLoginId string `gorm:"unique_index;default:null"` // github唯一标识
	GithubUrl     string                                    //github地址
	IsAdmin       bool                                      //是否是管理员
	AvatarUrl     string                                    // 头像链接
	NickName      string                                    // 昵称
	LockState     bool `gorm:"default:'0'"`                 //锁定状态
}



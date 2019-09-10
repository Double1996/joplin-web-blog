package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Subscriber struct {
	gorm.DB
	Email           string    `gorm:"unique_index"`
	VerifyState     bool      `gorm:"default:'0'"` // 验证状态
	SubscriberState bool      `gorm:"default:'1'"` // 订阅状态
	OutTime         time.Time // 过期时间
	SecretKey       string    // 秘钥
	Signature       string    // 签名
}

func (s *Subscriber) Insert() error {
	return DB.FirstOrCreate(s, "email=?", s.Email).Error
}

//func (s *Subscriber) Update() error	{
//	return
//}

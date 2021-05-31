// @File  : rabbitmq_error.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/6/1
// @slogan: 又是不想写代码的一天，神兽保佑，代码无BUG！
//         ┏┓      ┏┓
//        ┏┛┻━━━━━━┛┻┓
//        ┃     ღ    ┃
//        ┃  ┳┛   ┗┳ ┃
//        ┃     ┻    ┃
//        ┗━┓      ┏━┛
//          ┃      ┗━━━┓
//          ┃ 神兽咆哮!  ┣┓
//          ┃         ┏┛
//          ┗┓┓┏━━━┳┓┏┛
//           ┃┫┫   ┃┫┫
//           ┗┻┛   ┗┻┛

package model

import (
	"gorm.io/gorm"
	"time"
)

type AvfRabbitmqError struct {
	ID         int64     `gorm:"column:id" json:"id" form:"id"`
	Message    string    `gorm:"column:message" json:"message" form:"message"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
}

func (h *AvfRabbitmqError) TableName() string {
	return "avf_rabbitmq_error"
}

func (h *AvfRabbitmqError) Create(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}

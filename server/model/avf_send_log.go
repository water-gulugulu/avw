// @File  : avf_send_log.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/27
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
)

type AvfSendLog struct {
	ID        int64  `gorm:"column:id" db:"column:id" json:"id" form:"id"`
	Address   string `gorm:"column:address" db:"column:address" json:"address" form:"address"`
	IsSuccess int64  `gorm:"column:is_success" db:"column:is_success" json:"is_success" form:"is_success"`
}

func (h *AvfSendLog) TableName() string {
	return "avf_send_log"
}

func (h *AvfSendLog) Create(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}

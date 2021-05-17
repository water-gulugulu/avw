// 自动生成模板AvfOrderCard
package model

import (
	"errors"
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type AvfOrderCard struct {
	global.GVA_MODEL
	Uid        int     `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;type:int;size:10;"`                           // 用户ID
	OrderId    int     `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;type:int;size:10;"`              // 订单ID
	CardId     int     `json:"cardId" form:"cardId" gorm:"column:card_id;comment:卡牌ID;type:int;size:10;"`                 // 卡牌ID
	Star       int     `json:"star" form:"star" gorm:"column:star;comment:算力值;type:int;size:10;"`                         // 算力值
	Status     int     `json:"status" form:"status" gorm:"column:status;comment:状态 1-正常 2-转让中;type:int;size:10;"`         // 状态 1-正常 2-转让中
	CreateTime int     `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:int;size:10;"`     // 创建时间
	UpdateTime int     `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:int;size:10;"`     // 修改时间
	GiveType   int     `json:"giveType" form:"giveType" gorm:"column:give_type;comment:获得方式 1-抽奖 2-购买;type:int;size:10;"` // 获得方式 1-抽奖 2-购买
	Level      int     `json:"level" form:"level" gorm:"column:level;comment:1-N 2-R 3-SR 4-SSR;type:int;size:10;"`       // 等级 1-N 2-R 3-SR 4-SSR
	Card       AvfCard `json:"card"`                                                                                      // 卡牌信息
}

func (h *AvfOrderCard) TableName() string {
	return "avf_order_card"
}

func (h *AvfOrderCard) FindListByOrderId(DB *gorm.DB) (list []AvfOrderCard, err error) {
	if h.OrderId == 0 {
		return nil, errors.New("订单ID不能为空")
	}

	if err = DB.Table(h.TableName()).Where("order_id = ?", h.OrderId).Find(&list).Error; err != nil {
		return list, err
	}

	return
}
func (h *AvfOrderCard) FindListByUid(DB *gorm.DB, page, size int) (list []AvfOrderCard, total int64, err error) {
	if h.Uid == 0 {
		return nil, 0, errors.New("用户ID不能为空")
	}
	if page != 0 {
		page = page * size
	}
	DB = DB.Table(h.TableName()).Where("uid = ?", h.Uid)
	if err = DB.Count(&total).Error; err != nil {
		return list, 0, err
	}

	if err = DB.Order("id desc").Preload("Card").Limit(size).Offset(page).Find(&list).Error; err != nil {
		return list, 0, err
	}

	return
}

func (h *AvfOrderCard) CreateOrderCard(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}

func (h *AvfOrderCard) GetById(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id =?", h.ID).Preload("Card").First(&h).Error
}

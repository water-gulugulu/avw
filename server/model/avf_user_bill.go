// 自动生成模板AvfUserBill
package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type AvfUserBill struct {
	global.GVA_MODEL
	Uid        int    `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;type:int;size:10;"`                          // 用户ID
	CardId     int    `json:"cardId" form:"cardId" gorm:"column:card_id;comment:卡牌ID;type:int;size:10;"`                // 卡牌ID
	Address    string `json:"address" form:"address" gorm:"column:address;comment:钱包地址;type:varchar(255);size:255;"`    // 钱包地址
	Type       int    `json:"type" form:"type" gorm:"column:type;comment:类型 1-发放收益 2-转账 3-购买卡牌 4-提现;type:int;size:10;"` // 类型 1-发放收益 2-转账 3-购买卡牌 4-提现 -支付手续费
	Money      int    `json:"money" form:"money" gorm:"column:money;comment:金额;type:bigint;size:19;"`                   // 金额
	Fees       int    `json:"fees" form:"fees" gorm:"column:fees;comment:手续费;type:bigint;size:19;"`                     // 手续费
	Balance    int    `json:"balance" form:"balance" gorm:"column:balance;comment:余额;type:bigint;size:19;"`             // 余额
	Payment    int    `json:"payment" form:"payment" gorm:"column:payment;comment:收入支出 1-收入 2-支出;type:int;size:10;"`    // 收入支出
	PayType    int    `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付方式 1-avw 2-ht;type:int;size:10;"`  // 支付方式 1-avw 2-ht
	Detail     string `json:"detail" form:"detail" gorm:"column:detail;comment:描述;type:varchar(255);size:255;"`         // 详情
	CreateTime int    `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:int;size:10;"`    // 创建时间
}

func (AvfUserBill) TableName() string {
	return "avf_user_bill"
}

func (h *AvfUserBill) GetByUid(DB *gorm.DB) (list []AvfUserBill, err error) {
	DB = DB.Table(h.TableName()).Where("uid = ?", h.Uid)
	if h.CreateTime != 0 {
		DB = DB.Where("create_time between ? and ?", h.CreateTime, h.CreateTime+86400)
	}
	if err = DB.Order("id desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return
}

func (h *AvfUserBill) GetByUidAndCardId(DB *gorm.DB) (list []AvfUserBill, err error) {
	DB = DB.Table(h.TableName()).Where("uid = ? and card_id = ?", h.Uid, h.CardId)
	if h.CreateTime != 0 {
		DB = DB.Where("create_time between ? and ?", h.CreateTime, h.CreateTime+86400)
	}
	if h.PayType != 0 {
		DB = DB.Where("pay_type = ?", h.PayType)
	}
	if h.Type != 0 {
		DB = DB.Where("type = ?", h.Type)
	}
	if err = DB.Order("id desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return
}

// 自动生成模板AvfUserBill
package model

import (
	"fmt"
	"gin-vue-admin/global"
	"gorm.io/gorm"
	"strings"
)

// 如果含有time.Time 请自行import time包
type AvfUserBill struct {
	global.GVA_MODEL
	Uid        int     `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;type:int;size:10;"`                                  // 用户ID
	CardId     int     `json:"cardId" form:"cardId" gorm:"column:card_id;comment:卡牌ID;type:int;size:10;"`                        // 卡牌ID
	Address    string  `json:"address" form:"address" gorm:"column:address;comment:钱包地址;type:varchar(255);size:255;"`            // 钱包地址
	Type       int     `json:"type" form:"type" gorm:"column:type;comment:类型 1-发放收益 2-盲盒 3-购买卡牌 4-手续费 5-直推收益;type:int;size:10;"` // 类型 1-发放收益 2-盲盒 3-购买卡牌 4-手续费 5-直推收益
	Money      float64 `json:"money" form:"money" gorm:"column:money;comment:金额;type:decimal;size:9,4;"`                         // 金额
	Fees       float64 `json:"fees" form:"fees" gorm:"column:fees;comment:手续费;type:decimal;size:9,4;"`                           // 手续费
	Balance    float64 `json:"balance" form:"balance" gorm:"column:balance;comment:余额;type:decimal;size:9,4;"`                   // 余额
	Payment    int     `json:"payment" form:"payment" gorm:"column:payment;comment:收入支出 1-收入 2-支出;type:int;size:10;"`            // 收入支出
	PayType    int     `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付方式 1-avw 2-ht;type:int;size:10;"`          // 支付方式 1-avw 2-ht
	Detail     string  `json:"detail" form:"detail" gorm:"column:detail;comment:描述;type:varchar(255);size:255;"`                 // 详情
	TxHash     string  `json:"tx_hash" form:"tx_hash" gorm:"column:tx_hash;comment:交易hash;type:varchar(255);size:255;"`          // 交易hash
	CreateTime int     `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:int;size:10;"`            // 创建时间
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

func (h *AvfUserBill) Create(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}

func (h *AvfUserBill) GetList(DB *gorm.DB, page, size int, billType string) (list []*AvfUserBill, total int64, err error) {
	DB = DB.Table(h.TableName())
	if page != 0 {
		page = page * size
	}
	if h.Uid != 0 {
		DB = DB.Where("uid = ?", h.Uid)
	}
	if h.CreateTime != 0 {
		DB = DB.Where("create_time between ? and ?", h.CreateTime, h.CreateTime+86400)
	}
	fmt.Printf("billType:%s", billType)
	t := strings.Split(billType, ",")
	if len(billType) != 0 {
		DB = DB.Where("type IN(?)", t)
	}

	if err = DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = DB.Order("id desc").Limit(size).Offset(page).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return
}

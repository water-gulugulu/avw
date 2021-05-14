// 自动生成模板AvfCard
package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type AvfCard struct {
	global.GVA_MODEL
	Name            string    `json:"name" form:"name" gorm:"column:name;comment:卡牌名称;type:varchar(255);size:255;"`
	Star            int       `json:"star" form:"star" gorm:"column:star;comment:星力值;type:int;size:10;"`
	Money           float64   `json:"money" form:"money" gorm:"column:money;comment:卡牌价格;type:decimal;size:9,2;"`
	WalletAddress   string    `json:"walletAddress" form:"walletAddress" gorm:"column:wallet_address;comment:支付钱包地址;type:varchar(255);size:255;"`
	ContractAddress string    `json:"contractAddress" form:"contractAddress" gorm:"column:contract_address;comment:合约地址;type:varchar(255);size:255;"`
	Number          int       `json:"number" form:"number" gorm:"column:number;comment:发行数量;type:int;size:10;"`
	Level           int       `json:"level" form:"level" gorm:"column:level;comment:1-N 2-R 3-SR 4-SSR"`
	Author          string    `json:"author" form:"author" gorm:"column:author;comment:作者;type:varchar(255);size:255;"`
	Desc            string    `json:"desc" form:"desc" gorm:"column:desc;comment:描述;type:varchar(500);size:500;"`
	Status          *bool     `json:"status" form:"status" gorm:"column:status;comment:1-正常 2-禁用;type:tinyint;"`
	CreateDate      time.Time `json:"createDate" form:"createDate" gorm:"column:create_date;comment:创建时间;type:datetime;"`
	UpdateDate      time.Time `json:"updateDate" form:"updateDate" gorm:"column:update_date;comment:修改时间;type:datetime;"`
}

func (AvfCard) TableName() string {
	return "avf_card"
}

func (h AvfCard) GetList(DB *gorm.DB, p, size int) (list []AvfCard, total int64, err error) {
	DB = DB.Table(h.TableName()).Where("status = ?", 1)

	if err := DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := DB.Order("id desc").Limit(size).Offset(p).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return
}

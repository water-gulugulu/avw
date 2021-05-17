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
	Name            string    `json:"name" form:"name" gorm:"column:name;comment:卡牌名称;type:varchar(255);size:255;"`                                   // 卡牌名称
	Star            int       `json:"star" form:"star" gorm:"column:star;comment:星力值;type:int;size:10;"`                                              // 算力值
	Money           float64   `json:"money" form:"money" gorm:"column:money;comment:卡牌价格;type:decimal;size:9,2;"`                                     // 价格
	WalletAddress   string    `json:"walletAddress" form:"walletAddress" gorm:"column:wallet_address;comment:支付钱包地址;type:varchar(255);size:255;"`     // 收款钱包地址
	ContractAddress string    `json:"contractAddress" form:"contractAddress" gorm:"column:contract_address;comment:合约地址;type:varchar(255);size:255;"` // 合约地址
	Number          int       `json:"number" form:"number" gorm:"column:number;comment:发行数量;type:int;size:10;"`                                       // 发行数量
	Level           int       `json:"level" form:"level" gorm:"column:level;comment:1-N 2-R 3-SR 4-SSR"`                                              // 等级 1-N 2-R 3-SR 4-SSR
	Author          string    `json:"author" form:"author" gorm:"column:author;comment:作者;type:varchar(255);size:255;"`                               // 作者
	Desc            string    `json:"desc" form:"desc" gorm:"column:desc;comment:描述;type:varchar(500);size:500;"`                                     // 描述
	Status          *bool     `json:"status" form:"status" gorm:"column:status;comment:1-正常 0-禁用;type:tinyint;"`                                      // 状态 true-正常 false-禁用
	CreateDate      time.Time `json:"createDate" form:"createDate" gorm:"column:create_date;comment:创建时间;type:datetime;"`                             // 创建时间
	UpdateDate      time.Time `json:"updateDate" form:"updateDate" gorm:"column:update_date;comment:修改时间;type:datetime;"`                             // 修改时间
	Image           string    `json:"image" form:"image" gorm:"column:image;comment:图片地址;type:varchar(255);size:255;"`                                // 修改时间
}

func (h *AvfCard) TableName() string {
	return "avf_card"
}

func (h *AvfCard) GetList(DB *gorm.DB, p, size int) (list []AvfCard, total int64, err error) {
	if p > 0 {
		p = p * size
	}

	DB = DB.Table(h.TableName()).Where("status = ?", 1)

	if err := DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := DB.Order("id desc").Limit(size).Offset(p).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return
}
func (h *AvfCard) RandGetByLevel(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("level = ?", h.Level).First(&h).Error
}
func (h *AvfCard) GetById(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id = ?", h.ID).First(&h).Error
}

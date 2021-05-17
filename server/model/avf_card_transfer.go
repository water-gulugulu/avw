// 自动生成模板AvfCardTransfer
package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type AvfCardTransfer struct {
	global.GVA_MODEL
	RecordId int    `json:"record_id" form:"record_id" gorm:"column:record_id;comment:卡牌记录ID;type:int;size:10;"`                      // 卡牌记录ID
	Uid      int    `json:"uid" form:"uid" gorm:"column:uid;comment:出售人ID;type:int;size:10;"`                                         // 出售人ID
	CardId   int    `json:"card_id" form:"card_id" gorm:"column:card_id;comment:卡牌ID;type:int;size:10;"`                              // 卡牌ID
	Price    int    `json:"price" form:"price" gorm:"column:price;comment:卡牌出售价格;type:int;size:10;"`                                  // 价格
	Fees     int    `json:"fees" form:"fees" gorm:"column:fees;comment:手续费;type:int;size:10;"`                                        // 手续费
	BuyId    int    `json:"buy_id" form:"buy_id" gorm:"column:buy_id;comment:购买人ID;type:int;size:10;"`                                // 购买人ID
	Status   int    `json:"status" form:"status" gorm:"column:status;comment:状态 1-确认手续费 2-出售中 3-支付确认中 4-已完成 5-已撤销;type:int;size:10;"` // 状态
	CardName string `json:"cardName" form:"cardName" gorm:"column:card_name;comment:卡牌名称;type:varchar(255);size:255;"`                // 卡牌名称
	Level    int    `json:"level" form:"level" gorm:"column:level;comment:卡牌等级;type:int;size:10;"`                                    // 卡牌等级
	TxHash   string `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:交易hash;type:varchar(255);size:255;"`                    // 交易hash
	From     string `json:"from" form:"from" gorm:"column:from;comment:出售人钱包地址;type:varchar(255);size:255;"`                          // 出售人钱包地址
	To       string `json:"to" form:"to" gorm:"column:to;comment:购买人钱包地址;type:varchar(255);size:255;"`                                // 购买人钱包地址
	Block    string `json:"block" form:"block" gorm:"column:block;comment:区块号;type:varchar(255);size:255;"`                           // 区块号
	System   string `json:"system" form:"system" gorm:"column:system;comment:手续费钱包地址;type:varchar(255);size:255;"`                    // 系统手续费钱包地址
}

func (AvfCardTransfer) TableName() string {
	return "avf_card_transfer"
}

func (h *AvfCardTransfer) Update(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id = ?", h.ID).Updates(&h).Error
}

func (h *AvfCardTransfer) Create(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}

func (h *AvfCardTransfer) GetById(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id =?", h.ID).First(&h).Error
}

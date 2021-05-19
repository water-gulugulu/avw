// 自动生成模板AvfCardTransfer
package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type AvfCardTransfer struct {
	global.GVA_MODEL
	RecordId   int     `json:"record_id" form:"record_id" gorm:"column:record_id;comment:卡牌记录ID;type:int;size:10;"`                                     // 卡牌记录ID
	Uid        int     `json:"uid" form:"uid" gorm:"column:uid;comment:出售人ID;type:int;size:10;"`                                                        // 出售人ID
	CardId     int     `json:"card_id" form:"card_id" gorm:"column:card_id;comment:卡牌ID;type:int;size:10;"`                                             // 卡牌ID
	Price      int     `json:"price" form:"price" gorm:"column:price;comment:卡牌出售价格;type:int;size:10;"`                                                 // 价格
	Fees       int     `json:"fees" form:"fees" gorm:"column:fees;comment:手续费;type:int;size:10;"`                                                       // 手续费
	BuyId      int     `json:"buy_id" form:"buy_id" gorm:"column:buy_id;comment:购买人ID;type:int;size:10;"`                                               // 购买人ID
	Status     int     `json:"status" form:"status" gorm:"column:status;comment:状态 1-手续费待支付 2-确认手续费 3-出售中 4-待支付 5-支付确认中 6-已完成 7-已撤销;type:int;size:10;"` // 状态 1-手续费待支付 2-确认手续费 3-出售中 4-待支付 5-支付确认中 6-已完成 7-已撤销
	CardName   string  `json:"cardName" form:"cardName" gorm:"column:card_name;comment:卡牌名称;type:varchar(255);size:255;"`                               // 卡牌名称
	Level      int     `json:"level" form:"level" gorm:"column:level;comment:卡牌等级;type:int;size:10;"`                                                   // 卡牌等级
	TxHash     string  `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:交易hash;type:varchar(255);size:255;"`                                   // 交易hash
	From       string  `json:"from" form:"from" gorm:"column:from;comment:出售人钱包地址;type:varchar(255);size:255;"`                                         // 出售人钱包地址
	To         string  `json:"to" form:"to" gorm:"column:to;comment:购买人钱包地址;type:varchar(255);size:255;"`                                               // 购买人钱包地址
	Block      string  `json:"block" form:"block" gorm:"column:block;comment:区块号;type:varchar(255);size:255;"`                                          // 区块号
	System     string  `json:"system" form:"system" gorm:"column:system;comment:手续费钱包地址;type:varchar(255);size:255;"`                                   // 系统手续费钱包地址
	FeesHash   string  `json:"fees_hash" form:"fees_hash" gorm:"column:fees_hash;comment:手续费交易hash;type:varchar(255);size:255;"`                        // 手续费hash
	Card       AvfCard `gorm:"ForeignKey:CardId;References:ID"`                                                                                         // 卡牌
	ExpireTime int     `json:"expire_time" form:"expire_time" gorm:"column:expire_time;comment:过期时间;type:int;size:10;"`                                 // 支付过期时间
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
func (h *AvfCardTransfer) GetById2(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id =?", h.ID).Preload("Card").First(&h).Error
}

func (h *AvfCardTransfer) GetByHash(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("tx_hash = ?", h.TxHash).First(&h).Error
}

func (h *AvfCardTransfer) GetByFeesHash(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("fees_hash = ?", h.FeesHash).First(&h).Error
}

func (h *AvfCardTransfer) GetByCardIdAndUserIdAndNotCancel(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("card_id = ? and uid = ? and status != ?", h.CardId, h.Uid, h.Status).Order("id desc").First(&h).Error
}

func (h *AvfCardTransfer) GetByStatus(DB *gorm.DB) (list []*AvfCardTransfer, err error) {
	if err = DB.Table(h.TableName()).Where("status = ?", h.Status).Find(&list).Error; err != nil {
		return nil, err
	}

	return
}
func (h *AvfCardTransfer) GetList(DB *gorm.DB, page, size int) (list []*AvfCardTransfer, total int64, err error) {
	DB = DB.Table(h.TableName())
	if h.Status != 0 {
		DB = DB.Where("status = ?", h.Status)
	}
	if h.Level != 0 {
		DB = DB.Where("level = ?", h.Level)
	}
	if page != 0 {
		page = page * size
	}
	if err = DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = DB.Order("id desc").Limit(size).Offset(page).Preload("Card").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return
}

func (h *AvfCardTransfer) GetListByBuyId(DB *gorm.DB, page, size int) (list []*AvfCardTransfer, total int64, err error) {
	DB = DB.Table(h.TableName())
	if h.Level != 0 {
		DB = DB.Where("level = ?", h.Level)
	}
	if h.Level != 0 {
		DB = DB.Where("status = ?", h.Status)
	}
	if page != 0 {
		page = page * size
	}
	if err = DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = DB.Order("id desc").Limit(size).Offset(page).Preload("Card").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return
}

// 自动生成模板AvfOrder
package model

import (
	"errors"
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type AvfOrder struct {
	global.GVA_MODEL
	Uid      int    `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;type:int;size:10;"`                             // 用户ID
	OrderSn  string `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:订单编号;type:varchar(255);size:255;"`      // 订单编号
	Price    int64  `json:"price" form:"price" gorm:"column:price;comment:支付价格;type:int;size:10;"`                       // 价格
	Num      int    `json:"num" form:"num" gorm:"column:num;comment:购买数量;type:int;size:10;"`                             // 数量
	Number   int    `json:"number" form:"number" gorm:"force;column:number;comment:剩余数量;type:int;size:10;"`              // 剩余数量
	Status   int    `json:"status" form:"status" gorm:"column:status;comment:状态 1-待支付 2-待确认 3-已完成 4-已取消5-已关闭;type:int;"` // 状态 1-待支付 2-待确认 3-已完成 4-已取消5-已关闭
	PayTime  int    `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;type:int;size:10;"`                // 支付时间
	TxHash   string `json:"txHash" form:"txHash" gorm:"column:tx_hash;comment:事务哈希;type:varchar(255);size:255;"`         // 交易hash
	Block    string `json:"block" form:"block" gorm:"column:block;comment:区块编号;type:varchar(30);size:30;"`               // 区块号
	Gas      string `json:"gas" form:"gas" gorm:"column:gas;comment:手续费;type:varchar(50);size:50;"`                      // 手续费
	GasPrice string `json:"gasPrice" form:"gasPrice" gorm:"column:gas_price;comment:手续费价格;type:varchar(50);size:50;"`    // 手续费价格
	From     string `json:"from" form:"from" gorm:"column:from;comment:支付地址;type:varchar(255);size:255;"`                // 支付地址
	To       string `json:"to" form:"to" gorm:"column:to;comment:收款地址;type:varchar(255);size:255;"`                      // 收款地址
}

func (h *AvfOrder) TableName() string {
	return "avf_order"
}

func (h *AvfOrder) FindList(DB *gorm.DB, p, size int) (list []AvfOrder, total int64, err error) {
	DB = DB.Table(h.TableName()).Where("uid = ? AND status in(2,3)", h.Uid)

	if p != 0 {
		p = p * size
	}
	if err := DB.Count(&total).Error; err != nil {
		return list, 0, err
	}
	if err := DB.Order("id desc").Limit(size).Offset(p).Find(&list).Error; err != nil {
		return list, 0, err
	}
	return
}

func (h *AvfOrder) CreateOrder(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Create(&h).Error
}

func (h *AvfOrder) FindById(DB *gorm.DB) error {
	if h.ID == 0 {
		return errors.New("ID不能为空")
	}

	return DB.Table(h.TableName()).Where("id = ?", h.ID).First(&h).Error
}
func (h *AvfOrder) FindByIdAndUserId(DB *gorm.DB) error {
	if h.ID == 0 {
		return errors.New("ID不能为空")
	}
	if h.Uid == 0 {
		return errors.New("用户ID不能为空")
	}

	return DB.Table(h.TableName()).Where("id = ? and uid = ?", h.ID, h.Uid).First(&h).Error
}

func (h *AvfOrder) UpdateOrder(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("id = ?", h.ID).Updates(&h).Error
}

func (h *AvfOrder) UpdateOrderByNumber(DB *gorm.DB) error {
	update := make(map[string]interface{}, 0)
	update["number"] = h.Number
	update["updated_at"] = h.UpdatedAt

	return DB.Table(h.TableName()).Where("id = ?", h.ID).Updates(update).Error
}

func (h *AvfOrder) FindListByStatus(DB *gorm.DB) (list []*AvfOrder, err error) {
	if h.Status == 0 {
		return nil, errors.New("状态不能为空")
	}

	if err = DB.Table(h.TableName()).Where("status = ?", h.Status).Find(&list).Error; err != nil {
		return nil, err
	}

	return
}

func (h *AvfOrder) FindByHash(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("tx_hash = ?", h.TxHash).First(&h).Error
}

func (h *AvfOrder) GetByUid(DB *gorm.DB) error {
	return DB.Table(h.TableName()).Where("uid = ?", h.Uid).First(&h).Error
}

func (h *AvfOrder) GetListByUid(DB *gorm.DB) (list []*AvfOrder, err error) {
	if err = DB.Table(h.TableName()).Where("uid = ? and status = ?", h.Uid, 3).Find(&list).Error; err != nil {
		return nil, err
	}
	return
}

func (h *AvfOrder) GetListAll(DB *gorm.DB) (list []*AvfOrder, err error) {

	if err = DB.Table(h.TableName()).Order("id asc").Find(&list).Error; err != nil {
		return nil, err
	}
	return
}

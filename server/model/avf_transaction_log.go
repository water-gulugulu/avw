// @File  : avf_transaction_log.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/16
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

type AvfTransactionLog struct {
	ID         int64     `gorm:"column:id" json:"id" form:"id"`                            // ID
	OrderId    int       `gorm:"column:order_id" json:"order_id" form:"order_id"`          // 订单ID
	Block      string    `gorm:"column:block" json:"block" form:"block"`                   // 区块号
	TxHash     string    `gorm:"column:tx_hash" json:"tx_hash" form:"tx_hash"`             // 事务hash
	Form       string    `gorm:"column:form" json:"form" form:"form"`                      // 发起地址
	To         string    `gorm:"column:to" json:"to" form:"to"`                            // 接收地址
	GasPrice   string    `gorm:"column:gas_price" json:"gas_price" form:"gas_price"`       // 手续费价格
	Gas        string    `gorm:"column:gas" json:"gas" form:"gas"`                         // 手续费
	Value      string    `gorm:"column:value" json:"value" form:"value"`                   // value
	Nonce      string    `gorm:"column:nonce" json:"nonce" form:"nonce"`                   // nonce
	Data       string    `gorm:"column:data" json:"data" form:"data"`                      // data
	CheckNonce string    `gorm:"column:check_nonce" json:"check_nonce" form:"check_nonce"` // check_nonce
	Status     int64     `gorm:"column:status" json:"status" form:"status"`                // 状态 1-成功 2-失败
	CreateTime int64     `gorm:"column:create_time" json:"create_time" form:"create_time"` // 创建时间
	CreateDate time.Time `gorm:"column:create_date" json:"create_date" form:"create_date"` // 创建时间
}

func (h *AvfTransactionLog) TableName() string {
	return "avf_transaction_log"
}
func (h *AvfTransactionLog) CreateLog(DB *gorm.DB) {
	DB.Table(h.TableName()).Create(&h)
}

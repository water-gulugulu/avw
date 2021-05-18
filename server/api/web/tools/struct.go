// @File  : struct.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/15
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

package web_tools

import (
	"gin-vue-admin/model"
	"time"
)

// 订单列表返回
type OrderListResponse struct {
	List  []model.AvfOrder `json:"list"`  // 订单列表
	Total int64            `json:"total"` // 总条数
}

// 用户信息
type UserInfo struct {
	Id            uint    `json:"id"`             // 用户ID
	Pid           string  `json:"pid"`            // 上级地址
	Username      string  `json:"username"`       // 用户名
	Status        *bool   `json:"status"`         // 状态
	WalletAddress string  `json:"wallet_address"` // 钱包地址
	AVWBalance    float64 `json:"avw_balance"`    // avw余额
	HTBalance     float64 `json:"ht_balance"`     // ht余额
}

// 登录返回信息
type LoginResponse struct {
	Id            uint   `json:"id"`             // 用户ID
	Pid           string `json:"pid"`            // 上级地址
	Username      string `json:"username"`       // 用户名
	Status        *bool  `json:"status"`         // 状态
	WalletAddress string `json:"wallet_address"` // 钱包地址
	Token         string `json:"token"`          // token
}

// 卡牌列表返回
type CardListResponse struct {
	List  []model.AvfCard `json:"list"`  // 卡牌列表
	Total int64           `json:"total"` // 总条数
}

// 创建订单返回
type CreateOrderResponse struct {
	OrderSn string      `json:"order_sn"` // 订单号
	OrderId uint        `json:"order_id"` // 订单ID
	Price   interface{} `json:"price"`    // 订单价格
}

// 订单详情返回结果
type OrderDetailResponse struct {
	OrderInfo     model.AvfOrder `json:"order_info"`      // 订单详情
	OrderCardList []AvfOrderCard `json:"order_card_list"` // 卡牌列表
}
type AvfOrderCard struct {
	ID        uint      `json:"id"`         // 主键ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	Uid       int       `json:"uid"`        // 用户ID
	OrderId   int       `json:"orderId"`    // 订单ID
	CardId    int       `json:"cardId"`     // 卡牌ID
	Star      int       `json:"star"`       // 算力值
	Status    int       `json:"status"`     // 状态 1-正常 2-转让中
	GiveType  int       `json:"giveType"`   // 获得方式 1-抽奖 2-购买
	Level     int       `json:"level"`      // 等级 1-N 2-R 3-SR 4-SSR
	Image     string    `json:"image"`      // 卡牌图
}

// 我的卡牌列表返回信息
type MyCardResponse struct {
	List  []*model.AvfOrderCard `json:"list"`  // 卡牌列表
	Total int64                 `json:"total"` // 总数
}

// 转让卡牌返回
type TransferResponse struct {
	TransferId    int    `json:"transfer_id"`    // 转让卡牌ID
	Fees          int    `json:"fees"`           // 手续费
	Price         int    `json:"price"`          // 价格
	SystemAddress string `json:"system_address"` // 系统收款地址
}

// 我的卡牌详情返回信息
type MyCardDetailResponse struct {
	OrderCard       model.AvfOrderCard     `json:"order_card"`       // 订单卡牌记录
	Order           *model.AvfCardTransfer `json:"order"`            // 卡牌转让订单
	All             int                    `json:"all"`              // 全部收益
	Today           int                    `json:"today"`            // 今日收益
	Yesterday       int                    `json:"yesterday"`        // 昨日收益
	Fees            int                    `json:"fees"`             // 手续费
	Price           int                    `json:"price"`            // 最低价格
	PricePercentage int                    `json:"price_percentage"` // 手续费百分比
	FeesPercentage  int                    `json:"fees_percentage"`  // 最低价格百分比
}

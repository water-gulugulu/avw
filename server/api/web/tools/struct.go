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
	All             float64                `json:"all"`              // 全部收益
	Today           float64                `json:"today"`            // 今日收益
	Yesterday       float64                `json:"yesterday"`        // 昨日收益
	Fees            interface{}            `json:"fees"`             // 手续费
	Price           interface{}            `json:"price"`            // 最低价格
	PricePercentage int                    `json:"price_percentage"` // 手续费百分比
	FeesPercentage  int                    `json:"fees_percentage"`  // 最低价格百分比
}

// 卡牌市场返回信息
type CardMarketResponse struct {
	List  []*model.AvfCardTransfer `json:"list"`  // 出售卡牌列表
	Total int64                    `json:"total"` // 总数
}

// 卡牌市场详情
type CardMarketDetailResponse struct {
	CardId          int    `json:"card_id"`          // 卡牌ID
	Name            string `json:"name"`             // 卡牌名称
	ContractAddress string `json:"contract_address"` // 合约地址
	Author          string `json:"author"`           // 作者
	Desc            string `json:"desc"`             // 描述
	Star            int    `json:"star"`             // 算力值
	Image           string `json:"image"`            // 卡牌图片
	SellId          int    `json:"sell_id"`          // 出售人ID
	Price           int    `json:"price"`            // 出售价格
	Fees            int    `json:"fees"`             // 手续费
	Status          int    `json:"status"`           // 出售状态  1-手续费待支付 2-确认手续费 3-出售中 4-待支付 5-支付确认中 6-已完成 7-已撤销
	Level           int    `json:"level"`            // 卡牌等级
	OriginalPrice   int    `json:"original_price"`   // 卡牌原价
	From            string `json:"from"`             // 出售人钱包地址
	ExpireTime      int    `json:"expire_time"`      // 过期时间
}

// 我的团队返回信息
type MyTeamResponse struct {
	List       []AvfUser `json:"list"`        // 我的直推下级
	TeamCount  int       `json:"team_count"`  // 我的团队
	LowerCount int       `json:"lower_count"` // 我的下级
}

type AvfUser struct {
	Id            int       `json:"id"`             // 用户ID
	Pid           string    `json:"pid"`            // 上级地址
	Username      string    `json:"username"`       // 用户名
	WalletAddress string    `json:"wallet_address"` // 钱包地址
	CreatedAt     time.Time `json:"created_at"`     // 创建时间
	IsNumber      bool      `json:"is_number"`      // 是否正式会员
}

// 用户账单返回
type UserBillResponse struct {
	List  []*model.AvfUserBill `json:"list"`  // 账单列表
	Total int64                `json:"total"` // 总条数
}

// 用户账单返回
type MiningRecordResponse struct {
	List  []*AvfUserBill `json:"list"`  // 账单列表
	Total int64          `json:"total"` // 总条数
}

// 如果含有time.Time 请自行import time包
type AvfUserBill struct {
	ID         uint      `json:"id"`                                                                                               // id
	CreatedAt  time.Time `json:"created_at"`                                                                                       // 格式化创建时间
	UpdatedAt  time.Time `json:"updated_at"`                                                                                       // 格式化修改时间
	Uid        int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;type:int;size:10;"`                                  // 用户ID
	CardId     int       `json:"cardId" form:"cardId" gorm:"column:card_id;comment:卡牌ID;type:int;size:10;"`                        // 卡牌ID
	Address    string    `json:"address" form:"address" gorm:"column:address;comment:钱包地址;type:varchar(255);size:255;"`            // 钱包地址
	Type       int       `json:"type" form:"type" gorm:"column:type;comment:类型 1-发放收益 2-盲盒 3-购买卡牌 4-手续费 5-直推收益;type:int;size:10;"` // 类型 1-发放收益 2-盲盒 3-购买卡牌 4-手续费 5-直推收益
	Money      float64   `json:"money" form:"money" gorm:"column:money;comment:金额;type:decimal;size:9,4;"`                         // 金额
	Fees       float64   `json:"fees" form:"fees" gorm:"column:fees;comment:手续费;type:decimal;size:9,4;"`                           // 手续费
	Balance    float64   `json:"balance" form:"balance" gorm:"column:balance;comment:余额;type:decimal;size:9,4;"`                   // 余额
	Payment    int       `json:"payment" form:"payment" gorm:"column:payment;comment:收入支出 1-收入 2-支出;type:int;size:10;"`            // 收入支出
	PayType    int       `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付方式 1-avw 2-ht;type:int;size:10;"`          // 支付方式 1-avw 2-ht
	Detail     string    `json:"detail" form:"detail" gorm:"column:detail;comment:描述;type:varchar(255);size:255;"`                 // 详情
	TxHash     string    `json:"tx_hash" form:"tx_hash" gorm:"column:tx_hash;comment:交易hash;type:varchar(255);size:255;"`          // 交易hash
	CreateTime int       `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:int;size:10;"`            // 创建时间
	Star       int       `json:"star"`                                                                                             // 算力
	Level      int       `json:"level"`                                                                                            // 卡牌等级 1-N 2-R 3-SR 4-SSR
}

// 我的统计返回信息
type MyStatisticalResponse struct {
	AllForce          int     `json:"all_force"`          // 总算力
	AllEarnings       float64 `json:"all_earnings"`       // 总收益
	TodayEarnings     float64 `json:"today_earnings"`     // 昨日收益
	YesterdayEarnings float64 `json:"yesterday_earnings"` // 今日收益
}

// 开源统计返回信息
type OpenStatisticalResponse struct {
	RegUser        int `json:"reg_user"`        // 注册用户
	ActivationUser int `json:"activation_user"` // 激活用户
	Trading        int `json:"trading"`         // 	交易总额
	AllDayTrading  int `json:"all_day_trading"` // 24小时交易总额
	Input          int `json:"inflows"`         // 流入
	Output         int `json:"out_of"`          // 流出
}

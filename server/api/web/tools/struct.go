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

import "gin-vue-admin/model"

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
	OrderSn string `json:"order_sn"` // 订单号
	OrderId uint   `json:"order_id"` // 订单ID
	Price   int64  `json:"price"`    // 订单价格
}

// 订单详情返回结果
type OrderDetailResponse struct {
	OrderInfo     model.AvfOrder       `json:"order_info"`      // 订单详情
	OrderCardList []model.AvfOrderCard `json:"order_card_list"` // 卡牌列表
}

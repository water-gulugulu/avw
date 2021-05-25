// @File  : order.go
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

package web

import (
	"fmt"
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/api/web/tools/response"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// @Tags 前端接口
// @Summary 订单列表
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param page query string  false "页码 第一页为0"
// @Param size query string  false "数量 默认10"
// @Success 200 {object} web_tools.OrderListResponse
// @Router /web/order/list [get]
func GetOrderList(c *gin.Context) {
	UserId, e := web_tools.GetUserId(c)
	if e != nil {
		response.FailWithMessage("41003", c)
		return
	}
	page := c.Query("page")
	size := c.Query("size")

	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		page = "0"
	}

	Order := model.AvfOrder{
		Uid: int(UserId),
	}
	p, _ := strconv.Atoi(page)
	s, _ := strconv.Atoi(size)
	res := web_tools.OrderListResponse{}
	list, total, err := Order.FindList(global.GVA_DB, p, s)
	if err != nil {
		res.List = make([]model.AvfOrder, 0)
		response.OkWithData(res, c)
		return
	}
	res = web_tools.OrderListResponse{
		List:  list,
		Total: total,
	}
	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 创建订单
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param type body string  true "盲盒类型 1-一连抽 2-十连抽"
// @Param num body string  true "数量 购买盲盒数量"
// @Success 200 {object}  web_tools.CreateOrderResponse
// @Router /web/order/createOrder [post]
func CreateOrder(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	blindBoxType := c.PostForm("type")
	num := c.PostForm("num")
	if len(num) == 0 {
		response.FailWithMessage("41008", c)
		return
	}
	blindBox := global.GVA_CONFIG.BlindBox
	var price int64
	var n int = 1
	number, _ := strconv.Atoi(num)
	switch blindBoxType {
	case "1":
		one, _ := strconv.Atoi(blindBox.One)
		price = int64(one * number)
		n = number
	case "2":
		ten, _ := strconv.Atoi(blindBox.One)
		price = int64(ten * number)
		n = number * 10
	default:
		response.FailWithMessage("41007", c)
		return
	}
	orderSn := web_tools.CreateSn(time.Now())
	Order := model.AvfOrder{
		Uid:       int(UserId),
		OrderSn:   orderSn,
		Num:       n,
		Number:    n,
		Price:     price,
		Status:    1,
		GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	if err := Order.CreateOrder(global.GVA_DB); err != nil {
		response.FailWithMessage("60000", c)
		return
	}

	res := web_tools.CreateOrderResponse{
		OrderSn: orderSn,
		OrderId: Order.ID,
		Price:   price,
		Address: global.GVA_CONFIG.CollectionAddress.Address,
	}

	if global.GVA_CONFIG.CollectionAddress.Debug == "1" {
		res.Price = 0.0001
	}
	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 支付订单
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param order_id body string  true "订单ID"
// @Param tx_hash body string  true "发起交易后的事务hash"
// @Param address body string  true "发起交易的钱包地址"
// @Success 200 {string} string "{"status":200,"message":"200"}"
// @Router /web/order/payOrder [post]
func PayOrder(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	orderID := c.PostForm("order_id")
	TxHash := c.PostForm("tx_hash")
	address := c.PostForm("address")
	if len(orderID) == 0 {
		response.FailWithMessage("41009", c)
		return
	}
	if len(TxHash) == 0 {
		response.FailWithMessage("41010", c)
		return
	}
	if len(address) == 0 {
		response.FailWithMessage("41011", c)
		return
	}
	oid, _ := strconv.Atoi(orderID)
	Order := model.AvfOrder{
		TxHash: TxHash,
	}

	if err := Order.FindByHash(global.GVA_DB); err == nil || Order.ID != 0 {
		response.FailWithMessage("41013", c)
		return
	}
	Log := model.AvfTransactionLog{
		TxHash: TxHash,
	}
	if err := Log.GetByHash(global.GVA_DB); err == nil || Log.ID != 0 {
		response.FailWithMessage("41013", c)
		return
	}

	Order = model.AvfOrder{
		GVA_MODEL: global.GVA_MODEL{ID: uint(oid)},
		Uid:       int(UserId),
	}
	if err := Order.FindByIdAndUserId(global.GVA_DB); err != nil {
		response.FailWithMessage("60001", c)
		return
	}
	if Order.Status != 1 {
		response.FailWithMessage("60002", c)
		return
	}

	Order = model.AvfOrder{
		GVA_MODEL: global.GVA_MODEL{
			ID:        uint(oid),
			UpdatedAt: time.Now(),
		},
		Uid:    int(UserId),
		TxHash: TxHash,
		Status: 2,
		From:   address,
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := Order.UpdateOrder(tx); err != nil {
			return err
		}
		UserBill := model.AvfUserBill{
			GVA_MODEL:  global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Uid:        int(UserId),
			Address:    address,
			Type:       2,
			Money:      float64(Order.Price),
			Payment:    2,
			PayType:    2,
			Detail:     fmt.Sprintf("购买卡牌盲盒支付金额:%v", Order.Price),
			CreateTime: int(time.Now().Unix()),
		}
		if err := UserBill.Create(tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		response.FailWithMessage("60003", c)
		return
	}

	response.OkWithMessage("200", c)
	return
}

// @Tags 前端接口
// @Summary 订单详情
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param order_id query string  true "订单ID"
// @Success 200 {object} web_tools.OrderDetailResponse
// @Router /web/order/orderDetail [get]
func OrderDetail(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	orderID := c.Query("order_id")
	if len(orderID) == 0 {
		response.FailWithMessage("41009", c)
		return
	}

	oid, _ := strconv.Atoi(orderID)
	Order := model.AvfOrder{
		GVA_MODEL: global.GVA_MODEL{ID: uint(oid)},
		Uid:       int(UserId),
	}
	if err := Order.FindByIdAndUserId(global.GVA_DB); err != nil {
		response.FailWithMessage("60001", c)
		return
	}
	OrderCard := model.AvfOrderCard{
		OrderId: oid,
		Uid:     int(UserId),
	}
	list, _ := OrderCard.FindListByOrderId(global.GVA_DB)

	l := make([]web_tools.AvfOrderCard, 0)
	for _, item := range list {
		l = append(l, web_tools.AvfOrderCard{
			ID:        item.ID,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.CreatedAt,
			Uid:       item.Uid,
			OrderId:   item.OrderId,
			CardId:    item.CardId,
			Star:      item.Star,
			Status:    item.Status,
			GiveType:  item.GiveType,
			Level:     item.Level,
			Image:     item.Card.Image,
		})
	}

	res := web_tools.OrderDetailResponse{
		OrderInfo:     Order,
		OrderCardList: l,
	}

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 取消订单
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param order_id body string  true "订单ID"
// @Success 200 {string} string {"status":200,"msg":"成功"}
// @Router /web/order/cancelOrder [post]
func CancelOrder(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	orderID := c.PostForm("order_id")
	if len(orderID) == 0 {
		response.FailWithMessage("41009", c)
		return
	}

	oid, _ := strconv.Atoi(orderID)
	Order := model.AvfOrder{
		GVA_MODEL: global.GVA_MODEL{ID: uint(oid)},
		Uid:       int(UserId),
	}
	if err := Order.FindByIdAndUserId(global.GVA_DB); err != nil {
		response.FailWithMessage("60001", c)
		return
	}

	Order.Status = 4
	if err := Order.UpdateOrder(global.GVA_DB); err != nil {
		response.FailWithMessage("60001", c)
		return
	}

	response.Ok(c)
	return
}

// @Tags 前端接口
// @Summary 获取盲盒价格
// @accept application/json
// @Produce application/json
// @Success 200 {object} config.BlindBox
// @Router /web/order/getPrice [get]
func GetPrice(c *gin.Context) {
	response.OkWithData(global.GVA_CONFIG.BlindBox, c)
	return
}

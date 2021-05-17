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
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/api/web/tools/response"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"github.com/gin-gonic/gin"
	"log"
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
	number, _ := strconv.Atoi(num)
	switch blindBoxType {
	case "1":
		price = blindBox.One * int64(number)
	case "2":
		price = blindBox.Ten * int64(number)
	default:
		response.FailWithMessage("41007", c)
		return
	}
	orderSn := web_tools.CreateSn(time.Now())
	Order := model.AvfOrder{
		Uid:       int(UserId),
		OrderSn:   orderSn,
		Num:       number,
		Number:    number,
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
	if err := Order.UpdateOrder(global.GVA_DB); err != nil {
		response.FailWithMessage("60003", c)
		return
	}
	go LoopOrderStatus(TxHash, oid)

	response.OkWithMessage("200", c)
	return
}

// @Tags 前端接口
// @Summary 订单详情
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param order_id body string  true "订单ID"
// @Success 200 {object} web_tools.OrderDetailResponse
// @Router /web/order/orderDetail [get]
func OrderDetail(c *gin.Context) {
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
	OrderCard := model.AvfOrderCard{
		OrderId: oid,
	}
	list, _ := OrderCard.FindListByOrderId(global.GVA_DB)

	res := web_tools.OrderDetailResponse{
		OrderInfo:     Order,
		OrderCardList: list,
	}

	response.OkWithData(res, c)
	return
}

// 循环读取哈希来改变订单状态
func LoopOrderStatus(txHash string, OrderId int) {
	if len(txHash) == 0 {
		return
	}
	client, err := blockchian.NewClient()

	if err != nil {
		log.Printf("[%s]Failed to client RPC by Hash:%s error:%e", time.Now(), txHash, err)
		return
	}
	defer client.CloseClient()
	Order := model.AvfOrder{
		TxHash: txHash,
	}
	if err := Order.FindByHash(global.GVA_DB); err != nil {
		log.Printf("[%s]Failed to Hash:%s query Order error:%e", time.Now(), txHash, err)
		return
	}

	for {
		res, err2 := client.QueryTransactionByTxHash(txHash)
		if err2 != nil {
			log.Printf("[%s]Failed to query transaction error:%e", time.Now(), err)
			continue
		}
		if res.Status != 1 {
			log.Printf("[%s]Failed to status not 1", time.Now())
			continue
		}
		if res.From != Order.From {
			log.Printf("[%s]Failed to form no ok", time.Now())
			continue
		}
		if res.To != global.GVA_CONFIG.CollectionAddress.Address {
			log.Printf("[%s]Failed to to no ok", time.Now())
			continue
		}
		Order := model.AvfOrder{
			GVA_MODEL: global.GVA_MODEL{
				ID:        uint(OrderId),
				UpdatedAt: time.Now(),
			},
			TxHash:   txHash,
			Status:   3,
			PayTime:  int(time.Now().Unix()),
			Block:    res.Block.String(),
			Gas:      string(res.Gas),
			GasPrice: res.GasPrice.String(),
			From:     res.From,
			To:       res.To,
		}
		if err := Order.UpdateOrder(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to update Order error:%e", time.Now(), err)
			continue
		}

		Log := model.AvfTransactionLog{
			OrderId:  OrderId,
			Block:    res.Block.String(),
			TxHash:   txHash,
			Form:     res.From,
			To:       res.To,
			Gas:      strconv.Itoa(int(res.Gas)),
			GasPrice: res.GasPrice.String(),
			Value:    res.Value.String(),
			Nonce:    string(res.Nonce),
			Data:     string(res.Data),
			Status:   int64(res.Status),
		}
		Log.CreateLog(global.GVA_DB)
		break
	}
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

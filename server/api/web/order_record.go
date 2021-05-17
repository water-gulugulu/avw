// @File  : order_record.go
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

package web

import (
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
// @Summary 抽奖
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param order_id body string  true "订单ID"
// @Success 200 {object} model.AvfCard
// @Router /web/order_card/luckyDraw [post]
func LuckyDraw(c *gin.Context) {
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
	DB := global.GVA_DB

	oid, _ := strconv.Atoi(orderID)
	Order := model.AvfOrder{
		GVA_MODEL: global.GVA_MODEL{ID: uint(oid)},
		Uid:       int(UserId),
	}
	if err := Order.FindByIdAndUserId(DB); err != nil {
		response.FailWithMessage("60001", c)
		return
	}
	if Order.Status != 3 {
		response.FailWithMessage("60002", c)
		return
	}

	if Order.Number < 1 {
		response.FailWithMessage("60004", c)
		return
	}
	level := web_tools.Lottery(false)

	Card := model.AvfCard{
		Level: level,
	}
	if err := Card.RandGetByLevel(DB); err != nil {
		response.FailWithMessage("60005", c)
		return
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		Order.Number = Order.Number - 1
		Order.UpdatedAt = time.Now()

		if err := Order.UpdateOrder(tx); err != nil {
			return err
		}

		OrderCard := model.AvfOrderCard{
			OrderId:    oid,
			Uid:        int(UserId),
			CardId:     int(Card.ID),
			Level:      Card.Level,
			Star:       Card.Star,
			Status:     1,
			CreateTime: int(time.Now().Unix()),
			UpdateTime: int(time.Now().Unix()),
			GiveType:   1,
		}

		if err := OrderCard.CreateOrderCard(tx); err != nil {
			return err
		}

		// fmt.Printf("%s", OrderCard)
		return nil
	})
	if err != nil {
		response.FailWithMessage("60004", c)
		return
	}

	response.OkWithData(Card, c)
	return
}

// @Tags 前端接口
// @Summary 我的卡牌列表
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param page query string  false "页码"
// @Param size query string  false "数量 默认10"
// @Success 200 {object} web_tools.MyCardResponse
// @Router /web/order_card/myCard [get]
func MyCard(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
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

	p, _ := strconv.Atoi(page)
	s, _ := strconv.Atoi(size)

	OrderCard := model.AvfOrderCard{
		Uid: int(UserId),
	}
	res := web_tools.MyCardResponse{}
	list, total, err2 := OrderCard.FindListByUid(global.GVA_DB, p, s)
	if err2 != nil {
		res.List = make([]*model.AvfOrderCard, 0)
		response.OkWithData(res, c)
		return
	}
	res = web_tools.MyCardResponse{
		List:  list,
		Total: total,
	}

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 卡牌转让
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param record_id body string  true "卡牌记录ID"
// @Param price body string  true "转让价格"
// @Success 200 {object} web_tools.TransferResponse
// @Router /web/order_card/transferCard [post]
func TransferCard(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	recordId := c.PostForm("record_id")
	price := c.PostForm("price")
	if len(recordId) == 0 || recordId == "0" {
		response.FailWithMessage("41014", c)
		return
	}
	if len(price) == 0 || price == "0" {
		response.FailWithMessage("41017", c)
		return
	}
	rid, _ := strconv.Atoi(recordId)
	cardPrice, _ := strconv.Atoi(price)
	orderCard := model.AvfOrderCard{
		GVA_MODEL: global.GVA_MODEL{ID: uint(rid)},
	}
	DB := global.GVA_DB
	if err := orderCard.GetById(DB); err != nil {
		response.FailWithMessage("60006", c)
		return
	}
	if orderCard.Uid != int(UserId) {
		response.FailWithMessage("41015", c)
		return
	}
	if orderCard.Status == 2 {
		response.FailWithMessage("41016", c)
		return
	}
	proportion, _ := strconv.Atoi(global.GVA_CONFIG.CollectionAddress.Proportion)

	systemPrice := int(orderCard.Card.Money) * proportion / 100
	if cardPrice > systemPrice {
		response.FailWithMessageToSprintf("41018", c, proportion)
		return
	}
	Fees, _ := strconv.Atoi(global.GVA_CONFIG.CollectionAddress.Fees)
	fees := cardPrice * Fees / 100

	systemAddress := global.GVA_CONFIG.CollectionAddress.Address
	cardTransfer := model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		RecordId:  rid,
		Uid:       orderCard.Uid,
		CardId:    orderCard.CardId,
		Price:     cardPrice,
		Fees:      fees,
		BuyId:     0,
		Status:    1,
		CardName:  orderCard.Card.Name,
		Level:     orderCard.Level,
		TxHash:    "",
		From:      "",
		To:        "",
		Block:     "",
		System:    systemAddress,
	}

	orderCard = model.AvfOrderCard{
		GVA_MODEL: global.GVA_MODEL{
			ID:        uint(rid),
			UpdatedAt: time.Now(),
		},
		Status:     2,
		UpdateTime: int(time.Now().Unix()),
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := orderCard.Update(tx); err != nil {
			return err
		}
		if err := cardTransfer.Create(tx); err != nil {
			return err
		}

		return nil
	})
	res := web_tools.TransferResponse{
		TransferId:    int(cardTransfer.ID),
		Fees:          fees,
		Price:         cardPrice,
		SystemAddress: systemAddress,
	}

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 支付转让手续费
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param transfer_id body string  true "卡牌转让ID"
// @Param tx_hash body string  true "交易事务hash"
// @Param address body string  true "提交支付钱包地址"
// @Success 200 {object} web_tools.TransferResponse
// @Router /web/order_card/payFees [post]
func PayFees(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	transferId := c.PostForm("transfer_id")
	TxHash := c.PostForm("tx_hash")
	Address := c.PostForm("address")
	if len(transferId) == 0 || transferId == "0" {
		response.FailWithMessage("41019", c)
		return
	}
	if len(TxHash) == 0 {
		response.FailWithMessage("41010", c)
		return
	}
	if len(Address) == 0 {
		response.FailWithMessage("41011", c)
		return
	}

	tid, _ := strconv.Atoi(transferId)
	cardTransfer := model.AvfCardTransfer{
		TxHash: TxHash,
	}
	DB := global.GVA_DB

	if err := cardTransfer.GetByHash(DB); err == nil || cardTransfer.ID != 0 {
		response.FailWithMessage("41013", c)
		return
	}
	Log := model.AvfTransactionLog{
		TxHash: TxHash,
	}
	if err := Log.GetByHash(DB); err == nil || Log.ID != 0 {
		response.FailWithMessage("41013", c)
		return
	}

	cardTransfer = model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(tid)},
	}
	if err := cardTransfer.GetById(DB); err != nil {
		response.FailWithMessage("60006", c)
		return
	}

	if cardTransfer.Uid != int(UserId) {
		response.FailWithMessage("41020", c)
		return
	}
	if cardTransfer.Status != 1 {
		response.FailWithMessage("60007", c)
		return
	}
}

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
	"fmt"
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/api/web/tools/response"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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

		if err := Order.UpdateOrderByNumber(tx); err != nil {
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
		fmt.Printf("%s", err)
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
// @Summary 我的卡牌详情
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param record_id query string  true "卡牌记录ID"
// @Success 200 {object} web_tools.MyCardDetailResponse
// @Router /web/order_card/myCardDetail [get]
func MyCardDetail(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	orderId := c.Query("record_id")

	if len(orderId) == 0 {
		response.FailWithMessage("41014", c)
		return
	}

	oid, _ := strconv.Atoi(orderId)

	OrderCard := model.AvfOrderCard{
		GVA_MODEL: global.GVA_MODEL{ID: uint(oid)},
		Uid:       int(UserId),
	}
	DB := global.GVA_DB
	if err2 := OrderCard.GetById(DB); err2 != nil {
		response.FailWithMessage("60001", c)
		return
	}
	if OrderCard.Uid != int(UserId) {
		response.FailWithMessage("41015", c)
		return
	}
	// 用户账单
	UserBill := model.AvfUserBill{
		Uid:    int(UserId),
		CardId: OrderCard.CardId,
		Type:   1,
	}
	var yesterday, today, all float64

	list, err3 := UserBill.GetByUidAndCardId(DB)
	if err3 != nil {
		yesterday, today, all = 0, 0, 0
	}
	todayTime := web_tools.GetTodayZeroTimeStamp()
	yesterdayTime := int(today) - 86400
	// 统计昨日，今日，全部收益
	for _, item := range list {
		all = all + item.Money
		if item.CreateTime > todayTime && item.CreateTime < todayTime+86399 {
			today = today + item.Money
		}
		if item.CreateTime > yesterdayTime && item.CreateTime < yesterdayTime+86399 {
			yesterday = yesterday + item.Money
		}
	}

	Fees := global.GVA_CONFIG.CollectionAddress.Fees
	Proportion := global.GVA_CONFIG.CollectionAddress.Proportion
	fees, _ := strconv.Atoi(Fees)
	proportion, _ := strconv.Atoi(Proportion)

	res := web_tools.MyCardDetailResponse{
		OrderCard: OrderCard,
		All:       all,
		Today:     today,
		Yesterday: yesterday,
	}

	Order := model.AvfCardTransfer{
		CardId: OrderCard.CardId,
		Uid:    int(UserId),
		Status: 7,
	}
	if err := Order.GetByCardIdAndUserIdAndNotCancel(DB); err != nil {
		fmt.Printf("err:%s", err)
		res.Order = nil
	} else {
		res.Order = &Order
	}

	res.Fees = int(OrderCard.Card.Money) * fees / 100
	res.Price = int(OrderCard.Card.Money) * proportion / 100
	if global.GVA_CONFIG.CollectionAddress.Debug == "1" {
		res.Fees = 0.001
		res.Price = 0.001
	}
	res.FeesPercentage = fees
	res.PricePercentage = proportion
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
	if cardPrice < systemPrice {
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
		Status:    1,
		CardName:  orderCard.Card.Name,
		Level:     orderCard.Level,
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
			log.Printf("[%s]Failed to order_card update error:%e\n", time.Now(), err)
			return err
		}
		if err := cardTransfer.Create(tx); err != nil {
			log.Printf("[%s]Failed to create card_transfer error:%e\n", time.Now(), err)
			return err
		}

		return nil
	})
	if err != nil {
		response.FailWithMessage("60003", c)
		return
	}

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
// @Success 200 {string} string "{"code":0}"
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
		FeesHash: TxHash,
	}
	DB := global.GVA_DB

	if err := cardTransfer.GetByFeesHash(DB); err == nil || cardTransfer.ID != 0 {
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
	cardTransfer = model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{
			ID:        uint(tid),
			UpdatedAt: time.Now(),
		},
		FeesHash: TxHash,
		Status:   2,
		From:     Address,
		System:   global.GVA_CONFIG.CollectionAddress.Address,
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := cardTransfer.Update(tx); err != nil {
			return err
		}

		UserBill := model.AvfUserBill{
			GVA_MODEL:  global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Uid:        int(UserId),
			Address:    Address,
			Type:       4,
			Money:      float64(cardTransfer.Fees),
			Fees:       float64(cardTransfer.Fees),
			Payment:    2,
			PayType:    2,
			Detail:     fmt.Sprintf("转让卡牌支付手续费:%v", cardTransfer.Fees),
			CreateTime: int(time.Now().Unix()),
		}
		if err := UserBill.Create(tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		response.FailWithMessage("60007", c)
		return
	}
	response.Ok(c)
	return
}

// @Tags 前端接口
// @Summary 取消转让卡牌
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param transfer_id body string  true "卡牌转让ID"
// @Success 200 {string} string "{"code":0}"
// @Router /web/order_card/cancelTransfer [post]
func CancelTransfer(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	transferId := c.PostForm("transfer_id")

	if len(transferId) == 0 {
		response.FailWithMessage("41019", c)
		return
	}

	tid, _ := strconv.Atoi(transferId)

	CardTransfer := model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(tid)},
	}
	DB := global.GVA_DB
	if err2 := CardTransfer.GetById(DB); err2 != nil {
		response.FailWithMessage("60001", c)
		return
	}
	if CardTransfer.Uid != int(UserId) {
		response.FailWithMessage("41015", c)
		return
	}
	if CardTransfer.Status > 3 {
		response.FailWithMessage("60002", c)
		return
	}

	OrderCard := model.AvfOrderCard{
		GVA_MODEL: global.GVA_MODEL{ID: uint(CardTransfer.RecordId), UpdatedAt: time.Now()},
	}
	if err := OrderCard.GetById(DB); err != nil {
		response.FailWithMessage("60001", c)
		return
	}
	OrderCard = model.AvfOrderCard{
		GVA_MODEL: global.GVA_MODEL{ID: uint(CardTransfer.RecordId), UpdatedAt: time.Now()},
		Status:    1,
	}
	CardTransfer = model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(tid)},
		Status:    7,
	}
	err = DB.Transaction(func(tx *gorm.DB) error {

		if err := OrderCard.Update(tx); err != nil {
			return err
		}

		if err := CardTransfer.Update(tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		response.FailWithMessage("60003", c)
		return
	}

	response.Ok(c)
	return
}

// 挖矿
func Mining(c *gin.Context) {
	DB := global.GVA_DB

	client, err := blockchian.NewClient()
	if err != nil {
		log.Printf("[%s]client blockchian failed error:%e\n", time.Now(), err)
		return
	}
	CardRecord := model.AvfOrderCard{
		Status: 1,
	}
	list, err := CardRecord.GetListByMining(DB)
	if err != nil {
		log.Printf("[%s]query card list failed error:%e\n", time.Now(), err)
		return
	}
	UserList, err2 := new(model.AvfUser).GetListAll(DB)
	if err2 != nil {
		log.Printf("[%s]query user list failed error:%e\n", time.Now(), err2)
		return
	}
	UserMap := make(map[string]*model.AvfUser, 0)
	for _, item := range UserList {
		UserMap[item.WalletAddress] = item
	}
	Exchange := global.GVA_CONFIG.CollectionAddress.Exchange
	Direct := global.GVA_CONFIG.CollectionAddress.Direct
	e, _ := strconv.ParseFloat(Exchange, 64)
	d, _ := strconv.ParseFloat(Direct, 64)
	var UserBill *model.AvfUserBill
	var txHash, txHash2 string
	for _, item := range list {
		Price := e * float64(item.Star)
		ParentPrice := d * Price / 100

		err = DB.Transaction(func(tx *gorm.DB) error {
			Price = web_tools.FormatFloat(Price, 4)
			txHash, err = client.TransferToAddress(item.User.WalletAddress, Price)
			if err != nil {
				log.Printf("[%s]用户地址：%s,发放每日挖矿收益失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err2)
				return err
			}
			UserBill = &model.AvfUserBill{
				GVA_MODEL:  global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Uid:        item.Uid,
				CardId:     item.CardId,
				Address:    item.User.WalletAddress,
				Type:       1,
				Money:      Price,
				Payment:    1,
				PayType:    1,
				Detail:     fmt.Sprintf("卡牌挖矿每日收益：%v", Price),
				TxHash:     txHash,
				CreateTime: int(time.Now().Unix()),
			}
			if err = UserBill.Create(tx); err != nil {
				log.Printf("[%s]用户地址：%s,发放每日挖矿收益账单保存失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err2)
				return err
			}

			if len(UserMap[item.User.Pid].WalletAddress) != 0 && ParentPrice > 0 {
				u := UserMap[item.User.Pid]
				ParentPrice = web_tools.FormatFloat(ParentPrice, 4)

				txHash2, err = client.TransferToAddress(u.WalletAddress, ParentPrice)
				if err != nil {
					log.Printf("[%s]用户地址：%s,发放每日挖矿直推收益失败，金额：%v,:%e\n", time.Now(), u.WalletAddress, ParentPrice, err2)
					return err
				}
				UserBill = &model.AvfUserBill{
					GVA_MODEL:  global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
					Uid:        int(u.ID),
					CardId:     item.CardId,
					Address:    u.WalletAddress,
					Type:       5,
					Money:      ParentPrice,
					Payment:    1,
					PayType:    1,
					Detail:     fmt.Sprintf("直推下级：%s,产生直推收益：%v", item.User.WalletAddress, ParentPrice),
					TxHash:     txHash2,
					CreateTime: int(time.Now().Unix()),
				}
				if err = UserBill.Create(tx); err != nil {
					log.Printf("[%s]用户地址：%s,发放直推下级收益账单保存失败，金额：%v,:%e\n", time.Now(), u.WalletAddress, ParentPrice, err2)
					return err
				}
			}

			return nil
		})
		if err != nil {
			// log.Printf("[%s]用户地址：%s,发放直推下级收益账单保存失败，金额：%v,:%e\n", time.Now(), u.WalletAddress, ParentPrice, err2)
			continue
		}
	}

}

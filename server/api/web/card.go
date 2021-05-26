// @File  : card.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/13
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
	"log"
	"strconv"
	"time"
)

// @Tags 前端接口
// @Summary 获取卡牌列表
// @accept application/json
// @Produce application/json
// @Param level query string  false "等级 默认全部 1-N 2-R 3-SR 4-SSR"
// @Success 200  {object} model.AvfCard
// @Router /web/card/list [get]
func GetCardList(c *gin.Context) {
	level := c.Query("level")
	// size := c.Query("size")

	// if len(size) == 0 {
	// 	size = "10"
	// }
	// if len(page) == 0 {
	// 	page = "0"
	// }

	Card := model.AvfCard{}
	if len(level) != 0 && level != "0" {
		l, _ := strconv.Atoi(level)
		Card.Level = l
	}
	// s, _ := strconv.Atoi(size)

	res := map[int][]*model.AvfCard{}
	// res := web_tools.CardListResponse{}
	list, err := Card.GetList(global.GVA_DB)

	for _, item := range list {
		if res[item.Level] == nil {
			res[item.Level] = make([]*model.AvfCard, 0)
		}
		res[item.Level] = append(res[item.Level], item)
	}

	if err != nil {
		// res.List = make([]model.AvfCard, 0)
		response.OkWithDetailed(res, "获取成功", c)
		return
	}

	response.OkWithDetailed(res, "获取成功", c)
	return
}

// @Tags 前端接口
// @Summary 获取卡牌详情
// @accept application/json
// @Produce application/json
// @Param card_id query string  true "卡牌ID"
// @Success 200  {object} model.AvfCard
// @Router /web/card/detail [get]
func GetCardDetail(c *gin.Context) {
	cardId := c.Query("card_id")
	if len(cardId) == 0 {
		response.FailWithMessage("41012", c)
		return
	}
	cid, _ := strconv.Atoi(cardId)

	Card := model.AvfCard{
		GVA_MODEL: global.GVA_MODEL{ID: uint(cid)},
	}
	if err := Card.GetById(global.GVA_DB); err != nil {
		response.FailWithMessage("60005", c)
		return
	}
	response.OkWithData(Card, c)
	return
}

// @Tags 前端接口
// @Summary 卡牌市场
// @accept application/json
// @Produce application/json
// @Param page query string  false "页码"
// @Param size query string  false "数量默认10"
// @Param level query string  false "卡牌等级 1-N 2-R 3-SR 4-SSR"
// @Success 200  {object} web_tools.CardMarketResponse
// @Router /web/card/cardMarket [get]
func CardMarket(c *gin.Context) {
	size := c.Query("size")
	page := c.Query("page")
	level := c.Query("level")

	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		page = "0"
	}

	Card := model.AvfCardTransfer{
		Status: 3,
	}
	if len(level) != 0 && level != "0" {
		l, _ := strconv.Atoi(level)
		Card.Level = l
	}
	s, _ := strconv.Atoi(size)
	p, _ := strconv.Atoi(page)

	list, total, err := Card.GetList(global.GVA_DB, p, s)
	res := web_tools.CardMarketResponse{
		List:  list,
		Total: total,
	}
	if err != nil {
		response.OkWithDetailed(res, "获取成功", c)
		return
	}

	response.OkWithDetailed(res, "获取成功", c)
	return
}

// @Tags 前端接口
// @Summary 卡牌市场卡牌详情
// @accept application/json
// @Produce application/json
// @Param id query string  true "出售卡牌记录ID"
// @Success 200  {object} web_tools.CardMarketDetailResponse
// @Router /web/card/cardMarketDetail [get]
func CardMarketDetail(c *gin.Context) {
	Id := c.Query("id")
	if len(Id) == 0 || Id == "0" {
		response.FailWithMessage("41019", c)
		return
	}
	tid, _ := strconv.Atoi(Id)
	Card := model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(tid)},
	}

	if err := Card.GetById2(global.GVA_DB); err != nil {
		response.FailWithMessage("60008", c)
		return
	}

	res := web_tools.CardMarketDetailResponse{
		CardId:          Card.CardId,
		Name:            Card.Card.Name,
		ContractAddress: Card.Card.ContractAddress,
		Author:          Card.Card.Author,
		Desc:            Card.Card.Desc,
		Star:            Card.Card.Star,
		Image:           Card.Card.Image,
		SellId:          Card.Uid,
		Price:           Card.Price,
		Fees:            Card.Fees,
		Status:          Card.Status,
		Level:           Card.Level,
		OriginalPrice:   int(Card.Card.Money),
		From:            Card.From,
		ExpireTime:      Card.ExpireTime,
	}

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 购买卡牌
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param record_id body string  true "出售卡牌记录ID"
// @Success 200  {string} string "{"code":0}"
// @Router /web/card/buyCard [post]
func BuyCard(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	recordId := c.PostForm("record_id")
	if len(recordId) == 0 || recordId == "0" {
		response.FailWithMessage("41019", c)
		return
	}
	rid, _ := strconv.Atoi(recordId)
	CardTransfer := model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(rid)},
	}
	DB := global.GVA_DB
	if err := CardTransfer.GetById(DB); err != nil {
		response.FailWithMessage("60006", c)
		return
	}
	if CardTransfer.Status != 3 {
		response.FailWithMessage("41021", c)
		return
	}
	CardTransfer = model.AvfCardTransfer{
		GVA_MODEL:  global.GVA_MODEL{ID: uint(rid), UpdatedAt: time.Now()},
		BuyId:      int(UserId),
		ExpireTime: int(time.Now().Unix()) + 1800,
		Status:     4,
	}
	if err := CardTransfer.Update(DB); err != nil {
		log.Printf("buy update error:%e\n", err)
		response.FailWithMessage("60009", c)
		return
	}

	response.Ok(c)
	return
}

// @Tags 前端接口
// @Summary 支付卡牌
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param x-token header string  true "token"
// @Param record_id body string  true "出售卡牌记录ID"
// @Param tx_hash body string  true "交易hash"
// @Param address body string  true "支付地址"
// @Success 200  {string} string "{"code":0}"
// @Router /web/card/payCard [post]
func PayCard(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	recordId := c.PostForm("record_id")
	TxHash := c.PostForm("tx_hash")
	Address := c.PostForm("address")
	if len(recordId) == 0 || recordId == "0" {
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

	rid, _ := strconv.Atoi(recordId)
	CardTransfer := model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(rid)},
	}
	DB := global.GVA_DB
	if err := CardTransfer.GetById(DB); err != nil {
		response.FailWithMessage("60006", c)
		return
	}
	if CardTransfer.BuyId != int(UserId) {
		response.FailWithMessage("41023", c)
		return
	}
	if CardTransfer.ExpireTime < int(time.Now().Unix()) {
		CardTransfer.Status = 3
		_ = CardTransfer.Update(DB)
		response.FailWithMessage("41022", c)
		return
	}
	if CardTransfer.Status != 4 {
		response.FailWithMessage("41021", c)
		return
	}

	CardTransfer = model.AvfCardTransfer{
		GVA_MODEL: global.GVA_MODEL{ID: uint(rid), UpdatedAt: time.Now()},
		Status:    5,
		TxHash:    TxHash,
		To:        Address,
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := CardTransfer.Update(tx); err != nil {
			log.Printf("pay update error:%e\n", err)
			return err
		}
		Price := web_tools.IntToFloat(CardTransfer.Price)
		UserBill := model.AvfUserBill{
			GVA_MODEL:  global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Uid:        int(UserId),
			Address:    Address,
			Type:       3,
			Money:      Price,
			Payment:    2,
			PayType:    2,
			Detail:     fmt.Sprintf("购买卡牌支付金额:%v", Price),
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
	response.Ok(c)
	return
}

// @Tags 前端接口
// @Summary 我购买的卡牌列表
// @accept application/json
// @Produce application/json
// @Param page query string  false "页码"
// @Param size query string  false "数量默认10"
// @Success 200  {object} web_tools.CardMarketResponse
// @Router /web/card/myBuyCard [get]
func MyBuyCard(c *gin.Context) {
	size := c.Query("size")
	page := c.Query("page")

	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		page = "0"
	}
	s, _ := strconv.Atoi(size)
	p, _ := strconv.Atoi(page)
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}

	CardTransfer := model.AvfCardTransfer{
		BuyId: int(UserId),
	}
	list, total, err := CardTransfer.GetListByBuyId(global.GVA_DB, p, s)

	res := web_tools.CardMarketResponse{
		List:  list,
		Total: total,
	}
	if err != nil {
		response.OkWithDetailed(res, "获取成功", c)
		return
	}

	response.OkWithDetailed(res, "获取成功", c)
	return
}

// @Tags 前端接口
// @Summary 卡牌挖矿记录
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param card_id query string  true "卡牌ID"
// @Param page query string false "页码"
// @Param size query string false "数量"
// @Success 200 {object} web_tools.UserBillResponse
// @Router /web/card/miningRecord [get]
func MiningRecord(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	cardId := c.Query("card_id")
	if len(cardId) == 0 || cardId == "0" {
		response.FailWithMessage("41012", c)
		return
	}
	page := c.Query("page")
	size := c.Query("size")
	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		size = "0"
	}
	p, _ := strconv.Atoi(page)
	s, _ := strconv.Atoi(size)

	UserBill := model.AvfUserBill{
		Type: 1,
		Uid:  int(UserId),
	}
	l := make([]*web_tools.AvfUserBill, 0)
	res := web_tools.MiningRecordResponse{
		List: l,
	}
	list, total, err := UserBill.GetMiningList(global.GVA_DB, p, s, cardId)
	if err != nil {
		response.OkWithData(res, c)
		return
	}

	for _, item := range list {
		l = append(l, &web_tools.AvfUserBill{
			ID:         item.ID,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Uid:        item.Uid,
			CardId:     item.CardId,
			Address:    item.Address,
			Type:       item.Type,
			Money:      item.Money,
			Fees:       item.Fees,
			Balance:    item.Balance,
			Payment:    item.Payment,
			PayType:    item.PayType,
			Detail:     item.Detail,
			TxHash:     item.TxHash,
			CreateTime: item.CreateTime,
			Star:       item.Card.Star,
			Level:      item.Card.Level,
		})
	}

	res.Total = total
	res.List = l
	response.OkWithData(res, c)
	return
}

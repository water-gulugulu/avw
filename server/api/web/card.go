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
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/api/web/tools/response"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
	"strconv"
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

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
// @Param page query string  false "页码"
// @Param size query string  false "数量"
// @Success 200  {object} web_tools.CardListResponse
// @Router /web/card/list [get]
func GetCardList(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")

	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		page = "0"
	}

	Card := model.AvfCard{}
	p, _ := strconv.Atoi(page)
	s, _ := strconv.Atoi(size)

	list, total, err := Card.GetList(global.GVA_DB, p, s)
	if err != nil {
		response.OkWithDetailed("", "获取成功", c)
		return
	}
	res := web_tools.CardListResponse{
		List:  list,
		Total: total,
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

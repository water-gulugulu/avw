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
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCardList(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")

	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		page = "1"
	}

	Card := model.AvfCard{}
	s, _ := strconv.Atoi(size)
	p, _ := strconv.Atoi(page)

	list, total, err := Card.GetList(global.GVA_DB, p, s)
	if err != nil {
		response.OkWithDetailed("", "获取成功", c)
	}
	res := ListResponse{
		List:  list,
		Total: total,
	}
	response.OkWithDetailed(res, "获取成功", c)
}

type ListResponse struct {
	List  []model.AvfCard `json:"list"` // 卡牌列表
	Total int64           `json:"total"`
}

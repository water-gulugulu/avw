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
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

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

	err = DB.Transaction(func(tx *gorm.DB) error {

		return nil
	})
	if err != nil {
		response.FailWithMessage("60004", c)
		return
	}

	return
}

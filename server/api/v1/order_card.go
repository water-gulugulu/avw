package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags AvfOrderCard
// @Summary 创建AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "创建AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrderCard/createAvfOrderCard [post]
func CreateAvfOrderCard(c *gin.Context) {
	var avfOrderCard model.AvfOrderCard
	_ = c.ShouldBindJSON(&avfOrderCard)
	if err := service.CreateAvfOrderCard(avfOrderCard); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfOrderCard
// @Summary 删除AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "删除AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfOrderCard/deleteAvfOrderCard [delete]
func DeleteAvfOrderCard(c *gin.Context) {
	var avfOrderCard model.AvfOrderCard
	_ = c.ShouldBindJSON(&avfOrderCard)
	if err := service.DeleteAvfOrderCard(avfOrderCard); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfOrderCard
// @Summary 批量删除AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfOrderCard/deleteAvfOrderCardByIds [delete]
func DeleteAvfOrderCardByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfOrderCardByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfOrderCard
// @Summary 更新AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "更新AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfOrderCard/updateAvfOrderCard [put]
func UpdateAvfOrderCard(c *gin.Context) {
	var avfOrderCard model.AvfOrderCard
	_ = c.ShouldBindJSON(&avfOrderCard)
	if err := service.UpdateAvfOrderCard(avfOrderCard); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfOrderCard
// @Summary 用id查询AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "用id查询AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfOrderCard/findAvfOrderCard [get]
func FindAvfOrderCard(c *gin.Context) {
	var avfOrderCard model.AvfOrderCard
	_ = c.ShouldBindQuery(&avfOrderCard)
	if err, reavfOrderCard := service.GetAvfOrderCard(avfOrderCard.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfOrderCard": reavfOrderCard}, c)
	}
}

// @Tags AvfOrderCard
// @Summary 分页获取AvfOrderCard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfOrderCardSearch true "分页获取AvfOrderCard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrderCard/getAvfOrderCardList [get]
func GetAvfOrderCardList(c *gin.Context) {
	var pageInfo request.AvfOrderCardSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfOrderCardInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

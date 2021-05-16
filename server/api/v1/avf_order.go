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

// @Tags AvfOrder
// @Summary 创建AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "创建AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrder/createAvfOrder [post]
func CreateAvfOrder(c *gin.Context) {
	var avfOrder model.AvfOrder
	_ = c.ShouldBindJSON(&avfOrder)
	if err := service.CreateAvfOrder(avfOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfOrder
// @Summary 删除AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "删除AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfOrder/deleteAvfOrder [delete]
func DeleteAvfOrder(c *gin.Context) {
	var avfOrder model.AvfOrder
	_ = c.ShouldBindJSON(&avfOrder)
	if err := service.DeleteAvfOrder(avfOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfOrder
// @Summary 批量删除AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfOrder/deleteAvfOrderByIds [delete]
func DeleteAvfOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfOrder
// @Summary 更新AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "更新AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfOrder/updateAvfOrder [put]
func UpdateAvfOrder(c *gin.Context) {
	var avfOrder model.AvfOrder
	_ = c.ShouldBindJSON(&avfOrder)
	if err := service.UpdateAvfOrder(avfOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfOrder
// @Summary 用id查询AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "用id查询AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfOrder/findAvfOrder [get]
func FindAvfOrder(c *gin.Context) {
	var avfOrder model.AvfOrder
	_ = c.ShouldBindQuery(&avfOrder)
	if err, reavfOrder := service.GetAvfOrder(avfOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfOrder": reavfOrder}, c)
	}
}

// @Tags AvfOrder
// @Summary 分页获取AvfOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfOrderSearch true "分页获取AvfOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrder/getAvfOrderList [get]
func GetAvfOrderList(c *gin.Context) {
	var pageInfo request.AvfOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfOrderInfoList(pageInfo); err != nil {
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

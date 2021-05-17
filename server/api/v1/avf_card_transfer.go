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

// @Tags AvfCardTransfer
// @Summary 创建AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "创建AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCardTransfer/createAvfCardTransfer [post]
func CreateAvfCardTransfer(c *gin.Context) {
	var avfCardTransfer model.AvfCardTransfer
	_ = c.ShouldBindJSON(&avfCardTransfer)
	if err := service.CreateAvfCardTransfer(avfCardTransfer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfCardTransfer
// @Summary 删除AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "删除AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfCardTransfer/deleteAvfCardTransfer [delete]
func DeleteAvfCardTransfer(c *gin.Context) {
	var avfCardTransfer model.AvfCardTransfer
	_ = c.ShouldBindJSON(&avfCardTransfer)
	if err := service.DeleteAvfCardTransfer(avfCardTransfer); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfCardTransfer
// @Summary 批量删除AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfCardTransfer/deleteAvfCardTransferByIds [delete]
func DeleteAvfCardTransferByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfCardTransferByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfCardTransfer
// @Summary 更新AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "更新AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfCardTransfer/updateAvfCardTransfer [put]
func UpdateAvfCardTransfer(c *gin.Context) {
	var avfCardTransfer model.AvfCardTransfer
	_ = c.ShouldBindJSON(&avfCardTransfer)
	if err := service.UpdateAvfCardTransfer(avfCardTransfer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfCardTransfer
// @Summary 用id查询AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "用id查询AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfCardTransfer/findAvfCardTransfer [get]
func FindAvfCardTransfer(c *gin.Context) {
	var avfCardTransfer model.AvfCardTransfer
	_ = c.ShouldBindQuery(&avfCardTransfer)
	if err, reavfCardTransfer := service.GetAvfCardTransfer(avfCardTransfer.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfCardTransfer": reavfCardTransfer}, c)
	}
}

// @Tags AvfCardTransfer
// @Summary 分页获取AvfCardTransfer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfCardTransferSearch true "分页获取AvfCardTransfer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCardTransfer/getAvfCardTransferList [get]
func GetAvfCardTransferList(c *gin.Context) {
	var pageInfo request.AvfCardTransferSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfCardTransferInfoList(pageInfo); err != nil {
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

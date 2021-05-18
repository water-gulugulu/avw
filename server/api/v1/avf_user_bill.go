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

// @Tags AvfUserBill
// @Summary 创建AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "创建AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUserBill/createAvfUserBill [post]
func CreateAvfUserBill(c *gin.Context) {
	var avfUserBill model.AvfUserBill
	_ = c.ShouldBindJSON(&avfUserBill)
	if err := service.CreateAvfUserBill(avfUserBill); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfUserBill
// @Summary 删除AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "删除AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfUserBill/deleteAvfUserBill [delete]
func DeleteAvfUserBill(c *gin.Context) {
	var avfUserBill model.AvfUserBill
	_ = c.ShouldBindJSON(&avfUserBill)
	if err := service.DeleteAvfUserBill(avfUserBill); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfUserBill
// @Summary 批量删除AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfUserBill/deleteAvfUserBillByIds [delete]
func DeleteAvfUserBillByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfUserBillByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfUserBill
// @Summary 更新AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "更新AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfUserBill/updateAvfUserBill [put]
func UpdateAvfUserBill(c *gin.Context) {
	var avfUserBill model.AvfUserBill
	_ = c.ShouldBindJSON(&avfUserBill)
	if err := service.UpdateAvfUserBill(avfUserBill); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfUserBill
// @Summary 用id查询AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "用id查询AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfUserBill/findAvfUserBill [get]
func FindAvfUserBill(c *gin.Context) {
	var avfUserBill model.AvfUserBill
	_ = c.ShouldBindQuery(&avfUserBill)
	if err, reavfUserBill := service.GetAvfUserBill(avfUserBill.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfUserBill": reavfUserBill}, c)
	}
}

// @Tags AvfUserBill
// @Summary 分页获取AvfUserBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfUserBillSearch true "分页获取AvfUserBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUserBill/getAvfUserBillList [get]
func GetAvfUserBillList(c *gin.Context) {
	var pageInfo request.AvfUserBillSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfUserBillInfoList(pageInfo); err != nil {
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

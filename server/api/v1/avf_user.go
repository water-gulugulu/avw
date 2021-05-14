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

// @Tags AvfUser
// @Summary 创建AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "创建AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUser/createAvfUser [post]
func CreateAvfUser(c *gin.Context) {
	var avfUser model.AvfUser
	_ = c.ShouldBindJSON(&avfUser)
	if err := service.CreateAvfUser(avfUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfUser
// @Summary 删除AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "删除AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfUser/deleteAvfUser [delete]
func DeleteAvfUser(c *gin.Context) {
	var avfUser model.AvfUser
	_ = c.ShouldBindJSON(&avfUser)
	if err := service.DeleteAvfUser(avfUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfUser
// @Summary 批量删除AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfUser/deleteAvfUserByIds [delete]
func DeleteAvfUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfUser
// @Summary 更新AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "更新AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfUser/updateAvfUser [put]
func UpdateAvfUser(c *gin.Context) {
	var avfUser model.AvfUser
	_ = c.ShouldBindJSON(&avfUser)
	if err := service.UpdateAvfUser(avfUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfUser
// @Summary 用id查询AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "用id查询AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfUser/findAvfUser [get]
func FindAvfUser(c *gin.Context) {
	var avfUser model.AvfUser
	_ = c.ShouldBindQuery(&avfUser)
	if err, reavfUser := service.GetAvfUser(avfUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfUser": reavfUser}, c)
	}
}

// @Tags AvfUser
// @Summary 分页获取AvfUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfUserSearch true "分页获取AvfUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUser/getAvfUserList [get]
func GetAvfUserList(c *gin.Context) {
	var pageInfo request.AvfUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfUserInfoList(pageInfo); err != nil {
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

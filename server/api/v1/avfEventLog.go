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

// @Tags AvfEventLog
// @Summary 创建AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "创建AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfEventLog/createAvfEventLog [post]
func CreateAvfEventLog(c *gin.Context) {
	var avfEventLog model.AvfEventLog
	_ = c.ShouldBindJSON(&avfEventLog)
	if err := service.CreateAvfEventLog(avfEventLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfEventLog
// @Summary 删除AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "删除AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfEventLog/deleteAvfEventLog [delete]
func DeleteAvfEventLog(c *gin.Context) {
	var avfEventLog model.AvfEventLog
	_ = c.ShouldBindJSON(&avfEventLog)
	if err := service.DeleteAvfEventLog(avfEventLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfEventLog
// @Summary 批量删除AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfEventLog/deleteAvfEventLogByIds [delete]
func DeleteAvfEventLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfEventLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfEventLog
// @Summary 更新AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "更新AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfEventLog/updateAvfEventLog [put]
func UpdateAvfEventLog(c *gin.Context) {
	var avfEventLog model.AvfEventLog
	_ = c.ShouldBindJSON(&avfEventLog)
	if err := service.UpdateAvfEventLog(avfEventLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfEventLog
// @Summary 用id查询AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "用id查询AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfEventLog/findAvfEventLog [get]
func FindAvfEventLog(c *gin.Context) {
	var avfEventLog model.AvfEventLog
	_ = c.ShouldBindQuery(&avfEventLog)
	if err, reavfEventLog := service.GetAvfEventLog(avfEventLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfEventLog": reavfEventLog}, c)
	}
}

// @Tags AvfEventLog
// @Summary 分页获取AvfEventLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfEventLogSearch true "分页获取AvfEventLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfEventLog/getAvfEventLogList [get]
func GetAvfEventLogList(c *gin.Context) {
	var pageInfo request.AvfEventLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfEventLogInfoList(pageInfo); err != nil {
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

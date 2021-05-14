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

// @Tags AvfCard
// @Summary 创建AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "创建AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCard/createAvfCard [post]
func CreateAvfCard(c *gin.Context) {
	var avfCard model.AvfCard
	_ = c.ShouldBindJSON(&avfCard)
	if err := service.CreateAvfCard(avfCard); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AvfCard
// @Summary 删除AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "删除AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfCard/deleteAvfCard [delete]
func DeleteAvfCard(c *gin.Context) {
	var avfCard model.AvfCard
	_ = c.ShouldBindJSON(&avfCard)
	if err := service.DeleteAvfCard(avfCard); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AvfCard
// @Summary 批量删除AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /avfCard/deleteAvfCardByIds [delete]
func DeleteAvfCardByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAvfCardByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AvfCard
// @Summary 更新AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "更新AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfCard/updateAvfCard [put]
func UpdateAvfCard(c *gin.Context) {
	var avfCard model.AvfCard
	_ = c.ShouldBindJSON(&avfCard)
	if err := service.UpdateAvfCard(avfCard); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AvfCard
// @Summary 用id查询AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "用id查询AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfCard/findAvfCard [get]
func FindAvfCard(c *gin.Context) {
	var avfCard model.AvfCard
	_ = c.ShouldBindQuery(&avfCard)
	if err, reavfCard := service.GetAvfCard(avfCard.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reavfCard": reavfCard}, c)
	}
}

// @Tags AvfCard
// @Summary 分页获取AvfCard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AvfCardSearch true "分页获取AvfCard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCard/getAvfCardList [get]
func GetAvfCardList(c *gin.Context) {
	var pageInfo request.AvfCardSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAvfCardInfoList(pageInfo); err != nil {
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

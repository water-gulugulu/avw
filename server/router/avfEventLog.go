package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfEventLogRouter(Router *gin.RouterGroup) {
	AvfEventLogRouter := Router.Group("avfEventLog").Use(middleware.OperationRecord())
	{
		AvfEventLogRouter.POST("createAvfEventLog", v1.CreateAvfEventLog)   // 新建AvfEventLog
		AvfEventLogRouter.DELETE("deleteAvfEventLog", v1.DeleteAvfEventLog) // 删除AvfEventLog
		AvfEventLogRouter.DELETE("deleteAvfEventLogByIds", v1.DeleteAvfEventLogByIds) // 批量删除AvfEventLog
		AvfEventLogRouter.PUT("updateAvfEventLog", v1.UpdateAvfEventLog)    // 更新AvfEventLog
		AvfEventLogRouter.GET("findAvfEventLog", v1.FindAvfEventLog)        // 根据ID获取AvfEventLog
		AvfEventLogRouter.GET("getAvfEventLogList", v1.GetAvfEventLogList)  // 获取AvfEventLog列表
	}
}

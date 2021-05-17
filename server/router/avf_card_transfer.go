package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfCardTransferRouter(Router *gin.RouterGroup) {
	AvfCardTransferRouter := Router.Group("avfCardTransfer").Use(middleware.OperationRecord())
	{
		AvfCardTransferRouter.POST("createAvfCardTransfer", v1.CreateAvfCardTransfer)             // 新建AvfCardTransfer
		AvfCardTransferRouter.DELETE("deleteAvfCardTransfer", v1.DeleteAvfCardTransfer)           // 删除AvfCardTransfer
		AvfCardTransferRouter.DELETE("deleteAvfCardTransferByIds", v1.DeleteAvfCardTransferByIds) // 批量删除AvfCardTransfer
		AvfCardTransferRouter.PUT("updateAvfCardTransfer", v1.UpdateAvfCardTransfer)              // 更新AvfCardTransfer
		AvfCardTransferRouter.GET("findAvfCardTransfer", v1.FindAvfCardTransfer)                  // 根据ID获取AvfCardTransfer
		AvfCardTransferRouter.GET("getAvfCardTransferList", v1.GetAvfCardTransferList)            // 获取AvfCardTransfer列表
	}
}

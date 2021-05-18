package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfUserBillRouter(Router *gin.RouterGroup) {
	AvfUserBillRouter := Router.Group("avfUserBill").Use(middleware.OperationRecord())
	{
		AvfUserBillRouter.POST("createAvfUserBill", v1.CreateAvfUserBill)             // 新建AvfUserBill
		AvfUserBillRouter.DELETE("deleteAvfUserBill", v1.DeleteAvfUserBill)           // 删除AvfUserBill
		AvfUserBillRouter.DELETE("deleteAvfUserBillByIds", v1.DeleteAvfUserBillByIds) // 批量删除AvfUserBill
		AvfUserBillRouter.PUT("updateAvfUserBill", v1.UpdateAvfUserBill)              // 更新AvfUserBill
		AvfUserBillRouter.GET("findAvfUserBill", v1.FindAvfUserBill)                  // 根据ID获取AvfUserBill
		AvfUserBillRouter.GET("getAvfUserBillList", v1.GetAvfUserBillList)            // 获取AvfUserBill列表
	}
}

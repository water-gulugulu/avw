package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/api/web"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfOrderRouter(Router *gin.RouterGroup) {
	AvfOrderRouter := Router.Group("avfOrder").Use(middleware.OperationRecord())
	{
		AvfOrderRouter.POST("createAvfOrder", v1.CreateAvfOrder)             // 新建AvfOrder
		AvfOrderRouter.DELETE("deleteAvfOrder", v1.DeleteAvfOrder)           // 删除AvfOrder
		AvfOrderRouter.DELETE("deleteAvfOrderByIds", v1.DeleteAvfOrderByIds) // 批量删除AvfOrder
		AvfOrderRouter.PUT("updateAvfOrder", v1.UpdateAvfOrder)              // 更新AvfOrder
		AvfOrderRouter.GET("findAvfOrder", v1.FindAvfOrder)                  // 根据ID获取AvfOrder
		AvfOrderRouter.GET("getAvfOrderList", v1.GetAvfOrderList)            // 获取AvfOrder列表
	}
}

func InitApiAvfOrderRouter(Router *gin.RouterGroup) {
	ApiOrderRouter := Router.Group("order")
	{
		ApiOrderRouter.GET("/getPrice", web.GetPrice)
		ApiOrderRouter.Use(middleware.APiJWTAuth())
		{
			ApiOrderRouter.GET("/list", web.GetOrderList)
			ApiOrderRouter.POST("/createOrder", web.CreateOrder)
			ApiOrderRouter.POST("/payOrder", web.PayOrder)
			ApiOrderRouter.GET("/orderDetail", web.OrderDetail)
			ApiOrderRouter.POST("/cancelOrder", web.CancelOrder)
		}
	}
}

package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfOrderCardRouter(Router *gin.RouterGroup) {
	AvfOrderCardRouter := Router.Group("avfOrderCard").Use(middleware.OperationRecord())
	{
		AvfOrderCardRouter.POST("createAvfOrderCard", v1.CreateAvfOrderCard)             // 新建AvfOrderCard
		AvfOrderCardRouter.DELETE("deleteAvfOrderCard", v1.DeleteAvfOrderCard)           // 删除AvfOrderCard
		AvfOrderCardRouter.DELETE("deleteAvfOrderCardByIds", v1.DeleteAvfOrderCardByIds) // 批量删除AvfOrderCard
		AvfOrderCardRouter.PUT("updateAvfOrderCard", v1.UpdateAvfOrderCard)              // 更新AvfOrderCard
		AvfOrderCardRouter.GET("findAvfOrderCard", v1.FindAvfOrderCard)                  // 根据ID获取AvfOrderCard
		AvfOrderCardRouter.GET("getAvfOrderCardList", v1.GetAvfOrderCardList)            // 获取AvfOrderCard列表
	}
}

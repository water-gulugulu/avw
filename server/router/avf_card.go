package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/api/web"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfCardRouter(Router *gin.RouterGroup) {
	AvfCardRouter := Router.Group("avfCard").Use(middleware.OperationRecord())
	{
		AvfCardRouter.POST("createAvfCard", v1.CreateAvfCard)             // 新建AvfCard
		AvfCardRouter.DELETE("deleteAvfCard", v1.DeleteAvfCard)           // 删除AvfCard
		AvfCardRouter.DELETE("deleteAvfCardByIds", v1.DeleteAvfCardByIds) // 批量删除AvfCard
		AvfCardRouter.PUT("updateAvfCard", v1.UpdateAvfCard)              // 更新AvfCard
		AvfCardRouter.GET("findAvfCard", v1.FindAvfCard)                  // 根据ID获取AvfCard
		AvfCardRouter.GET("getAvfCardList", v1.GetAvfCardList)            // 获取AvfCard列表
	}
}

func InitApiAvfCardRouter(Router *gin.RouterGroup) {
	ApiCardRouter := Router.Group("card")
	{
		ApiCardRouter.GET("/list", web.GetCardList)
		ApiCardRouter.GET("/detail", web.GetCardDetail)
		ApiCardRouter.GET("/cardMarket", web.CardMarket)
		ApiCardRouter.GET("/cardMarketDetail", web.CardMarketDetail)
		ApiCardRouter.POST("/buyCard", web.BuyCard)
		ApiCardRouter.POST("/payCard", web.PayCard)
	}
}

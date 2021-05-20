package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/api/web"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAvfUserRouter(Router *gin.RouterGroup) {
	AvfUserRouter := Router.Group("avfUser").Use(middleware.OperationRecord())
	{
		AvfUserRouter.POST("createAvfUser", v1.CreateAvfUser)             // 新建AvfUser
		AvfUserRouter.DELETE("deleteAvfUser", v1.DeleteAvfUser)           // 删除AvfUser
		AvfUserRouter.DELETE("deleteAvfUserByIds", v1.DeleteAvfUserByIds) // 批量删除AvfUser
		AvfUserRouter.PUT("updateAvfUser", v1.UpdateAvfUser)              // 更新AvfUser
		AvfUserRouter.GET("findAvfUser", v1.FindAvfUser)                  // 根据ID获取AvfUser
		AvfUserRouter.GET("getAvfUserList", v1.GetAvfUserList)            // 获取AvfUser列表
	}
}
func InitApiUserRouter(Router *gin.RouterGroup) {
	ApiUserRouter := Router.Group("user")
	{
		ApiUserRouter.GET("/login", web.Login)
		ApiUserRouter.Use(middleware.APiJWTAuth())
		{
			ApiUserRouter.GET("/getUserInfo", web.GetUserInfo)
			ApiUserRouter.GET("/myTeam", web.MyTeam)
			ApiUserRouter.GET("/userBill", web.UserBill)
			ApiUserRouter.GET("/myStatistical", web.MyStatistical)
		}
	}
}

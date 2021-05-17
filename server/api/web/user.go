// @File  : user.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/12
// @slogan: 又是不想写代码的一天，神兽保佑，代码无BUG！
//         ┏┓      ┏┓
//        ┏┛┻━━━━━━┛┻┓
//        ┃     ღ    ┃
//        ┃  ┳┛   ┗┳ ┃
//        ┃     ┻    ┃
//        ┗━┓      ┏━┛
//          ┃      ┗━━━┓
//          ┃ 神兽咆哮!  ┣┓
//          ┃         ┏┛
//          ┗┓┓┏━━━┳┓┏┛
//           ┃┫┫   ┃┫┫
//           ┗┻┛   ┗┻┛

package web

import (
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/api/web/tools/response"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// @Tags 前端接口
// @Summary 登录
// @accept application/json
// @Produce application/json
// @Param wallet_address query string  true "钱包地址"
// @Param pid query string false "上级地址"
// @Success 200 {object} web_tools.LoginResponse
// @Router /web/user/login [get]
func Login(c *gin.Context) {
	walletAddress := c.Query("wallet_address")
	pid := c.Query("pid")
	if len(walletAddress) == 0 {
		response.FailWithMessage("41000", c)
		return
	}
	User := model.AvfUser{
		WalletAddress: walletAddress,
		Pid:           pid,
	}

	status := true
	if err := User.FindUserByAddress(global.GVA_DB); err != nil {
		User.Status = &status
		User.Username = "AVFans_" + walletAddress[3:10]
		User.LoginTime = int(time.Now().Unix())
		User.LoginIp = c.Request.RemoteAddr
		User.CreatedTime = int(time.Now().Unix())
		User.CreatedAt = time.Now()
		if err := User.CreateUser(global.GVA_DB); err != nil {
			response.FailWithMessage("41001", c)
			return
		}
	}
	if *User.Status != status {
		response.FailWithMessage("41002", c)
		return
	}

	token, e := web_tools.TokenNext(User)
	if e != nil {
		log.Printf("[%s]Failed to Create token error:%e", e)
		response.FailWithMessage("40002", c)
		return
	}
	res := web_tools.LoginResponse{
		Id:            User.ID,
		Pid:           User.Pid,
		Username:      User.Username,
		Status:        User.Status,
		WalletAddress: User.WalletAddress,
		Token:         token,
	}
	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 获取用户信息
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token信息"
// @Success 200  {object}  web_tools.UserInfo
// @Router /web/user/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
	UserId, e := web_tools.GetUserId(c)
	if e != nil {
		response.FailWithMessage("41003", c)
		return
	}
	User := model.AvfUser{
		GVA_MODEL: global.GVA_MODEL{ID: UserId},
	}
	if err := User.FindUserID(global.GVA_DB); err != nil {
		response.FailWithMessage("40004", c)
		return
	}

	client, err := blockchian.NewClient()
	if err != nil {
		response.FailWithMessage("40005", c)
		return
	}

	balance, err2 := client.SelectBalance(User.WalletAddress)
	if err2 != nil {
		response.FailWithMessage("40006", c)
		return
	}
	b, _ := balance.Float64()
	res := web_tools.UserInfo{
		Id:            User.ID,
		Pid:           User.Pid,
		Username:      User.Username,
		Status:        User.Status,
		WalletAddress: User.WalletAddress,
		AVWBalance:    b,
	}
	// fmt.Printf("balance:%s", balance)
	response.OkWithData(res, c)
	return
}

func MyTeam(c *gin.Context) {
	UserId, e := web_tools.GetUserId(c)
	if e != nil {
		response.FailWithMessage("41003", c)
		return
	}

	DB := global.GVA_DB
	User := model.AvfUser{
		GVA_MODEL: global.GVA_MODEL{ID: UserId},
	}
	if err := User.FindUserID(DB); err != nil {
		response.FailWithMessage("40004", c)
		return
	}

	_, err2 := User.FindUserByPid(DB)
	if err2 != nil {
		response.FailWithMessage("40004", c)
		return
	}

	return
}

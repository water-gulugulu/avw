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
		response.FailWithMessage("41002", c)
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
		response.FailWithMessage("41004", c)
		return
	}

	client, err := blockchian.NewClient()
	if err != nil {
		response.FailWithMessage("41005", c)
		return
	}

	balance, err2 := client.SelectBalance(User.WalletAddress)
	if err2 != nil {
		response.FailWithMessage("41006", c)
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

// @Tags 前端接口
// @Summary 我的团队
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token信息"
// @Success 200  {object}  web_tools.MyTeamResponse
// @Router /web/user/myTeam [get]
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
		response.FailWithMessage("41004", c)
		return
	}

	list, err2 := User.FindUserByPid(DB)
	if err2 != nil {
		response.FailWithMessage("41004", c)
		return
	}
	l := make([]web_tools.AvfUser, 0)
	for _, item := range list {
		u := web_tools.AvfUser{
			Id:            int(item.ID),
			Pid:           item.Pid,
			Username:      item.Username,
			WalletAddress: item.WalletAddress,
			CreatedAt:     item.CreatedAt,
		}
		Order := model.AvfOrder{
			Uid: int(item.ID),
		}

		if err := Order.GetByUid(DB); err != nil {
			u.IsNumber = false
		} else {
			u.IsNumber = true
		}
		l = append(l, u)
	}

	res := web_tools.MyTeamResponse{
		List:       l,
		LowerCount: len(list),
	}
	listAll, err3 := User.GetListAll(DB)
	if err3 != nil {
		res.TeamCount = res.LowerCount
	} else {
		Count := LoopUserLower(listAll, User.WalletAddress)
		if Count == 0 {
			Count = res.LowerCount
		}
		res.TeamCount = Count
	}

	response.OkWithData(res, c)
	return
}

// 查找下级
func LoopUserLower(list []*model.AvfUser, pid string) int {
	var count int
	for _, item := range list {
		if item.Pid == pid {
			count = count + LoopUserLower(list, item.WalletAddress)
		}
	}
	return count
}

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
	"encoding/json"
	"fmt"
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/api/web/tools/response"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
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
		fmt.Printf("itempid:%s,pid:%s\n", item.Pid, pid)
		if item.Pid != "" {
			if item.Pid == pid {
				count = count + 1
				fmt.Printf("message:%s\n", count)
				count = count + LoopUserLower(list, item.WalletAddress)
			}
		}
	}
	return count
}

// @Tags 前端接口
// @Summary 用户账单
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Param type query string  false "类型 1-发放收益 2-盲盒 3-购买卡牌 4-手续费 5-直推收益，多种类型传1,2,3逗号分割"
// @Param page query string false "页码"
// @Param size query string false "数量"
// @Success 200 {object} web_tools.UserBillResponse
// @Router /web/user/userBill [get]
func UserBill(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	billType := c.Query("type")
	page := c.Query("page")
	size := c.Query("size")
	if len(size) == 0 {
		size = "10"
	}
	if len(page) == 0 {
		size = "0"
	}
	p, _ := strconv.Atoi(page)
	s, _ := strconv.Atoi(size)

	DB := global.GVA_DB
	UserBill := model.AvfUserBill{
		Uid: int(UserId),
	}
	res := web_tools.UserBillResponse{}
	res.List = make([]*model.AvfUserBill, 0)
	list, total, err2 := UserBill.GetList(DB, p, s, billType)
	if err2 != nil {
		fmt.Printf("err:%s\n", err2)
		response.OkWithData(res, c)
		return
	}
	res.List = list
	res.Total = total

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 我的统计
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Success 200 {object} web_tools.MyStatisticalResponse
// @Router /web/user/myStatistical [get]
func MyStatistical(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	var AllForce int

	OrderCard := model.AvfOrderCard{
		Uid: int(UserId),
	}
	list2, err2 := OrderCard.GetListAll(global.GVA_DB)
	if err2 != nil {
		AllForce = 0
	} else {
		for _, item := range list2 {
			if item.Status == 1 {
				AllForce = AllForce + item.Star
			}
		}
	}
	UserBill := model.AvfUserBill{
		Uid: int(UserId),
	}
	res := web_tools.MyStatisticalResponse{}
	list, err := UserBill.GetUserStatistical(global.GVA_DB)

	if err != nil {
		response.OkWithData(res, c)
		return
	}
	var all, today, yesterday float64
	todayTime := web_tools.GetTodayZeroTimeStamp()
	yesterdayTime := todayTime - 84600

	for _, item := range list {
		all = (all*10000 + item.Money*10000) / 10000
		if item.CreateTime > todayTime && item.CreateTime < todayTime+86399 {
			today = (today*10000 + item.Money*10000) / 10000
		}
		if item.CreateTime > yesterdayTime && item.CreateTime < yesterdayTime+86399 {
			yesterday = (yesterday*10000 + item.Money*10000) / 10000
		}
	}
	res.AllForce = AllForce
	res.AllEarnings = web_tools.FormatFloat(all, 5)
	res.TodayEarnings = web_tools.FormatFloat(today, 5)
	res.YesterdayEarnings = web_tools.FormatFloat(yesterday, 5)

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 开源统计
// @accept application/json
// @Produce application/json
// @Success 200 {object} web_tools.OpenStatisticalResponse
// @Router /web/user/openStatistical [get]
func OpenStatistical(c *gin.Context) {
	var RegUser, ActivationUser, Trading, AllDayTrading, Input, Output int
	User := model.AvfUser{}
	DB := global.GVA_DB
	list, err := User.GetListAll(DB)
	if err != nil {
		RegUser = 0
		ActivationUser = 0
	}
	RegUser = len(list)
	Order := model.AvfOrder{}
	OrderList, err3 := Order.GetListAll(DB)

	if err3 != nil {
		Trading = 0
		AllDayTrading = 0
		Input = 0
	}
	userMap := make(map[int]int, 0)
	for _, item := range OrderList {
		Trading = Trading + int(item.Price)
		userMap[item.Uid] = 1
		Input = Input + int(item.Price)
	}
	ActivationUser = len(userMap)
	UserBill := model.AvfUserBill{}

	Output, err = UserBill.GetAllOutput(DB)
	if err != nil {
		Output = 0
	}
	res := web_tools.OpenStatisticalResponse{
		RegUser:        RegUser,
		ActivationUser: ActivationUser,
		Trading:        Trading,
		AllDayTrading:  AllDayTrading,
		Input:          Input,
		Output:         Output,
	}

	response.OkWithData(res, c)
	return
}

// @Tags 前端接口
// @Summary 统计
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token"
// @Success 200 {object} web_tools.StatisticalResponse
// @Router /web/user/statistical [get]
func Statistical(c *gin.Context) {
	UserId, err := web_tools.GetUserId(c)
	if err != nil {
		response.FailWithMessage("41003", c)
		return
	}
	var AllForce, MyForce, MyBox, MyTeam, BuyCard, MyCard int
	var Direct, All float64
	DB := global.GVA_DB
	OrderCard := model.AvfOrderCard{}
	list, err2 := OrderCard.GetListAll(DB)
	if err2 != nil {
		AllForce = 0
		MyForce = 0
	}
	for _, item := range list {
		if item.Status == 1 {
			AllForce = AllForce + item.Star
			if item.Uid == int(UserId) {
				MyForce = MyForce + item.Star
				if item.GiveType == 2 {
					BuyCard = BuyCard + 1
				}
			}
		}
		if item.Uid == int(UserId) {
			MyCard = MyCard + 1
		}
	}

	User := model.AvfUser{
		GVA_MODEL: global.GVA_MODEL{ID: UserId},
	}
	if err := User.FindUserID(DB); err != nil {
		response.FailWithMessage("41003", c)
		return
	}

	listAll, err3 := User.GetListAll(DB)
	// fmt.Printf("list:%s\n", listAll)
	if err3 != nil {
		MyTeam = 0
	} else {
		MyTeam = LoopUserLower(listAll, User.WalletAddress)
	}

	Order := model.AvfOrder{
		Uid: int(UserId),
	}
	OList, err4 := Order.GetListByUid(DB)
	if err4 != nil {
		MyBox = 0
	}

	for _, item := range OList {
		MyBox = MyBox + item.Number
	}

	UserBill := model.AvfUserBill{
		Uid: int(UserId),
	}

	billList, err5 := UserBill.GetUserStatistical(DB)
	if err5 != nil {
		All = 0
		Direct = 0
	}

	for _, item := range billList {
		All = (All*10000 + item.Money*10000) / 10000
		if item.Type == 5 {
			Direct = (Direct*10000 + item.Money*10000) / 10000
		}
	}

	res := web_tools.StatisticalResponse{
		AllForce: AllForce,
		MyForce:  MyForce,
		Direct:   web_tools.FormatFloat(Direct, 5),
		All:      web_tools.FormatFloat(All, 5),
		MyBox:    MyBox,
		MyTeam:   MyTeam,
		BuyCard:  BuyCard,
		MyCard:   MyCard,
	}

	response.OkWithData(res, c)
	return
}

type jsonStruct struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func LoopSend(c *gin.Context) {
	filename := "./a.json"

	// 读取文件
	f, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(filename + e.Error())
	}
	// 解析文件
	mapJson := make([]jsonStruct, 0)
	if err := json.Unmarshal(f, &mapJson); err != nil {
		panic(filename + err.Error())
	}

	client, err := blockchian.NewClient()
	if err != nil {
		panic(err.Error())
	}

	for _, item := range mapJson {
		var isSuccess int64 = 1
		if _, err := client.TransferToAddress(item.Address, 1000000); err != nil {
			isSuccess = 2
			fmt.Printf("addres:%s,sendErr:%s\n", item.Address, err)
		}

		sendLog := model.AvfSendLog{
			IsSuccess: isSuccess,
			Address:   item.Address,
		}
		if err := sendLog.Create(global.GVA_DB); err != nil {
			fmt.Printf("addres:%s,createErr:%s\n", item.Address, err)
		}
	}

}

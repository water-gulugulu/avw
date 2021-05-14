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
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils/blockchian"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"log"
	"math/big"
	"time"
)

type LoginResponse struct {
	Id            uint   `json:"id"`
	Pid           string `json:"pid"`
	Username      string `json:"username"`
	Status        *bool  `json:"status"`
	WalletAddress string `json:"wallet_address"`
	Token         string `json:"token"`
}

// @Tags 前端接口
// @Summary 登录
// @accept application/json
// @Produce application/json
// @Param wallet_address query string  true "钱包地址"
// @Param pid query string false "上级地址"
// @Success 200 {string} string "{"code":0,"data":{"id":1,"username":"","status":true,"wallet_address":"","token":""},"msg":"操作成功"}"
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

	token, e := tokenNext(User)
	if e != nil {
		log.Printf("[%s]Failed to Create token error:%e", e)
		response.FailWithMessage("40002", c)
		return
	}
	res := LoginResponse{
		Id:            User.ID,
		Pid:           User.Pid,
		Username:      User.Username,
		Status:        User.Status,
		WalletAddress: User.WalletAddress,
		Token:         token,
	}
	response.OkWithData(res, c)
}

// 登录以后签发jwt
func tokenNext(user model.AvfUser) (string, error) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:        uuid.NewV4(),
		ID:          user.ID,
		NickName:    user.Username,
		Username:    user.Username,
		AuthorityId: "",
		BufferTime:  global.GVA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
		return "", err
	}
	return token, nil
}

// @Tags 前端接口
// @Summary 获取用户信息
// @accept application/json
// @Produce application/json
// @Param x-token header string  true "token信息"
// @Success 200  {string} string "{"code":0,"data":{"user":{"UUID":"","ID":1,"Username":"","NickName":"","AuthorityId":"","BufferTime":86400,"exp":1621475281,"iss":"qmPlus","nbf":1620869481},"avw_balance":"9.5","ht_balance":null},"msg":"操作成功"}"
// @Router /web/user/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		response.FailWithMessage("40003", c)
		return
	}
	token := claims.(*request.CustomClaims)
	User := model.AvfUser{
		GVA_MODEL: global.GVA_MODEL{ID: token.ID},
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
	// block, err3 := client.ReadNewHeaderBlock()
	// fmt.Printf("block:%s,err:%e\n", block, err3)
	// info, err4 := client.ReadBlockInfo(block)
	// fmt.Printf("info:%s,err:%e\n", info, err4)

	balance, err2 := client.SelectBalance(User.WalletAddress)
	if err2 != nil {
		response.FailWithMessage("40006", c)
		return
	}

	res := UserInfo{
		User:       token,
		AVWBalance: balance,
	}
	fmt.Printf("balance:%s", balance)
	response.OkWithData(res, c)
}

type UserInfo struct {
	User       *request.CustomClaims `json:"user"`
	AVWBalance *big.Float            `json:"avw_balance"`
	HTBalance  *big.Float            `json:"ht_balance"`
}

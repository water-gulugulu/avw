// @File  : tools.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/15
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

package web_tools

import (
	"bytes"
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils/blockchian"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/w3liu/go-common/constant/timeformat"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const char = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const number = "0123456789"

// 通过token获取用户ID
func GetUserId(c *gin.Context) (uint, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return 0, errors.New("读取token信息失败")
	}
	token := claims.(*request.CustomClaims)

	return token.ID, nil
}

// GetRandPassword 获取随机密码
func GetRandPassword(num int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := num
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf) // E.g. "3i[g0|)z"

	return str
}
func RandChar(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		s.WriteByte(char[rand.Int63()%int64(len(char))])
	}
	return s.String()
}
func RandNumber(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		s.WriteByte(number[rand.Int63()%int64(len(number))])
	}
	return s.String()
}

var num int64

// CreateSn 生成24位订单号 前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
func CreateSn(t time.Time) string {
	s := t.Format(timeformat.Continuity)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// sup 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

// 抽奖算法
func Lottery(noSSR bool) int {
	start := 0
	var end, level int
	var list = make([]int, 0)
	if noSSR {
		list = []int{9000, 800, 200}
	} else {
		list = []int{9000, 800, 190, 10}
	}
	randNumber := rand.Intn(10000)
	for k, probability := range list {
		end += probability
		if start <= randNumber && end > randNumber {
			level = k
		}
		start = end
	}
	return level + 1
}

// 登录以后签发jwt
func TokenNext(user model.AvfUser) (string, error) {
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

var client *blockchian.ClientManage

// 循环读取哈希来改变订单状态
func LoopOrderStatus(txHash string, OrderId int) {
	if len(txHash) == 0 {
		return
	}
	var err error
	if client == nil {
		client, err = blockchian.NewClient()
		if err != nil {
			log.Printf("[%s]Failed to client RPC by Hash:%s error:%e\n", time.Now(), txHash, err)
			return
		}
	}

	defer client.CloseClient()

	Order := model.AvfOrder{
		TxHash: txHash,
	}
	if err := Order.FindByHash(global.GVA_DB); err != nil {
		log.Printf("[%s]Failed to Hash:%s query Order error:%e\n", time.Now(), txHash, err)
		return
	}

	for {
		res, err2 := client.QueryTransactionByTxHash(txHash)
		if err2 != nil {
			log.Printf("[%s]Failed to query transaction error:%e\n", time.Now(), err)
			continue
		}
		if res.Status != 1 {
			log.Printf("[%s]Failed to status not 1\n", time.Now())
			break
		}
		res.From = strings.ToUpper(res.From)
		Order.From = strings.ToUpper(Order.From)
		if res.From != Order.From {
			log.Printf("[%s]Failed to form:%s orderForm:%s\n", time.Now(), res.From, Order.From)
			break
		}
		Price := Order.Price * 100000000000000000

		if global.GVA_CONFIG.CollectionAddress.Debug == "1" {
			P := 0.001 * 100000000000000000
			Price = int64(P)
		}
		price := strconv.Itoa(int(Price))
		if price != res.Value.String() {
			log.Printf("[%s]Failed to money not same money:%v,%v \n", time.Now(), Price, res.Value)
			break
		}

		res.To = strings.ToUpper(res.To)
		Order.To = strings.ToUpper(global.GVA_CONFIG.CollectionAddress.Address)
		if res.To != Order.To {
			log.Printf("[%s]Failed to to:%s orderTo:%s\n", time.Now(), res.To, Order.To)
			break
		}

		Order = model.AvfOrder{
			GVA_MODEL: global.GVA_MODEL{
				ID:        uint(OrderId),
				UpdatedAt: time.Now(),
			},
			TxHash:   txHash,
			Status:   3,
			PayTime:  int(time.Now().Unix()),
			Block:    res.Block.String(),
			Gas:      string(res.Gas),
			GasPrice: res.GasPrice.String(),
			From:     res.From,
			To:       res.To,
		}
		if err := Order.UpdateOrder(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to update Order error:%e\n", time.Now(), err)
			break
		}

		Log := model.AvfTransactionLog{
			OrderId:    OrderId,
			Block:      res.Block.String(),
			TxHash:     txHash,
			Form:       res.From,
			To:         res.To,
			Gas:        strconv.Itoa(int(res.Gas)),
			GasPrice:   res.GasPrice.String(),
			Value:      res.Value.String(),
			Nonce:      string(res.Nonce),
			Data:       string(res.Data),
			Status:     int64(res.Status),
			CreateDate: time.Now(),
			CreateTime: time.Now().Unix(),
			Type:       1,
		}
		if err := Log.CreateLog(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to update Order error:%e,Log:%s\n", time.Now(), err, Log)
		}
		break
	}
}

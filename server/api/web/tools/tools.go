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
	Rand "crypto/rand"
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/w3liu/go-common/constant/timeformat"
	"go.uber.org/zap"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"
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

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := Rand.Int(Rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := Rand.Int(Rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
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

// 获取当天零点时间
func GetTodayZeroTimeStamp() int {
	timeStr := time.Now().Format("2006-01-02")
	// fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	return int(timeNumber)
}

func FormatFloat(num float64, decimal int) float64 {
	// 默认乘1
	d := float64(100000000000000000)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return math.Trunc(num*d) / d
	// return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}

func IntToFloat(num int) float64 {
	totalAmount := num
	numrator, _ := new(big.Float).SetPrec(uint(1024)).SetString(strconv.Itoa(totalAmount))

	denominator := big.NewFloat(1)
	denominator1, _ := numrator.Mul(numrator, denominator).Float64()
	return denominator1
}

// 截取小数位数
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f/pow10_n)*pow10_n) / pow10_n
}

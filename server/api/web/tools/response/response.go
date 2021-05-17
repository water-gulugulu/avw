package response

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR       = 7
	TOKEN_ERROR = 8
	SUCCESS     = 0
)

var language string = "Chinese"

func Result(code int, data interface{}, msg string, c *gin.Context) {
	if code == ERROR {
		lang := c.Request.Header.Get("language")
		if len(lang) != 0 {
			language = lang
		}
		if language != "Chinese" && lang != "Japan" && lang != "English" {
			language = "Chinese"
		}
		message := global.GVA_ERRCODE[language][msg]

		if len(message) != 0 {
			msg = message
		}
		fmt.Println(msg)
	}
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
func FailWithDetailedByToken(data interface{}, message string, c *gin.Context) {
	Result(TOKEN_ERROR, data, message, c)
}

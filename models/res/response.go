package res

import (
	"github.com/gin-gonic/gin"
	"goblog_server/utils"
	"net/http"
)

const (
	Success = 0
	Error   = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// ListResponse 返回元素数量 以及一个泛型的数组
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// 响应成功
func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "Success", c)
}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}
func OKWithList(list any, count int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		Count: count,
		List:  list,
	}, c)
}

// 响应失败
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}

// 失败一般不需要传回data any
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code] //在map中查错误码
	// 若能查询得到错误类型
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	// 若查询不到错误类型
	Result(Error, map[string]any{}, "未知错误", c)
}

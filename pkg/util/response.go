package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// https://juejin.cn/post/7258119695010824250
// https://www.jb51.net/jiaoben/293311yib.html

type Response struct {
	Code      int         `json:"code" example:"200"` // 响应状态码
	Data      interface{} `json:"data"`               // 响应数据
	Msg       interface{} `json:"msg"`                // 响应信息
	RequestId string      `json:"requestId"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	var msgID string
	data, ok := c.Get("msgID")
	if !ok {
		msgID = uuid.New().String()
		c.Set("msgID", msgID)
		return msgID
	}
	msgID = cast.ToString(data)
	return msgID
}

// 失败数据处理
func ErrorWithErr(c *gin.Context, code int, msg string, err error) {
	var res Response
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}
	// res.RequestId = GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(code))
}

// 失败数据处理
func Error(c *gin.Context, code int, msg interface{}) {
	var res Response
	res.Msg = msg
	// res.RequestId = GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(code))
}

// 通常成功数据处理
func OK(c *gin.Context, msg string, data interface{}) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	// res.RequestId = GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnOK())
}

// 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res Page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, msg, res)
}

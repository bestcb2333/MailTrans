package util

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 构造响应体内容的函数
func Resp(msg string, data any) gin.H {
	return gin.H{"msg": msg, "data": data}
}

// 输出错误信息
func Error(c *gin.Context, status int, msg string, err error) {
	c.AbortWithStatusJSON(status, Resp(msg, nil))
	var errstring string
	if err != nil {
		errstring = err.Error()
	}
	log.Println(msg, errstring)
}

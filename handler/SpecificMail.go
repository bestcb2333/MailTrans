package handler

import (
	"github.com/McaxDev/MailTrans/util"
	"github.com/gin-gonic/gin"
)

// 获取特定邮件的内容的handler
func SpecificMail(c *gin.Context) {

	// 建立邮箱连接
	conn, err := util.ConnectMail()
	if err != nil {
		c.AbortWithStatusJSON(500, util.Resp("无法连接到邮件服务器", err))
		return
	}
	defer conn.Logout()
}

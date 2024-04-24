package handler

import (
	"github.com/McaxDev/MailTrans/util"
	"github.com/gin-gonic/gin"
)

// 获取所有邮件的handler
func AllMail(c *gin.Context) {

	// 与邮件服务器建立连接
	conn, err := util.ConnectMail()
	if err != nil {
		util.Error(c, 500, "与邮件服务器建立连接失败", err)
		return
	}
	defer conn.Logout()
}

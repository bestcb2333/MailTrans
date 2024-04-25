package handler

import (
	"strconv"

	"github.com/McaxDev/MailTrans/util"
	"github.com/gin-gonic/gin"
)

// 获取特定邮件的内容的handler
func SpecificMail(c *gin.Context) {

	// 建立邮箱连接
	conn, err := util.ConnectMail()
	if err != nil {
		util.Error(c, 500, "无法连接到邮件服务器", err)
		return
	}
	defer conn.Logout()

	// 从查询字符串参数获取邮件ID
	uidStr := c.Query("uid")
	uidint, err := strconv.Atoi(uidStr)
	if err != nil {
		util.Error(c, 400, "查询字符串参数UID无效", err)
		return
	}

	content, err := util.GetContent(conn, uint32(uidint))
	if err != nil {
		util.Error(c, 500, "查询邮件内容失败", err)
		return
	}

	msg := <-content

	htmlBody, err := util.ExtractText(msg)
	if err != nil {
		util.Error(c, 500, "获取邮件内容文本失败", err)
		return
	}
	c.AbortWithStatusJSON(200, util.Resp("邮件内容查询成功", htmlBody))
}

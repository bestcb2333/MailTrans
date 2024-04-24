package handler

import (
	"strconv"

	"github.com/McaxDev/MailTrans/util"
	"github.com/emersion/go-imap"
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
	uid := uint32(uidint)

	seqset := new(imap.SeqSet)
	seqset.AddNum(uid)

	//获取完整的邮件体
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}
	messages := make(chan *imap.Message, 1)
	go func() {
		if err := conn.UidFetch(seqset, items, messages); err != nil {
			util.Error(c, 500, "获取邮件内容失败", err)
			return
		}
	}()

	// 读取邮件内容
	for msg := range messages {
		r := msg.GetBody(section)
		if r == nil {
			util.Error(c, 500, "邮件服务器没有发送邮件内容", err)
			return
		}
		c.AbortWithStatusJSON(200, util.Resp("邮件内容查询成功", r))
	}
}

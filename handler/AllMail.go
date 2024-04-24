package handler

import (
	"github.com/McaxDev/MailTrans/util"
	"github.com/emersion/go-imap"
	"github.com/gin-gonic/gin"
)

type Email struct {
	UID     uint32
	From    string
	Subject string
}

// 获取所有邮件的handler
func AllMail(c *gin.Context) {

	// 与邮件服务器建立连接
	conn, err := util.ConnectMail()
	if err != nil {
		util.Error(c, 500, "与邮件服务器建立连接失败", err)
		return
	}
	defer conn.Logout()

	// 获取所有邮件的UID
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	ids, err := conn.UidSearch(criteria)
	if err != nil {
		util.Error(c, 500, "从邮件服务器获取邮件UID列表失败", err)
		return
	}

	// 获取邮件的概要信息
	seqset := new(imap.SeqSet)
	seqset.AddNum(ids...)
	items := []imap.FetchItem{imap.FetchEnvelope}
	messages := make(chan *imap.Message, 10)
	go func() {
		if err := conn.UidFetch(seqset, items, messages); err != nil {
			util.Error(c, 500, "获取邮件的概要信息失败", err)
			return
		}
	}()

	// 将邮件列表变成切片
	var emails []Email
	for msg := range messages {
		if msg == nil {
			continue
		}
		email := Email{
			UID:     msg.Uid,
			From:    msg.Envelope.From[0].PersonalName,
			Subject: msg.Envelope.Subject,
		}
		emails = append(emails, email)
	}

	// 将切片返回
	c.AbortWithStatusJSON(200, util.Resp("查询成功", emails))
}

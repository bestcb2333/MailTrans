package handler

import (
	"sort"

	"github.com/McaxDev/MailTrans/config"
	"github.com/McaxDev/MailTrans/util"
	"github.com/emersion/go-imap"
	"github.com/gin-gonic/gin"
)

type Email struct {
	Time    string
	UID     uint32
	From    string
	Subject string
	Preview string
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

	// 创建检索条件变量
	criteria := imap.NewSearchCriteria()

	// 仅允许特定的主题关键词的邮件
	filter := config.Config.Filter
	for _, value := range filter {
		criteriaA := imap.NewSearchCriteria()
		criteriaA.Header.Add("Subject", value[0])
		criteriaB := imap.NewSearchCriteria()
		criteriaB.Header.Add("Subject", value[1])
		arrayCriteria := [2]*imap.SearchCriteria{criteriaA, criteriaB}
		criteria.Or = append(criteria.Or, arrayCriteria)
	}

	// 查询字符串参数 收件人
	receiver := c.Query("receiver")
	if receiver == "" {
		util.Error(c, 400, "You have to provide a receiver email.", nil)
		return
	}
	criteria.Header.Add("To", receiver)

	// 搜索满足条件的邮件UID
	ids, err := conn.UidSearch(criteria)
	if err != nil {
		util.Error(c, 500, "Failed to fetch uids of emails.", err)
		return
	}

	// 对UID进行排序
	sort.Slice(ids, func(i, j int) bool { return ids[i] > ids[j] })

	// 只截取前五个UID
	if len(ids) > 5 {
		ids = ids[:5]
	}

	// 获取邮件的概要信息
	messages, err := util.GetContent(conn, ids...)
	if err != nil {
		util.Error(c, 500, "Failed to fetch summary of email.", err)
		return
	}

	// 将邮件列表变成切片
	var emails []Email
	for msg := range messages {
		if msg == nil {
			continue
		}
		preview, err := util.ExtractText(msg)
		if err != nil {
			preview = "Failed to load detailed information of email."
		}
		email := Email{
			Time:    msg.Envelope.Date.Format("2006-01-02 15:04"),
			UID:     msg.Uid,
			From:    msg.Envelope.From[0].PersonalName,
			Subject: msg.Envelope.Subject,
			Preview: preview,
		}
		emails = append(emails, email)
	}

	// 将切片返回
	c.AbortWithStatusJSON(200, util.Resp("Fetched successfully.", emails))
}

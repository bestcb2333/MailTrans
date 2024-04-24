package util

import (
	"github.com/McaxDev/MailTrans/config"
	"github.com/emersion/go-imap/client"
)

// 连接到邮件服务器的函数
func ConnectMail() (*client.Client, error) {

	conf := config.Config

	// 连接到IMAP服务器
	c, err := client.DialTLS(conf.MailSrvIp+":"+conf.MailSrvPort, nil)
	if err != nil {
		return nil, err
	}

	// 登录到IMAP服务器
	if err := c.Login(conf.MailAccount, conf.MailPassword); err != nil {
		return nil, err
	}

	// 选择收件箱
	if _, err := c.Select("INBOX", false); err != nil {
		return nil, err
	}

	return c, nil
}

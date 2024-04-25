package util

import (
	"crypto/tls"

	"github.com/McaxDev/MailTrans/config"
	"github.com/emersion/go-imap/client"
)

// 连接到邮件服务器的函数
func ConnectMail() (c *client.Client, err error) {
	conf := config.Config

	// 先使用普通连接
	c, err = client.Dial(conf.MailSrvIp + ":" + conf.MailSrvPort)
	if err != nil {
		return nil, err
	}

	// 尝试升级到TLS
	tlsConfig := &tls.Config{
		ServerName: conf.MailSrvIp, // 通常设置为服务器的域名
	}
	if err = c.StartTLS(tlsConfig); err != nil {
		c.Logout() // 出错时确保释放资源
		return nil, err
	}

	// 登录到IMAP服务器
	if err = c.Login(conf.MailAccount, conf.MailPassword); err != nil {
		c.Logout() // 登录失败也确保释放资源
		return nil, err
	}

	// 选择收件箱
	if _, err = c.Select("INBOX", false); err != nil {
		c.Logout()
		return nil, err
	}

	return c, nil
}

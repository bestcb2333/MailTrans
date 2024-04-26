package config

// 配置文件模板
type ConfigTemplate struct {
	Port         string
	MailSrvIp    string
	MailSrvPort  string
	MailAccount  string
	MailPassword string
	Filter       [][2]string
}

// 配置文件变量
var Config = ConfigTemplate{
	Port:         "8910",
	MailSrvIp:    "127.0.0.1",
	MailSrvPort:  "143",
	MailAccount:  "",
	MailPassword: "",
	Filter: [][2]string{
		{"临时", "更新"},
		{"代码", "同户设备"},
	},
}

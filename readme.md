## 配置文件使用说明
### Port 这个程序运行的端口
    "Port": "5678"
### MailSrvIp 邮箱服务器d的地址
    "MailSrvIp": "mail.ckui.net"
### MailSrvPort 邮箱服务器IMAP端口
    "MailSrvPort": "143"
### MailAccount 邮箱账号
    "MailAccount": "code@ckui.net"
### 邮箱密码
    "MailPassword": "Code888123"
### 主题筛选器
* 整个大括号里可以有无限个小括号，每个小括号里只能有两个关键词。
* 小括号里的两个关键词之间是或者关系，大括号里的小括号之间是并且关系。
* 假如你有三封邮件，主题分别是 Google如何更新设备 Google你的用户代码 Google你的设备名称 私人邮件，
  * 如果你只想显示 Google如何更新设备 和 Google你的用户代码，这时 更新 和 代码 有一个词就可以了，是或者关系，即
     ```
     "Filter": [
	    ["更新", "代码"]
     ]
	 ```
  * 如果你只想显示 Google如何更新设备 和 Google你的设备名称，这里 Google 和 设备 是两个主题共有词，可以设置并且关系，即
     ```
	 "Filter": [
	    ["Google", "任意内容"],
		["设备", "任意内容"]
	 ]
	 ```


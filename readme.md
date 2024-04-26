## 后端软件使用方法
* 在服务器上创建一个文件夹用于存放后端程序，并进入这个文件夹
```
sudo mkdir /srv/MailTrans/
cd /srv/MailTrans/
```
* 下载后端程序文件
```
sudo wget https://cloud.mcax.cn/s/tmp/download?path=%2F&files=MailTrans&downloadStartSecret=x5ya27tpdzp
```
* 为程序授予可执行权限
```
sudo chmod +x MailTrans
```
* 执行程序，以生成配置文件
```
sudo ./MailTrans
```
* 修改配置文件内容
```
sudo vim config.json
```
* 修改完成再次尝试启动
```
sudo ./MailTrans
```
* 配置systemd配置文件 
```
sudo cat << EOF | sudo tee /etc/systemd/system/MailTrans.service
[Unit]
Description=Mail Transformer
After=network.target

[Service]
ExecStart=/srv/MailTrans
Restart=always
User=root

[Install]
WantedBy=multi-user.target
EOF
```
* 重载systemd配置文件
```
sudo systemctl daemon-reload
```
* 通过systemd启动后端程序
```
sudo systemctl start MailTrans
```
* 设置后端程序为开机自动启动
```
sudo systemctl enable MailTrans
```
## 配置文件使用说明
### Port 这个程序运行的端口
    "Port": "5678"
### MailSrvIp 邮箱服务器d的地址
    "MailSrvIp": "mail.youremail.com"
### MailSrvPort 邮箱服务器IMAP端口
    "MailSrvPort": "143"
### MailAccount 邮箱账号
    "MailAccount": "yourEmail@gmail.com"
### 邮箱密码
    "MailPassword": "yourPassWord"
### 主题筛选器
* 整个大括号里可以有无限个小括号，每个小括号里只能有两个关键词。
* 小括号里的两个关键词之间是或者关系，大括号里的小括号之间是并且关系。
* 假如你有四封邮件，主题分别是 **Google如何更新设备** **Google你的用户代码** **Google你的设备名称** **私人邮件**，
  * 如果你只想显示 **Google如何更新设备** 和 **Google你的用户代码**，这时 **更新** 和 **代码** 有一个词就可以了，是或者关系，即
     ```
     "Filter": [
	    ["更新", "代码"]
     ]
	 ```
  * 如果你只想显示 **Google如何更新设备** 和 **Google你的设备名称**，这里 **Google** 和 **设备** 是两个主题共有词，可以设置并且关系，即
     ```
	 "Filter": [
	    ["Google", "任意内容"],
		["设备", "任意内容"]
	 ]
	 ```


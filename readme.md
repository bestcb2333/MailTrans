this code is a customer's order and I open-sourced it. I made $80 from this source code.


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
* 任意一封邮件，只要标题里面有字符串数组里面任意一个字符串，就通过
	 ```


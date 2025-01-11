# HostLoc 每日刷分脚本
LOC 每日签到脚本，支持多账号、Telegram 推送, 开启 TLS 指纹伪装.

## 获取二进制文件

### (1) 从 Github Release 中下载
```shell
# 以 amd64 架构的系统为例
mkdir HostLoc_CheckIn
cd HostLoc_CheckIn
wget https://github.com/LordPenguin666/Hostloc-daily-checkin-tls/releases/download/v1.0.0/hostloc-check-in-linux-amd64.tar.gz
tar xvf hostloc-check-in-linux-amd64.tar.gz
```

### (2) 编译 (需要拥有 go 环境)
```shell
git clone git@github.com:LordPenguin666/Hostloc-daily-checkin-tls.git
cd Hostloc-daily-checkin-tls
make default
```

## 使用方法

1. 复制配置文件 example.json `cp example.json config.json`；
2. 修改配置文件 `vim config.json`；
3. (可选) 你也可以通过 `./hostloc -c /path/to/your/config` 指定配置文件路径；
4. 使用 `./hostloc` 运行脚本。

## 后台运行
可以使用 `tmux` 或 `screen` 等工具后台运行。

```bash
tmux new -s hostloc
./hostloc
```

你也可以使用 systemd 等工具将其作为服务运行。

```bash 
# /usr/lib/systemd/system/hostloc.service
[Unit]
Description=HostLoc CheckIn Service
After=network.target
Wants=network.target

[Service]
WorkingDirectory=/path/to/your/hostloc
ExecStart=/path/to/your/hostloc/hostloc
Restart=on-abnormal
RestartSec=5s
KillMode=mixed

[Install]
WantedBy=multi-user.target
```

## 配置文件说明
- 可以设置启动程序时立即开始签到
- 可以通过配置文件指定多个帐号，也可以配置 Telegram 推送

### 定时任务配置
```json
{
  "time": "0 5 * * *" // 每天 5 点执行
}
```

### 立即开始签到
```json
{
  "startup": true
}
```

### 多帐号配置

```json
{
  "accounts": [
    {"username": "第一个帐号名", "password": "密码"},
    {"username": "第二个帐号名", "password": "密码"}, 
    {"username": "cuper", "password": "114514"}
  ]
}
```

### Telegram 推送

```json
{
  "telegram": {
    "enable": true, // 开启推送
    "token": "这里填写 bot token",
    "chat_id": "这里填写对话 id"
  }
}
```


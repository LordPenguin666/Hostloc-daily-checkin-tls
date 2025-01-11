package config

type Config struct {
	Time     string     `json:"time"`
	Startup  bool       `json:"startup"`
	Telegram Telegram   `json:"telegram"`
	Accounts []Accounts `json:"accounts"`
}

type Telegram struct {
	Enable bool   `json:"enable"`
	Token  string `json:"token"`
	ChatID int64  `json:"chat_id"`
}

type Accounts struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

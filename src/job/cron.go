package job

import (
	"HostLoc-Daily-CheckIn/src/request"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"time"
)

func (j *Job) Start() {
	if j.config.Startup {
		j.checkIn()
	}

	s, err := gocron.NewScheduler()
	if err != nil {
		j.logger.Fatal(err.Error())
	}

	if _, err = s.NewJob(gocron.CronJob(j.config.Time, false), gocron.NewTask(j.checkIn)); err != nil {
		j.logger.Fatal(err.Error())
	}

	s.Start()
}

func (j *Job) checkIn() {
	j.logger.Info("签到任务开始!")
	var messages string

	for n, account := range j.config.Accounts {
		req := request.New()
		j.logger.Info("正在进行第 " + strconv.Itoa(n+1) + " 个账号签到")

		// 获取 formhash
		r, err := req.Member()
		if err != nil {
			j.logger.Error("获取 [formhash] 失败! Error: " + err.Error())
			continue
		}

		formhash := j.getHashForm(r)
		if formhash == "" {
			j.logger.Error("正则匹配 [formhash] 失败! ")
			continue
		}
		j.logger.Info("正则匹配到 [formhash] 成功, Value=" + formhash)

		r, err = req.MainPage()
		if err != nil {
			j.logger.Error("初始化 [cookies] 失败! Error: " + err.Error())
			continue
		}
		j.logger.Info("初始化 [cookies] 成功! ")
		req.UpdateCookies(r.Cookies())
		time.Sleep(time.Second)

		r, err = req.Login(&account, formhash)
		if err != nil {
			j.logger.Error(fmt.Sprintf("用户 [%v] 登陆失败\v. Error: %v", account.Username, err.Error()))
			continue
		}

		j.logger.Info(fmt.Sprintf("用户 [%v] 准备执行空间访问任务.", account.Username))
		req.UpdateCookies(r.Cookies())
		time.Sleep(time.Second)

		r, err = req.CheckCoin()
		if err != nil {
			j.logger.Error(fmt.Sprintf("用户 [%v] 获取金币数据失败! Error: %v", account.Username, err.Error()))
			continue
		}

		coins := j.getCoin(r)
		j.logger.Info(fmt.Sprintf("用户 [%v] 当前金钱 [%v].", account.Username, coins))
		time.Sleep(time.Second)

		uids := j.randomUID(20, 70000)
		for k, uid := range uids {
			j.logger.Info(fmt.Sprintf("用户 [%v] 正在进行第 %v 次空间访问.", account.Username, k+1))
			r, err = req.Space(uid)
			if err != nil {
				j.logger.Error(fmt.Sprintf("用户 [%v] 访问空间失败! Error: %v", account.Username, err.Error()))
				time.Sleep(time.Second * 10)
				continue
			}

			if strings.Contains(r.String(), account.Username) == true {
				j.logger.Info(fmt.Sprintf("用户 [%v] 空间访问成功 [Space UID: %v]", account.Username, uid))
			} else {
				j.logger.Error(fmt.Sprintf("用户 [%v] 空间访问失败 [Space UID: %v]", account.Username, uid))
			}

			time.Sleep(time.Second * 5)
		}

		r, err = req.CheckCoin()
		if err != nil {
			j.logger.Error(fmt.Sprintf("用户 [%v] 获取金币数据失败! Error: %v", account.Username, err.Error()))
			continue
		}

		newCoins := j.getCoin(r)
		msg := fmt.Sprintf("用户 [%v] 金钱: %v -> %v", account.Username, coins, newCoins)
		j.logger.Info(msg)
		messages += msg + "\n"

		time.Sleep(time.Second * 10)
	}

	if !j.config.Telegram.Enable {
		return
	}

	if j.config.Telegram.API == "" || j.config.Telegram.ChatID == 0 {
		j.logger.Info("Telegram 配置不全, 取消推送")
		return
	}

	if messages == "" {
		messages = "没有账号进行了签到"
	}

	messages = "[LOC 签到小助手]\n\n" + messages
	fmt.Println(messages)

	bot, err := tgbotapi.NewBotAPI(j.config.Telegram.API)
	if err != nil {
		j.logger.Error("Telegram 推送通知失败! Error: " + err.Error())
		return
	}

	text := tgbotapi.NewMessage(j.config.Telegram.ChatID, messages)
	if _, err = bot.Send(text); err != nil {
		j.logger.Error("Telegram 推送通知失败! Error: " + err.Error())
	} else {
		j.logger.Info("Telegram 推送通知成功!")
	}
}

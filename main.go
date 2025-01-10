package main

import (
	"HostLoc-Daily-CheckIn/src/config"
	"HostLoc-Daily-CheckIn/src/job"
	"HostLoc-Daily-CheckIn/src/logger"
	"flag"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "c", "config.json", "配置文件路径, 默认当前路径下的 config.json")
	flag.Parse()
}

func main() {
	log := logger.New()
	conf := config.ReadConfig(fileName, log)

	j := job.NewJob(
		job.WithLogger(log),
		job.WithConfig(conf),
	)

	j.Start()

	select {}

}

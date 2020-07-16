package initialize

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

// Cron 定时任务
var Cron *cron.Cron

// InitCron 初始化cron
func InitCron() {
	fmt.Println("cron start...")
	c := cron.New(cron.WithChain(cron.Recover(cron.PrintfLogger(log.New(os.Stdout, "", log.LstdFlags)))))
	Cron = c
	c.AddFunc("*/1 * * * *", cronjob)
	c.Start()
}

func cronjob() {
	fmt.Println("cron test...")
}

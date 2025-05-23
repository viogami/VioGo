package cron

import (
	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New(cron.WithSeconds())
	
	c.AddFunc("0 */5 * * * *", clearChatHistory) // 每5分钟执行一次

	c.Start()
	defer c.Stop()

	select {}
}

package cron

import (
	"github.com/robfig/cron/v3"
)

type CronTask struct {
}

func init() {
	aiCron := NewAICronTask()

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */5 * * * *", aiCron.ClearChatHistory) // 每5分钟执行一次

	c.Start()
	defer c.Stop()

	select {}
}

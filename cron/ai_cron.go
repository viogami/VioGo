package cron

import (
	"log/slog"
	"time"

	"github.com/viogami/viogo/AI/deepseek"
)

type AICronTask struct {
	CronTask
}

func (c *AICronTask) ClearChatHistory() {
	s := deepseek.NewDeepSeekService()
	s.ClearContext()

	now := time.Now()
	slog.Info("AICronTask:ClearChatHistory执行成功!", "time", now.Format("2006-01-02 15:04:05"))

}

func NewAICronTask() *AICronTask {
	return &AICronTask{}
}

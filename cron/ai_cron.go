package cron

import (
	"github.com/viogami/viogo/AI/deepseek"
)

func clearChatHistory() {
	s := deepseek.NewDeepSeekService()
	s.ClearContext()
}

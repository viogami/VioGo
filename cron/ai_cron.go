package cron

import (
	"github.com/viogami/viogo/ai/deepseek"
)

func clearChatHistory() {
	s := deepseek.NewDeepSeekService()
	s.ClearContext()
}

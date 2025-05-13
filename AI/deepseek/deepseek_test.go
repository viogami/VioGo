package deepseek_test

import (
	"os"
	"testing"

	"github.com/viogami/viogo/AI/deepseek"
)

func TestInvokeDeepSeekAPI(t *testing.T) {
	os.Setenv("DeepSeekAPIKey", "sk-xxxx")

	service := deepseek.NewDeepSeekService()

	text := "什么是超讽刺和元讽刺？"

	resp := service.InvokeDeepSeekAPI(text)
	t.Logf("Response: %s", resp)
}

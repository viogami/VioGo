package deepseek_test

import (
	"os"
	"testing"

	"github.com/viogami/viogo/ai/deepseek"
)

func TestInvokeDeepSeekAPI(t *testing.T) {
	os.Setenv("DeepSeekAPIKey", "sk-xxx")
	// 运行test需要修改config的init函数读取配置文件的位置
	// 	data, err := os.ReadFile("../../config.yaml")

	service := deepseek.NewDeepSeekService()
	
	text := "什么是超讽刺和元讽刺？"

	resp := service.InvokeDeepSeekAPI(text)
	t.Logf("Response: %s", resp)
}

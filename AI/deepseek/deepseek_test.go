package deepseek_test

import (
	"testing"

	"github.com/viogami/viogo/AI/deepseek"
)

func TestInvokeDeepSeekAPI(t *testing.T) {
	// Create a new DeepSeekService instance
	service := deepseek.NewDeepSeekService()

	// Define the input text
	text := "Hello, how are you?"

	// Call the InvokeDeepSeekAPI method
	resp := service.InvokeDeepSeekAPI(text)
	t.Logf("Response: %s", resp)
}

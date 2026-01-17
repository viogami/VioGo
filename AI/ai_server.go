package ai

import (
	"github.com/viogami/viogo/ai/deepseek"
	"github.com/viogami/viogo/ai/openai"
)

type AIServer struct {
	GptServer      *openai.ChatGPTService
	DeepSeekServer *deepseek.DeepSeekService
}

func (s *AIServer) ProcessMessage(message string) string {
	return s.GptServer.InvokeChatGPTAPI(message)
}

func (s *AIServer) ProcessSharpReviews(message string) string {
	return s.DeepSeekServer.InvokeDeepSeekAPI(message)
}

func NewAIServer() *AIServer {
	s := new(AIServer)
	s.GptServer = openai.GetInstance()
	s.DeepSeekServer = deepseek.GetInstance()
	return s
}

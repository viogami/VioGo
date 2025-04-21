package AI

import (
	"github.com/viogami/viogo/AI/deepseek"
	openai "github.com/viogami/viogo/AI/openai"
)

type AIServer struct {
	GptServer      *openai.ChatGPTService
	DeepSeekServer *deepseek.DeepSeekService
}

func (s *AIServer) ProcessMessage(message string) string {
	return s.GptServer.InvokeChatGPTAPI(message)
}

func(s *AIServer) ProcessSharpReviews(message string) string {
	return s.DeepSeekServer.InvokeDeepSeekAPI(message)
}

func NewAIServer() *AIServer {
	s := new(AIServer)
	s.GptServer = openai.NewChatGPTService()
	s.DeepSeekServer = deepseek.NewDeepSeekService()
	return s
}

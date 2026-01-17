package openai

import (
	"context"
	"log/slog"
	"os"
	"sync"

	goopenai "github.com/sashabaranov/go-openai"
	config "github.com/viogami/viogo/conf"
)

var (
	instance *ChatGPTService
	once     sync.Once
)

func GetInstance() *ChatGPTService {
	once.Do(func() {
		instance = NewChatGPTService()
	})
	return instance
}

type ChatGPTService struct {
	Role             string
	Character        string
	characterSetting string

	client *goopenai.Client
}

func NewChatGPTService() *ChatGPTService {
	URL_PROXY := config.AppConfig.AIConfig.ChatGPTUrlProxy
	APIKey := os.Getenv("ChatGPTAPIKey")

	s := new(ChatGPTService)
	s.Role = goopenai.ChatMessageRoleUser
	s.Character = "vio"
	s.characterSetting = gpt_preset[s.Character]

	conf := goopenai.DefaultConfig(APIKey)
	conf.BaseURL = URL_PROXY

	s.client = goopenai.NewClientWithConfig(conf)

	return s
}

func (s *ChatGPTService) SetCharacter(character string) {
	s.Character = character
	s.characterSetting = gpt_preset[character]
}

func (s *ChatGPTService) InvokeChatGPTAPI(text string) string {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		goopenai.ChatCompletionRequest{
			Model: goopenai.GPT4o20241120,
			Messages: []goopenai.ChatCompletionMessage{
				{
					Role:    goopenai.ChatMessageRoleSystem,
					Content: s.characterSetting,
				},
				{
					Role:    s.Role,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		slog.Error("Error calling ChatGPT API", "error", err)
		Resp := "AI调用失败了😥 error:\n" + err.Error()
		return Resp
	}
	return resp.Choices[0].Message.Content
}

func (s *ChatGPTService) InvokeChatGPTAPIWithRole(text string, role string) string {
	s.Role = role
	return s.InvokeChatGPTAPI(text)
}

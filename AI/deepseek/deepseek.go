package deepseek

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	config "github.com/viogami/viogo/conf"
)

type DeepSeekService struct {
	Client *http.Client

	Model    string
	Messages []message
}

type DeepSeekPayLoad struct {
	Messages         []message `json:"messages"`
	Model            string    `json:"model"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	PresencePenalty  float64   `json:"presence_penalty"`
	MaxTokens        int       `json:"max_tokens"`
	ResponseFormat   format    `json:"response_format"`
	Stop             string    `json:"stop"`
	Stream           bool      `json:"stream"`
	StreamOptions    []string  `json:"stream_options"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	Tools            []string  `json:"tools"`
	ToolChoice       string    `json:"tool_choice"`
	LogProbs         bool      `json:"logprobs"`
	TopLogProbs      []string  `json:"top_logprobs"`
}

type message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
type format struct {
	Type string `json:"type"`
}

func (s *DeepSeekService) InvokeDeepSeekAPI(text string) string {
	url := config.AppConfig.AIConfig.DeepSeekUrl
	key := os.Getenv("DeepSeekAPIKey")

	s.Messages = append(s.Messages, message{
		Content: text,
		Role:    "user",
	})
	// 只保留最后 20 条消息
	s.trimContext(20)

	payload := DeepSeekPayLoad{
		Messages:         s.Messages,
		Model:            s.Model,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		MaxTokens:        2048,
		ResponseFormat: format{
			Type: "text",
		},
		Stop:          "",
		Stream:        false,
		StreamOptions: nil,
		Temperature:   1,
		TopP:          1,
		Tools:         nil,
		ToolChoice:    "none",
		LogProbs:      false,
		TopLogProbs:   nil,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return string(err.Error())
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		return string(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+key)

	res, err := s.Client.Do(req)
	if err != nil {
		return string(err.Error())
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return string(err.Error())
	}
	// 将 AI 响应添加到上下文
	responseText := string(responseBody)
	s.Messages = append(s.Messages, message{
		Content: responseText,
		Role:    "assistant",
	})

	return responseText
}

func (s *DeepSeekService) setPrompt(prompt string) {
	s.Messages = append(s.Messages, message{
		Content: prompt,
		Role:    "system",
	})
}

func (s *DeepSeekService) trimContext(n int) {
	// 保留第一条系统消息和最后n条消息
	if len(s.Messages) > n {
		s.Messages = append(s.Messages[:1], s.Messages[len(s.Messages)-n:]...)
	}
}

func (s *DeepSeekService) ClearContext() {
	s.Messages = []message{}
}

func NewDeepSeekService() *DeepSeekService {
	s := new(DeepSeekService)
	s.Model = "deepseek-chat"

	s.Client = &http.Client{}
	s.Messages = []message{}
	s.setPrompt("你是一个喜欢锐评的贴吧老哥，喜欢用讽刺的语气来表达观点，语言简明干练，一针见血，诙谐幽默又不失风度。")
	return s
}

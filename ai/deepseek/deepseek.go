package deepseek

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	config "github.com/viogami/viogo/conf"
)

var (
	instance *DeepSeekService
	once     sync.Once
)

func GetInstance() *DeepSeekService {
	once.Do(func() {
		instance = NewDeepSeekService()
	})
	return instance
}

const MAX_MSG_HISTORY = 20

type DeepSeekService struct {
	Client *http.Client

	Model    string
	Messages []message
}

func (s *DeepSeekService) InvokeDeepSeekAPI(text string) string {
	url := config.AppConfig.AIConfig.DeepSeekUrl
	key := os.Getenv("DeepSeekAPIKey")

	s.Messages = append(s.Messages, message{
		Content: text,
		Role:    "user",
	})
	// 只保留最后 20 条消息
	s.trimContext(MAX_MSG_HISTORY)

	payload := newDeepSeekPayLoad()
	payload.Messages = s.Messages
	payload.Model = s.Model

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		return err.Error()
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
		return err.Error()
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Sprintf("请求失败，状态码：%d，响应：%s", res.StatusCode, string(responseBody))
	}

	// 解析响应
	encodeRes := DeepSeekResponse{}
	if err := json.Unmarshal(responseBody, &encodeRes); err != nil {
		return fmt.Sprintf("解析响应失败：%s", err.Error())
	}
	err = s.checkResponse(encodeRes)
	if err != nil {
		return err.Error()
	}
	// 将 AI 响应添加到上下文
	responseText := encodeRes.Choices[0].Message.Content
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

func (s *DeepSeekService) checkResponse(encodeRes DeepSeekResponse) error {
	if len(encodeRes.Choices) == 0 {
		return fmt.Errorf("没有返回内容")
	}
	if encodeRes.Choices[0].FinishReason != "stop" {
		return fmt.Errorf("请求失败，结束原因：%s", encodeRes.Choices[0].FinishReason)
	}
	if encodeRes.Choices[0].Message.Role != "assistant" {
		return fmt.Errorf("请求失败，角色错误：%s", encodeRes.Choices[0].Message.Role)
	}
	if encodeRes.Choices[0].Message.Content == "" {
		return fmt.Errorf("请求失败，内容为空")
	}
	return nil
}

func (s *DeepSeekService) ClearContext() {
	newMsgs := make([]message, 0)
	for _, msg := range s.Messages {
		msg.Role = "system"
		newMsgs = append(newMsgs, msg)
	}
	s.Messages = newMsgs
}

func NewDeepSeekService() *DeepSeekService {
	s := new(DeepSeekService)
	s.Model = "deepseek-chat"
	s.Client = &http.Client{}
	s.Messages = []message{}
	s.setPrompt("你是一个喜欢锐评的贴吧老哥，喜欢用讽刺的语气来表达观点，语言简明干练，一针见血，诙谐幽默又不失风度。")
	return s
}

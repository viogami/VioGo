package deepseek

type DeepSeekPayLoad struct {
	Messages         []message `json:"messages"`
	Model            string    `json:"model"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	PresencePenalty  float64   `json:"presence_penalty"`
	MaxTokens        int       `json:"max_tokens"`
	ResponseFormat   format    `json:"response_format"`
	Stop             []string  `json:"stop"`
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

func newDeepSeekPayLoad() *DeepSeekPayLoad {
	return &DeepSeekPayLoad{
		Messages:         []message{},
		Model:            "deepseek-chat",
		MaxTokens:        2048,
		ResponseFormat: format{
			Type: "text",
		},
		Temperature:   1,
		TopP:          1,
		ToolChoice:    "none",
	}
}

type DeepSeekResponse struct {
	Id                string   `json:"id"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Object            string   `json:"object"`
	Choices           []choice `json:"choices"`
	Usage             usage    `json:"usage"`
}
type choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Message      message `json:"message"`
}
type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

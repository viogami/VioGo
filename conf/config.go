package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var AppConfig Config

type Config struct {
	Services Services `yaml:"Services"`
	AIConfig AIconfig `yaml:"AIconfig"`
}

type Services struct {
	RedisEnabled bool `yaml:"redis_enabled"`
}

type AIconfig struct {
	ChatGPTUrlProxy string `yaml:"chatgpt_url_proxy"`
	DeepSeekUrl     string `yaml:"deepseek_url"`
}

// 初始化配置
func init() {
	// 读取配置文件
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		slog.Error("读取配置文件失败", "error", err)
	}

	// 解析 YAML
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		slog.Error("解析配置文件失败", "error", err)
	}
}

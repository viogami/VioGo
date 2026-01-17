package cqCode

import (
	"fmt"
	"regexp"
	"strings"
)

type CQmsg struct {
	CQcodes []CQCode
	Text    string
}

// 判断是否at我
func (cq *CQmsg) IsAtme(selfId int64) bool {
	CQcodes := cq.CQcodes
	for _, CQcode := range CQcodes {
		if CQcode.Type == "at" && CQcode.Data["qq"] == fmt.Sprintf("%d", selfId) {
			return true
		}
	}
	return false
}

func ParseCQmsg(input string) CQmsg {
	re := regexp.MustCompile(`\[CQ:([^,]+)(?:,([^=]+)=([^,]+))?\]`)

	matches := re.FindAllStringSubmatch(input, -1)

	result := CQmsg{}

	// 初始化 Text 为原始 input
	result.Text = input

	// 处理每个CQ码句段
	for _, match := range matches {
		cqCode := CQCode{
			Type: match[1],
			Data: make(map[string]any),
		}
		if match[2] != "" && match[3] != "" {
			// 如果有参数，将参数添加到 map 中
			cqCode.Data[match[2]] = match[3]
		}
		result.CQcodes = append(result.CQcodes, cqCode)
		// 替换掉当前CQ码句段
		result.Text = strings.Replace(result.Text, match[0], "", 1)
	}
	// 去除文本中的多余空格
	result.Text = strings.TrimSpace(result.Text)

	return result
}

type CQCode struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

// 生成CQ码字符串
func (cq *CQCode) GenerateCQCode() string {
	cqCode := fmt.Sprintf("[CQ:%s", cq.Type)

	for key, value := range cq.Data {
		cqCode += fmt.Sprintf(",%s=%s", key, value)
	}
	return cqCode + "]"
}

func NewCQCode(cqType string, data map[string]any) CQCode {
	return CQCode{
		Type: cqType,
		Data: data,
	}
}

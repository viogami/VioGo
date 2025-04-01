package gocq

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/viogami/Gobot-vio/utils"
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
			Data: make(map[string]interface{}),
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
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// 生成CQ码字符串
func (cq *CQCode) GenerateCQCode() string {
	cqCode := fmt.Sprintf("[CQ:%s", cq.Type)

	for key, value := range cq.Data {
		cqCode += fmt.Sprintf(",%s=%s", key, value)
	}
	return cqCode + "]"
}

func NewCQCode(cqType string, data map[string]interface{}) CQCode {
	return CQCode{
		Type: cqType,
		Data: data,
	}
}

// 获得猎杀枪声的CQ码
func GetCQCode_HuntSound(input string) string {
	sound := utils.HuntSound{
		Name:     "",
		Distance: "",
	}
	parts := strings.Split(input, " ")
	if len(parts) == 2 {
		sound.Name = parts[1]
	}
	if len(parts) == 3 {
		sound.Name = parts[1]
		sound.Distance = parts[2]
	}

	return "[CQ:record,file=" + utils.GetHuntSound(sound) + "]"
}

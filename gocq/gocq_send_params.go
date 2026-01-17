package gocq

import "github.com/viogami/viogo/gocq/cq-code"

type SendMsgParams struct {
	MessageType string `json:"message_type"` // 消息类型, 支持 private、group , 分别对应私聊、群组, 如不传入, 则根据传入的 *_id 参数判断
	UserID      int64  `json:"user_id"`      // 对方 QQ 号 ( 消息类型为 private 时需要 )
	GroupID     int64  `json:"group_id"`     // 群号 ( 消息类型为 group 时需要 )
	Message     string `json:"message"`      // 要发送的内容
	AutoEscape  bool   `json:"auto_escape"`  // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效
}

func (params SendMsgParams) toMap() map[string]any {
	return map[string]any{
		"message_type": params.MessageType,
		"user_id":      params.UserID,
		"group_id":     params.GroupID,
		"message":      params.Message,
		"auto_escape":  params.AutoEscape,
	}
}

type SendGroupForwardMsgParams struct {
	GroupID int64           `json:"group_id"` // 群号
	Message []cqCode.CQCode `json:"messages"` // 消息列表
}

func (params SendGroupForwardMsgParams) toMap() map[string]any {
	return map[string]any{
		"group_id": params.GroupID,
		"messages": params.Message,
	}
}

type SendPrivateForwardMsgParams struct {
	UserID  int64           `json:"user_id"`  // 对方 QQ 号
	Message []cqCode.CQCode `json:"messages"` // 消息列表
}

func (params SendPrivateForwardMsgParams) toMap() map[string]any {
	return map[string]any{
		"user_id":  params.UserID,
		"messages": params.Message,
	}
}

type SendSetuMsgParams struct {
	Tags []string `json:"tags"` // 色图标签
	R18  int      `json:"r18"`  // 是否 R18
	Num  int      `json:"num"`  // 色图数量
}

type SendSetGroupBanParams struct {
	GroupID  int64  `json:"group_id"` // 群号
	UserID   int64  `json:"user_id"`  // 对方 QQ 号
	Duration uint32 `json:"duration"` // 禁言时长, 单位秒, 0 为取消禁言,默认30*60
}

func (params SendSetGroupBanParams) toMap() map[string]any {
	return map[string]any{
		"group_id": params.GroupID,
		"user_id":  params.UserID,
		"duration": params.Duration,
	}
}

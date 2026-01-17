package cqEvent

import (
	"encoding/json"
	"log"
)

// 下面的Message结构体为array形式，并未使用
// 当前使用的是string形式，修改请到gocq的config文件中改变上报属性
type Message struct {
	// Type string      `json:"type"`
	// Data any `json:"data"`
}
type AtData struct {
	QQ int64 `json:"qq"`
}
type TextData struct {
	Text string `json:"text"`
}

type Sender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

// 私聊消息
type PrivateMessage struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageID   int32  `json:"message_id"`
	UserID      int64  `json:"user_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int32  `json:"font"`
	Sender      Sender `json:"sender"`
	TargetID    int64  `json:"target_id"`
	TempSource  int    `json:"temp_source"`
}

type Anonymous struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

// 群聊消息
type GroupMessage struct {
	Time        int64     `json:"time"`
	SelfID      int64     `json:"self_id"`
	PostType    string    `json:"post_type"`
	MessageType string    `json:"message_type"`
	SubType     string    `json:"sub_type"`
	MessageID   int32     `json:"message_id"`
	UserID      int64     `json:"user_id"`
	Message     string    `json:"message"`
	RawMessage  string    `json:"raw_message"`
	Font        int32     `json:"font"`
	Sender      Sender    `json:"sender"`
	GroupID     int64     `json:"group_id"`
	Anonymous   Anonymous `json:"anonymous"`
}

func Get_msg_info(p []byte, msgType string) any {
	var privateMessage PrivateMessage
	var groupMessage GroupMessage
	switch msgType {
	case "private":
		err := json.Unmarshal(p, &privateMessage)
		if err != nil {
			log.Println("Error parsing JSON to privateMessage:", err)
		}
		return privateMessage
	case "group":
		err := json.Unmarshal(p, &groupMessage)
		if err != nil {
			log.Println("Error parsing JSON to groupMessage:", err)
		}
		return groupMessage
	}
	return nil
}

// 私聊消息快速操作
// @params
// reply	message	要回复的内容	不回复
// auto_escape	boolean	消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 reply 字段是字符串时有效	不转义
func PrivateMsgFastOperate(reply string, auto_escape bool) map[string]any {
	sendMessage := map[string]any{
		"reply":       reply,
		"auto_escape": auto_escape,
	}
	return sendMessage
}

// 群消息快速操作
// @params
// reply	message	要回复的内容	不回复
// auto_escape	boolean	消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 reply 字段是字符串时有效	不转义
// at_sender	boolean	是否要在回复开头 at 发送者 ( 自动添加 ) , 发送者是匿名用户时无效	at 发送者
// delete	boolean	撤回该条消息	不撤回
// kick	boolean	把发送者踢出群组 ( 需要登录号权限足够 ) , 不拒绝此人后续加群请求, 发送者是匿名用户时无效	不踢出
// ban	boolean	禁言该消息发送者, 对匿名用户也有效	不禁言
// ban_duration	number	若要执行禁言操作时的禁言时长	30 分钟
func GroupMsgFastOperate(reply string, auto_escape bool, at_sender bool, delete bool, kick bool, ban bool, ban_duration int) map[string]any {
	sendMessage := map[string]any{
		"reply":        reply,
		"auto_escape":  auto_escape,
		"at_sender":    at_sender,
		"delete":       delete,
		"kick":         kick,
		"ban":          ban,
		"ban_duration": ban_duration,
	}
	return sendMessage
}

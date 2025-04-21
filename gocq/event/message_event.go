package event

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/viogami/viogo/gocq/command"
	"github.com/viogami/viogo/gocq/cqCode"
	"github.com/viogami/viogo/utils"
)

type MessageEvent struct {
	Event

	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageID   int32  `json:"message_id"`
	UserID      int64  `json:"user_id"`
	GroupID     int64  `json:"group_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	Sender      Sender `json:"sender"`
}

type Sender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

func (m *MessageEvent) LogInfo() {
	slog.Info("MessageEvent",
		"message_type", m.MessageType,
		"sub_type", m.SubType,
		"message_id", m.MessageID,
		"user_id", m.UserID,
		"group_id", m.GroupID,
		"message", m.Message,
		// "raw_message", m.RawMessage,
		// "font", m.Font
		"sender_id", m.Sender.UserID,
	)
}

func (m *MessageEvent) Handle() {
	key := fmt.Sprintf("message_%d", m.UserID)
	pushToRedis(key, m.Message, 1) // 将消息推送到redis

	cqmsg := cqCode.ParseCQmsg(m.Message)
	f := m.parseCommand(cqmsg)
	if f == nil {
		slog.Info("MessageEvent", "接受到普通群聊消息", m.Message)
		return
	}
	params := command.CommandParams{
		MessageId:   m.MessageID,
		MessageType: m.MessageType,
		Message:     m.Message,
		GroupId:     m.GroupID,
		UserId:      m.UserID,

		SetuParams: command.SetuParams{
			Tags: utils.ReadTags(cqmsg.Text),
		},
	}

	f.Execute(params)
}

func (m *MessageEvent) parseCommand(cqmsg cqCode.CQmsg) command.Command {
	cmdStr := cqmsg.Text

	// 遍历 CommandMap，检查是否有以 cmdStr 开头的命令
	var matchedCommand command.Command
	for key, cmd := range command.CommandMap {
		if strings.HasPrefix(cmdStr, key) {
			matchedCommand = cmd
			break
		}
	}

	t_info := command.COMMAND_INFO_CMD_TYPE
	t_private := command.COMMAND_TYPE_PRIVATE
	t_group := command.COMMAND_TYPE_GROUP
	t_all := command.COMMAND_TYPE_ALL

	// 判断是否是私聊消息
	if m.MessageType == t_private {
		if matchedCommand == nil {
			return command.CommandMap["/chat"]
		}
		if matchedCommand.GetInfo(t_info) == t_private || matchedCommand.GetInfo(t_info) == t_all {
			return matchedCommand
		}
		return command.CommandMap["/chat"]
	}

	// 判断是否是群聊消息
	if m.MessageType == t_group && cqmsg.IsAtme(m.SelfID) {
		if matchedCommand == nil {
			return command.CommandMap["/chat"]
		}
		if matchedCommand.GetInfo(t_info) == t_group || matchedCommand.GetInfo(t_info) == t_all {
			return matchedCommand
		}
	}
	return nil
}

func NewMessageEvent(p []byte) (*MessageEvent, error) {
	messageEvent := new(MessageEvent)
	err := json.Unmarshal(p, &messageEvent)
	if err != nil {
		slog.Error("Error parsing JSON to MessageEvent:", "err", err)
		return nil, err
	}
	return messageEvent, nil
}

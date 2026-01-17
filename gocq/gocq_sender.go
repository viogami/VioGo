package gocq

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/viogami/viogo/gocq/cq-code"
)

var maxWaitingTime = 10 * time.Second // 响应超时时间

type GocqSender struct {
	writeMutex sync.Mutex // 添加互斥锁，ws无并发安全
	conn       *websocket.Conn
}

func (s *GocqSender) sendToGocq(action string, params map[string]any) (resp RHttpResq, err error) {
	s.writeMutex.Lock()

	// 生成唯一echo值
	echoValue := fmt.Sprintf("%s:%d", action, time.Now().UnixNano())

	// 创建消息
	messageSend := map[string]any{
		"action": action,
		"params": params,
		"echo":   echoValue, // 添加echo字段
	}

	// 创建一个channel用于接收响应
	responseChan := make(chan RHttpResq, 1)
	Instance.ResponseMap.Store(echoValue, responseChan)

	// 发送请求
	err = s.conn.WriteJSON(messageSend)
	s.writeMutex.Unlock() // 发送后立即释放锁，允许其他请求发送

	if err != nil {
		Instance.ResponseMap.Delete(echoValue)
		return RHttpResq{}, err
	}

	// 等待响应
	select {
	case resp := <-responseChan:
		if resp.Status != "ok" {
			Instance.ResponseMap.Delete(echoValue)
			return resp, fmt.Errorf("api请求失败,状态码: %s, 错误信息: %s", resp.Status, resp.Msg)
		}
		slog.Info("收到api响应", "response", resp)
		return resp, nil
	case <-time.After(maxWaitingTime): // 超时时间
		Instance.ResponseMap.Delete(echoValue)
		return RHttpResq{}, fmt.Errorf("等待响应超时")
	}
}

func (s *GocqSender) SendMsg(params SendMsgParams) {
	action := "send_msg"

	if params.MessageType == "group" {
		cq := cqCode.CQCode{
			Type: "at",
			Data: map[string]any{
				"qq": fmt.Sprintf("%d", params.UserID),
			},
		}
		params.Message = cq.GenerateCQCode() + params.Message
	}
	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送消息失败", "error", err)
		return
	}
	slog.Info("发送消息成功", "message", params.Message, "userId", params.UserID, "groupId", params.GroupID)
}

func (s *GocqSender) SendGroupForwardMsg(params SendGroupForwardMsgParams) {
	action := "send_group_forward_msg"

	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送群聊合并消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SendPrivateForwardMsg(params SendPrivateForwardMsgParams) {
	action := "send_private_forward_msg"

	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("发送私聊合并消息失败", "error", err)
		return
	}
}

func (s *GocqSender) SetGroupBan(params SendSetGroupBanParams) {
	action := "set_group_ban"

	_, err := s.sendToGocq(action, params.toMap())
	if err != nil {
		slog.Error("设置群禁言失败", "error", err)
		return
	}
}

func (s *GocqSender) GetMsg(msgid int32) map[string]any {
	action := "get_msg"
	params := map[string]any{
		"message_id": msgid,
	}

	resp, err := s.sendToGocq(action, params)
	if err != nil {
		slog.Error("获取消息失败", "error", err)
		return nil
	}

	return resp.Data
}

func NewGocqSender(conn *websocket.Conn) *GocqSender {
	return &GocqSender{
		writeMutex: sync.Mutex{}, // 初始化互斥锁
		conn:       conn,
	}
}

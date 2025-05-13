package command

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	config "github.com/viogami/viogo/conf"
	"github.com/viogami/viogo/gocq"
	"github.com/viogami/viogo/utils"
)

type redisRecord struct {
	MessageId  int32       `json:"message_id"`
	OperatorId json.Number `json:"operator_id"`
	UserId     json.Number `json:"user_id"`
}

type cmdGetRecall struct {
	BaseCmd
}

func (c *cmdGetRecall) Execute(params CommandParams) {
	client := gocq.Instance.RedisClient
	sender := gocq.Instance.Sender
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     "",
		AutoEscape:  false,
	}

	// 从 Redis 中获取上一次撤回的消息 ID
	key := fmt.Sprintf("group_recall_%d", params.GroupId)
	record, err := client.RPop(context.Background(), key).Result()
	if err != nil {
		slog.Warn("获取上一次撤回的消息 ID 失败", "error", err)
		msgParams.Message = "找不到撤回的消息!可能已经过期了~"
		sender.SendMsg(msgParams)
		return
	}
	redisData := new(redisRecord)
	if err := json.Unmarshal([]byte(record), &redisData); err != nil {
		slog.Error("解析上一次撤回的消息 ID 失败", "error", err)
		return
	}
	// 获取消息 ID 和消息内容
	messageId := redisData.MessageId
	userId := redisData.UserId
	operatorId := redisData.OperatorId

	resp := sender.GetMsg(messageId)
	time := utils.Time2Str(resp["time"])
	if resp["message"] == nil {
		msgParams.Message = "没有收到撤回的消息~"
		sender.SendMsg(msgParams)
		return
	}

	msgParams.Message = fmt.Sprintf("撤回时间:%s\n发送者:%s\n撤回者:%s\n消息内容:%s", time, userId, operatorId, resp["message"])
	sender.SendMsg(msgParams)
	slog.Info("执行指令:撤回了什么")
}

func newCmdGetRecall() *cmdGetRecall {
	if config.AppConfig.Services.RedisEnabled == false {
		return nil
	}
	inst := new(cmdGetRecall)
	inst.Command = "撤回了什么"
	inst.Description = "撤回了什么"
	inst.CmdType = COMMAND_TYPE_GROUP
	return inst
}

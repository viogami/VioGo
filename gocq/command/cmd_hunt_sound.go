package command

import (
	"log/slog"

	"github.com/viogami/viogo/gocq"
	"github.com/viogami/viogo/gocq/cqCode"
	"github.com/viogami/viogo/utils"
)

type cmdHuntSound struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdHuntSound) Execute(params CommandParams) {
	sender := gocq.Instance.Sender

	hs := utils.NewRandHuntSound() // 随机枪声，固定5m
	reply := c.cqReply(hs.Sound)
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("执行指令:打一枪听听", "reply", reply)

	sender.SendMsg(msgParams)
}

func (c *cmdHuntSound) cqReply(soundUrl string) string {
	ret := cqCode.CQCode{
		Type: "record",
		Data: map[string]any{
			"file": soundUrl,
		},
	}
	return ret.GenerateCQCode()
}

func (c *cmdHuntSound) GetInfo(index int) string {
	switch index {
	case COMMAND_INFO_COMMAND:
		return c.Command
	case COMMAND_INFO_DESCRIPTION:
		return c.Description
	case COMMAND_INFO_CMD_TYPE:
		return c.CmdType
	}
	return ""
}

func newCmdHuntSound() *cmdHuntSound {
	return &cmdHuntSound{
		Command:     "打一枪听听",
		Description: "随机猎杀对决枪声，固定5米",
		CmdType:     COMMAND_TYPE_ALL,
	}
}

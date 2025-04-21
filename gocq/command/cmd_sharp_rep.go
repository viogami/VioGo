package command

import (
	"log/slog"

	"github.com/viogami/viogo/AI"
	"github.com/viogami/viogo/gocq"
)

type cmdSharpRep struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdSharpRep) Execute(params CommandParams) {
	sender := gocq.Instance.Sender
	reply := AI.NewAIServer().ProcessSharpReviews(params.Message)
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	sender.SendMsg(msgParams)
	slog.Info("执行指令：锐评一下")
}

func (c *cmdSharpRep) GetInfo(index int) string {
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

func newCmdSharpRep() *cmdSharpRep {
	return &cmdSharpRep{
		Command:     "锐评一下",
		Description: "锐评一下xxx",
		CmdType:     COMMAND_TYPE_ALL,
	}
}

package command

import (
	"log/slog"

	"github.com/viogami/viogo/AI"
	"github.com/viogami/viogo/gocq"
)

type cmdChat struct {
	BaseCmd
}

func (c *cmdChat) Execute(params CommandParams) {
	sender := gocq.Instance.Sender

	reply := AI.NewAIServer().ProcessMessage(params.Message)
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		UserID:      params.UserId,
		GroupID:     params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("调用ai执行指令:/chat")

	sender.SendMsg(msgParams)
}

func newCmdChat() *cmdChat {
	inst := new(cmdChat)
	inst.Command = "/chat"
	inst.Description = "聊天指令"
	inst.CmdType = COMMAND_TYPE_ALL
	return inst
}

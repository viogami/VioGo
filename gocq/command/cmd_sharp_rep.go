package command

import (
	"log/slog"

	AI "github.com/viogami/viogo/ai"
	"github.com/viogami/viogo/gocq"
)

type cmdSharpRep struct {
	BaseCmd
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

func newCmdSharpRep() *cmdSharpRep {
	inst := new(cmdSharpRep)
	inst.Command = "锐评一下"
	inst.Description = "锐评一下xxx"
	inst.CmdType = COMMAND_TYPE_ALL
	return inst
}

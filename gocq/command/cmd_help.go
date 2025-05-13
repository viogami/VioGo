package command

import (
	"log/slog"

	"github.com/viogami/viogo/gocq"
)

type cmdHelp struct {
	BaseCmd
}

func (c *cmdHelp) Execute(params CommandParams) {
	sender := gocq.Instance.Sender

	reply := ""
	if params.MessageType == "group" {
		reply = c.groupReply()
	} else if params.MessageType == "private" {
		reply = c.privateReply()
	}
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		UserID:      params.UserId,
		GroupID:     params.GroupId,
		Message:     reply,
		AutoEscape:  false,
	}
	slog.Info("执行指令:help", "reply", reply)

	sender.SendMsg(msgParams)
}

func (c *cmdHelp) privateReply() string {
	reply := "指令列表:\n"
	for _, v := range CommandList[1:] {
		if v.GetInfo(2) == "private" || v.GetInfo(2) == "all" {
			reply += "[" + v.GetInfo(0) + "]:" + v.GetInfo(1) + "\n"
		}
	}
	return reply
}

func (c *cmdHelp) groupReply() string {
	reply := "指令列表:"
	for _, v := range CommandList[1:] {
		if v.GetInfo(2) == "group" || v.GetInfo(2) == "all" {
			reply += "\n" + "[" + v.GetInfo(0) + "]:" + v.GetInfo(1)
		}
	}
	return reply

}

func newCmdHelp() *cmdHelp {
	inst := new(cmdHelp)
	inst.Command = "help"
	inst.Description = "指令列表"
	inst.CmdType = COMMAND_TYPE_ALL
	return inst
}

package command

import (
	"log/slog"

	"github.com/viogami/viogo/gocq"
)

type cmdHelp struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
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

func (c *cmdHelp) GetInfo(index int) string {
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
	return &cmdHelp{
		Command:     "help",
		Description: "查看指令列表",
		CmdType:     COMMAND_TYPE_ALL,
	}
}

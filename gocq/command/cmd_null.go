package command

import (
	"log/slog"
)

type cmdNull struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdNull) Execute(params CommandParams) {
	slog.Info("空指令")
}

func (c *cmdNull) GetInfo(index int) string {
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

func newCmdNull() *cmdNull {
	return &cmdNull{
		Command:     "",
		Description: "空指令",
		CmdType:     COMMAND_TYPE_ALL,
	}
}

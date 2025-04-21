package command

import (
	"crypto/rand"
	"log/slog"
	"math/big"

	"github.com/viogami/viogo/gocq"
)

var maxDuration = 600 // 最大禁言时间
type cmdBanLottery struct {
	Command     string // 指令名称
	Description string // 指令描述
	CmdType     string // 指令类型
}

func (c *cmdBanLottery) Execute(params CommandParams) {
	sender := gocq.Instance.Sender

	duration, _ := rand.Int(rand.Reader, big.NewInt(int64(maxDuration)))
	banParams := gocq.SendSetGroupBanParams{
		GroupID:  params.GroupId,
		UserID:   params.UserId,
		Duration: uint32(duration.Int64()),
	}
	sender.SetGroupBan(banParams)

	reply := "恭喜你获得了" + duration.String() + "秒的禁言时间！"
	msgParams := gocq.SendMsgParams{
		MessageType: params.MessageType,
		GroupID:     params.GroupId,
		UserID:      params.UserId,
		Message:     reply,
		AutoEscape:  false,
	}
	sender.SendMsg(msgParams)
	slog.Info("执行指令:禁言抽奖")
}

func (c *cmdBanLottery) GetInfo(index int) string {
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

func newCmdBanLottery() *cmdBanLottery {
	return &cmdBanLottery{
		Command:     "禁言抽奖",
		Description: "禁言抽奖0~600秒",
		CmdType:     COMMAND_TYPE_GROUP,
	}
}

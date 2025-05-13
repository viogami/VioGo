package command

import (
	"crypto/rand"
	"log/slog"
	"math/big"

	"github.com/viogami/viogo/gocq"
)

var maxDuration = 600 // 最大禁言时间
type cmdBanLottery struct {
	BaseCmd
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

func newCmdBanLottery() *cmdBanLottery {
	inst := new(cmdBanLottery)
	inst.Command = "禁言抽奖"
	inst.Description = "禁言抽奖"
	inst.CmdType = COMMAND_TYPE_GROUP
	return inst
}

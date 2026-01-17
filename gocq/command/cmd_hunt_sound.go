package command

import (
	"log/slog"

	"github.com/viogami/viogo/gocq"
	"github.com/viogami/viogo/gocq/cq-code"
	"github.com/viogami/viogo/utils"
)

type cmdHuntSound struct {
	BaseCmd
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

func newCmdHuntSound() *cmdHuntSound {
	inst := new(cmdHuntSound)
	inst.Command = "打一枪听听"
	inst.Description = "随机猎杀对决枪声，固定5米"
	inst.CmdType = COMMAND_TYPE_ALL
	return inst
}

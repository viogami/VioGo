package command

import (
	"fmt"
	"log/slog"

	"github.com/viogami/viogo/gocq"
	"github.com/viogami/viogo/gocq/cq-code"
	"github.com/viogami/viogo/utils"
)

type cmdSetu struct {
	BaseCmd
}

func (c *cmdSetu) Execute(params CommandParams) {
	reply := c.getSetuReply(gocq.SendSetuMsgParams{
		Tags: params.Tags,
		R18:  0,
		Num:  1,
	})
	slog.Info("执行指令:来份涩图", "reply", reply)
	sender := gocq.Instance.Sender

	if params.MessageType == "private" {
		msgParams := gocq.SendPrivateForwardMsgParams{
			UserID:  params.UserId,
			Message: reply,
		}
		sender.SendPrivateForwardMsg(msgParams)
		return
	}
	// 如果是群消息，使用 SendGroupForwardMsg 发送
	msgParams := gocq.SendGroupForwardMsgParams{
		GroupID: params.GroupId,
		Message: reply,
	}
	sender.SendGroupForwardMsg(msgParams)
}

func (c *cmdSetu) getSetuReply(params gocq.SendSetuMsgParams) []cqCode.CQCode {
	reply := []cqCode.CQCode{
		{
			Type: "node",
			Data: map[string]any{
				"name": "LV",
				"uin":  "1524175162",
				"content": []cqCode.CQCode{
					{
						Type: "text",
						Data: map[string]any{
							"text": fmt.Sprintf("tags:%s", params.Tags),
						},
					},
				},
			},
		},
	}
	content := []cqCode.CQCode{}
	setuInfo := utils.GetSetu(params.Tags, params.R18, params.Num)
	if setuInfo.Error != "" {
		slog.Error("随机色图api调用出错", "error", setuInfo.Error)
		return nil
	}
	if len(setuInfo.Data) == 0 {
		slog.Error("随机色图api调用出错:tag搜索不到,返回数据为空")
		return nil
	}
	for _, data := range setuInfo.Data {
		content = append(content, cqCode.NewCQCode("image", map[string]any{
			"file": data.Urls.Regular,
			"url":  data.Urls.Regular,
		}))
	}
	reply[0].Data["content"] = append(reply[0].Data["content"].([]cqCode.CQCode), content...)
	return reply
}

func newCmdSetu() *cmdSetu {
	inst := new(cmdSetu)
	inst.Command = "来份涩图"
	inst.Description = "随机涩图,指令后可接tag,用逗号分隔"
	inst.CmdType = COMMAND_TYPE_ALL
	return inst
}

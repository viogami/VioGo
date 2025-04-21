package command

type Command interface {
	Execute(params CommandParams) // 执行指令
	GetInfo(index int) string     // 获取指令信息
}

type CommandParams struct {
	MessageId   int32
	MessageType string
	Message     string
	GroupId     int64
	UserId      int64

	SetuParams
}

type SetuParams struct {
	Tags []string // 标签
}

var CommandList = []Command{
	newCmdNull(),
	newCmdHelp(),
	newCmdChat(),
	newCmdSetu(),
	newCmdSetuR18(),
	newCmdHuntSound(),
	newCmdBanLottery(),
	newCmdGetRecall(),
	newCmdSharpRep(),
}
var CommandMap = map[string]Command{
	"help":    newCmdHelp(),
	"/chat":   newCmdChat(),
	"来份涩图":    newCmdSetu(),
	"来份r18涩图": newCmdSetuR18(),
	"打一枪听听":   newCmdHuntSound(),
	"禁言抽奖":    newCmdBanLottery(),
	"撤回了什么":   newCmdGetRecall(),
	"锐评一下":      newCmdSharpRep(),
}

const (
	COMMAND_TYPE_ALL     = "all"     // 所有类型
	COMMAND_TYPE_GROUP   = "group"   // 群聊类型
	COMMAND_TYPE_PRIVATE = "private" // 私聊类型
	COMMAND_TYPE_ADMIN   = "admin"   // 管理员类型
	COMMAND_TYPE_OWNER   = "owner"   // 群主类型
)

const (
	COMMAND_INFO_COMMAND     = 0 // 指令名称
	COMMAND_INFO_DESCRIPTION = 1 // 指令描述
	COMMAND_INFO_CMD_TYPE    = 2 // 指令类型
)

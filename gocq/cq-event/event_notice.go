package cqEvent

import (
	"encoding/json"
	"log"
)

// 以下在手表协议无法使用
// 群内戳一戳（双击头像）
// 群红包运气王提示
// 群成员荣誉变更提示
// 群成员名片更新

// ----------------------------------------------------
// 私聊消息撤回
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	friend_recall	通知类型
// user_id	int64		好友 QQ 号
// message_id	int64		被撤回的消息 ID
type PrivateRecallNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
	MessageID  int64  `json:"message_id"`
}

// 群消息撤回
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	group_recall	通知类型
// group_id	int64		群号
// user_id	int64		消息发送者 QQ 号
// operator_id	int64		操作者 QQ 号
// message_id	int64		被撤回的消息 ID
type GroupRecallNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
	MessageID  int64  `json:"message_id"`
}

// 群成员增加
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	group_increase	通知类型
// sub_type	string	approve、invite	事件子类型, 分别表示管理员已同意入群、管理员邀请入群
// group_id	int64	-	群号
// operator_id	int64	-	操作者 QQ 号
// user_id	int64	-	加入者 QQ 号
type GroupIncreaseNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
}

// 群成员减少
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	group_decrease	通知类型
// sub_type	string	leave、kick、kick_me	事件子类型, 分别表示主动退群、成员被踢、登录号被踢
// group_id	int64	-	群号
// operator_id	int64	-	操作者 QQ 号 ( 如果是主动退群, 则和 user_id 相同 )
// user_id	int64	-	离开者 QQ 号
type GroupDecreaseNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
}

// 群管理员变动
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	group_admin	通知类型
// sub_type	string	set、unset	事件子类型, 分别表示设置和取消管理员
// group_id	int64	-	群号
// user_id	int64	-	管理员 QQ 号
type GroupAdminChangeNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
}

// 群文件上传
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	group_upload	通知类型
// group_id	int64	-	群号
// user_id	int64	-	发送者 QQ 号
// file	object	-	文件信息
type file1 struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}
type GroupUploadFileNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	File       file1  `json:"file"`
}

// 群禁言
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	group_ban	通知类型
// sub_type	string	ban、lift_ban	事件子类型, 分别表示禁言、解除禁言
// group_id	int64	-	群号
// operator_id	int64	-	操作者 QQ 号
// user_id	int64	-	被禁言 QQ 号 (为全员禁言时为0)
// duration	int64	-	禁言时长, 单位秒 (为全员禁言时为-1)
type GroupBanNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
	Duration   int64  `json:"duration"`
}

// 好友添加
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	friend_add	通知类型
// user_id	int64	-	新添加好友 QQ 号
type FriendAddNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
}

// 好友戳一戳（双击头像）
// time	int64		时间
// * self_id	int64		BOT QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	notify	消息类型
// sub_type	string	poke	提示类型
// sender_id	int64		发送者 QQ 号
// user_id	int64		发送者 QQ 号
// target_id	int64		被戳者 QQ 号
type FriendPokeNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	SenderID   int64  `json:"sender_id"`
	UserID     int64  `json:"user_id"`
	TargetID   int64  `json:"target_id"`
}

// 群成员头衔变更
// time	int64		时间
// * self_id	int64		BOT QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	notify	消息类型
// sub_type	string	title	提示类型
// group_id	int64		群号
// user_id	int64		变更头衔的用户 QQ 号
// title	string		获得的新头衔
type GroupTitleChangeNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	Title      string `json:"title"`
}

// 接收到离线文件
// time	int64		时间
// * self_id	int64		BOT QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	offline_file	消息类型
// user_id	int64		发送者id
// file	object		文件数据
type file2 struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Url  string `json:"url"`
}
type OfflineFileNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
	File       file2  `json:"file"`
}

// 其他客户端在线状态变更
// post_type	string 参考	notice	上报类型
// * notice_type	string 参考	client_status	消息类型
// client	Device*		客户端信息
// online	bool		当前是否在线
type Device struct {
	AppID      int64  `json:"app_id"`
	DeviceName string `json:"device_name"`
	DeviceKind string `json:"device_kind"`
}
type OtherDeviceChangeNotice struct {
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	Client     Device `json:"client"`
	Online     bool   `json:"online"`
}

// 精华消息变更
// time	int64		时间
// * self_id	int64		BOT QQ 号
// * post_type	string 参考	notice	上报类型
// * notice_type	string 参考	essence	消息类型
// sub_type	string	add,delete	添加为add,移出为delete
// group_id	int64		群号
// sender_id	int64		消息发送者ID
// operator_id	int64		操作者ID
// message_id	int32		消息ID
type EssenceChangeNotice struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	SenderID   int64  `json:"sender_id"`
	OperatorID int64  `json:"operator_id"`
	MessageID  int32  `json:"message_id"`
}

// 获取通知信息
func Get_notice_info(p []byte, NoticeType string) any {
	var (
		group_upload_notice   GroupUploadFileNotice
		group_admin_notice    GroupAdminChangeNotice
		group_increase_notice GroupIncreaseNotice
		group_decrease_notice GroupDecreaseNotice
		group_ban_notice      GroupBanNotice
		group_recall_notice   GroupRecallNotice
		essence_notice        EssenceChangeNotice

		friend_add_notice     FriendAddNotice
		private_recall_notice PrivateRecallNotice

		offline_file_notice  OfflineFileNotice
		client_status_notice OtherDeviceChangeNotice

		titleChange_notice GroupTitleChangeNotice
		friend_poke_notice FriendPokeNotice
	)
	switch NoticeType {
	// -----------群通知-----------
	case "group_upload":
		err := json.Unmarshal(p, &group_upload_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_upload_notice:", err)
		}
		return group_upload_notice
	case "group_admin":
		err := json.Unmarshal(p, &group_admin_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_admin_notice:", err)
		}
		return group_admin_notice
	case "group_increase":
		err := json.Unmarshal(p, &group_increase_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_increase_notice:", err)
		}
		return group_increase_notice
	case "group_decrease":
		err := json.Unmarshal(p, &group_decrease_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_decrease_notice:", err)
		}
		return group_decrease_notice
	case "group_ban":
		err := json.Unmarshal(p, &group_ban_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_ban_notice:", err)
		}
		return group_ban_notice
	case "group_recall":
		err := json.Unmarshal(p, &group_recall_notice)
		if err != nil {
			log.Println("Error parsing JSON to group_recall_notice:", err)
		}
		return group_recall_notice
	case "essence":
		err := json.Unmarshal(p, &essence_notice)
		if err != nil {
			log.Println("Error parsing JSON to essence_notice:", err)
		}
		return essence_notice
	// -----------好友通知-----------
	case "friend_add":
		err := json.Unmarshal(p, &friend_add_notice)
		if err != nil {
			log.Println("Error parsing JSON to friend_add_notice:", err)
		}
		return friend_add_notice
	case "friend_recall":
		err := json.Unmarshal(p, &private_recall_notice)
		if err != nil {
			log.Println("Error parsing JSON to friend_recall_notice:", err)
		}
		return private_recall_notice
	// -----------设备通知-----------
	case "offline_file":
		err := json.Unmarshal(p, &offline_file_notice)
		if err != nil {
			log.Println("Error parsing JSON to offline_file_notice:", err)
		}
		return offline_file_notice
	case "client_status":
		err := json.Unmarshal(p, &client_status_notice)
		if err != nil {
			log.Println("Error parsing JSON to client_status_notice:", err)
		}
		return client_status_notice
	// -----------系统通知-----------
	case "notify":
		type notify_notice struct {
			SubType string `json:"sub_type"`
		}
		var notify notify_notice
		err := json.Unmarshal(p, &notify)
		if err != nil {
			log.Println("Error parsing JSON to notify_notice:", err)
		}
		// 以下为notify的子类型
		switch notify.SubType {
		case "title":
			err := json.Unmarshal(p, &titleChange_notice)
			if err != nil {
				log.Println("Error parsing JSON to titleChange_notice:", err)
			}
			return titleChange_notice
		case "poke":
			err := json.Unmarshal(p, &friend_poke_notice)
			if err != nil {
				log.Println("Error parsing JSON to friend_poke_notice:", err)
			}
			return friend_poke_notice
		}
	}
	return nil
}

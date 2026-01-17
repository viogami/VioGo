package cqEvent

import (
	"encoding/json"
	"log"
)

// StatusStatistics - 统计信息
// packet_received	uint64	收包数
// packet_sent	uint64	发包数
// packet_lost	uint64	丢包数
// message_received	uint64	消息接收数
// message_sent	uint64	消息发送数
// disconnect_times	uint32	连接断开次数
// lost_times	uint32	连接丢失次数
// last_message_time	int64	最后一次消息时间
type StatusStatistics struct {
	PacketReceived  uint64 `json:"packet_received"`
	PacketSent      uint64 `json:"packet_sent"`
	PacketLost      uint64 `json:"packet_lost"`
	MessageReceived uint64 `json:"message_received"`
	MessageSent     uint64 `json:"message_sent"`
	DisconnectTimes uint32 `json:"disconnect_times"`
	LostTimes       uint32 `json:"lost_times"`
	LastMessageTime int64  `json:"last_message_time"`
}

// Status	-	应用程序状态
// app_initialized	bool	程序是否初始化完毕
// app_enabled	bool	程序是否可用
// plugins_good	bool	插件正常(可能为 null)
// app_good	bool	程序正常
// online	bool	是否在线
// stat	Status_Statistics	统计信息
type Status struct {
	AppInitialized bool             `json:"app_initialized"`
	AppEnabled     bool             `json:"app_enabled"`
	PluginsGood    bool             `json:"plugins_good"`
	AppGood        bool             `json:"app_good"`
	Online         bool             `json:"online"`
	Stat           StatusStatistics `json:"stat"`
}

// 心跳包
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	meta_event	上报类型
// * meta_event_type	string 参考	heartbeat	元事件类型
// status	Status 参考	-	应用程序状态
// interval	int64	-	距离上一次心跳包的时间(单位是毫秒
type Heartbeat struct {
	Time          int64  `json:"time"`
	SelfID        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	Status        Status `json:"status"`
	Interval      int64  `json:"interval"`
}

// Lifecycle - 生命周期
// time	int64	-	事件发生的时间戳
// * self_id	int64	-	收到事件的机器人 QQ 号
// * post_type	string 参考	meta_event	上报类型
// * meta_event_type	string 参考	lifecycle	元事件类型
// sub_type	string	enable, disable, connect	子类型
type Lifecycle struct {
	Time          int64  `json:"time"`
	SelfID        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
}

func Get_meta_event(p []byte, metaType string) any {
	var heartbeat Heartbeat
	var lifecycle Lifecycle

	switch metaType {
	case "heartbeat":
		err := json.Unmarshal(p, &heartbeat)
		if err != nil {
			log.Println("Error parsing JSON to heartbeat:", err)
		}
		return heartbeat
	case "lifecycle":
		err := json.Unmarshal(p, &lifecycle)
		if err != nil {
			log.Println("Error parsing JSON to lifecycle:", err)
		}
		return lifecycle
	}
	return nil
}

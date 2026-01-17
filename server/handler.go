package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	AI "github.com/viogami/viogo/ai"
	"github.com/viogami/viogo/gocq"
	"github.com/viogami/viogo/gocq/event"
)

// gptMsgHandle 处理POST请求
func gptMsgHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 获取表单数据
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}
		// 读取请求体
		postmsg := r.Form.Get("usermsg")
		if postmsg != "" {
			slog.Info("POST request,the usermsg:", "postmsg", postmsg)
		} else {
			http.Error(w, "Error:Don`t find the key:usermsg in the POST,maybe it`s a nil", http.StatusBadRequest)
		}
		// 调用ChatGPT API
		gptResponse := AI.NewAIServer().ProcessMessage(postmsg)
		fmt.Fprintln(w, gptResponse)
	} else {
		http.Error(w, "Error: wrong HTTP method:"+r.Method+",required POST.", http.StatusMethodNotAllowed)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	gocq.Instance.Sender = gocq.NewGocqSender(conn)

	// 启动一个 goroutine 读取 WebSocket 消息并发送到队列
	go func() {
		defer close(gocq.Instance.MsgQueue) // 确保在退出时关闭消息队列
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				slog.Error("Error reading message from WebSocket:", "error", err)
				return
			}
			// 检查是否为 API 响应
			var rawMsg gocq.RHttpResq
			if err := json.Unmarshal(p, &rawMsg); err == nil {
				// 分发到对应的响应 channel
				echo := rawMsg.Echo
				if ch, ok := gocq.Instance.ResponseMap.LoadAndDelete(echo); ok {
					ch.(chan gocq.RHttpResq) <- rawMsg
					continue
				}
			}

			gocq.Instance.MsgQueue <- p
		}
	}()

	// 主循环处理消息队列中的消息
	for msg := range gocq.Instance.MsgQueue {
		// 判断是否为事件消息
		if event.IsEvent(msg) {
			e, err := event.ParseEvent(msg)
			if err != nil {
				slog.Warn("Failed to parse event", "error", err)
				continue
			}
			go e.LogInfo()
			go e.Handle()
		} else {
			// 处理非事件消息（如 API 响应）
			slog.Warn("Received non-event message", "message", string(msg))
		}
	}
}

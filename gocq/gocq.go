package gocq

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type GocqServer struct {
	conn       *websocket.Conn
	writeMutex sync.Mutex // 添加互斥锁，ws无并发安全
}

var GocqInstance *GocqServer

func NewGocqServer(conn *websocket.Conn) *GocqServer {
	// 初始化GocqData实例
	if GocqDataInstance == nil {
		GocqDataInstance = NewGocqData()
	}
	return &GocqServer{
		conn:       conn,
		writeMutex: sync.Mutex{}, // 初始化互斥锁
	}
}

func (g *GocqServer) SendToGocq(action string, params map[string]any) error {
	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()

	messageSend := map[string]interface{}{
		"action": action,
		"params": params,
	}
	return g.conn.WriteJSON(messageSend)
}

func (g *GocqServer) SendMessageWithEcho(action string, params map[string]any, echo string) error {
	messageSend := map[string]interface{}{
		"action": action,
		"params": params,
		"echo":   echo,
	}
	return g.conn.WriteJSON(messageSend)
}

func (g *GocqServer) IsConnected() bool {
	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()

	// 尝试发送一个ping消息来检查连接
	err := g.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second))
	return err == nil
}

func (g *GocqServer) Reconnect(url string) error {
	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()
	// 关闭旧连接
	if g.conn != nil {
		g.conn.Close()
	}
	// 建立新连接
	newConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	g.conn = newConn
	return nil
}

package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/wbw295/zap-config"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/url"
	"time"
)

type AgentCommand string

const (
	writeWait = 20 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024

	AgentStatus   AgentCommand = "status"
)

type Message struct {
	Command AgentCommand `json:"Command"`
	Data    string `json:"Data"`
}

func main() {

	// 发起 websocket 连接请求
	//url := url.URL{Scheme: "ws", Host: *flag.String("addr", "localhost:8080", "http service address"), Path: "/ws"}
	url := url.URL{Scheme: "ws", Host: *flag.String("addr", "localhost:15752", "http service address"), Path: "/ws"}
	log.Infof("connecting to %s", url.String())
	header := http.Header{
		"client-id": []string{"f276ba17-72c3-4566-adae-90cc4ca96315"},
		"fingerprint": []string{"abcsdjhcsadhc"},
	}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), header)
	if err != nil {
		log.Fatal("dial:", zap.Error(err))
	}
	defer conn.Close()
	conn.SetPingHandler(func(message string) error {
		log.Debug("received ping")
		err := conn.WriteControl(websocket.PongMessage, []byte(message), time.Now().Add(writeWait))
		if err == websocket.ErrReadLimit {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})
	conn.SetPongHandler(func(appData string) error {
		log.Debug("received pong")
		err := conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Error("connection set read deadline occur error", zap.Error(err))
		}
		return err
	})

	// 状态上报
	info := &Message{
		Command: AgentStatus,
		Data: "ok",
	}
	data, err := json.Marshal(info)
	if err != nil {
		log.Fatal("json marshal error:", zap.Error(err))
	}
	if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Fatal("client register write data occur error", zap.Error(err))
	}
	go func() {
		for {
			if err = conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Fatal("client send heartbeat message occur error", zap.Error(err))
			}
			time.Sleep(pingPeriod)
		}
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("read occur error", zap.Error(err))
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatal("IsUnexpectedCloseError error", zap.Error(err))
			}
		}
		log.Infof("Received msg: %+v", string(data))
	}
}


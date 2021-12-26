package driver

import (
	"RouterForGuild/dto/tencent_guild"
	wss "github.com/gorilla/websocket"
	"sync"
	"time"
)


type GuildClient struct {
	version         int
	seq             uint64
	conn            *wss.Conn
	mu              sync.Mutex
	msgQueue        messageChan
	session         tencent_guild.Session
	user            tencent_guild.WebSocketUser
	closeChan       closeErrorChan
	heartBeatTicker *time.Ticker
}

type messageChan chan []byte
type closeErrorChan chan error

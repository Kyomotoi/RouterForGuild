package driver

import (
	"RouterForGuild/dto/tencent_guild/websocket"
	"RouterForGuild/global"
	"encoding/json"
	wss "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
	"sync/atomic"
	"time"
)

// Connect 建立ws链接
func (c *GuildClient) Connect() error {
	var err error
	if c.session.URL == "" {
		log.Error("")
		return err
	}

	log.Infof("开始尝试连接到WebSocket服务器：%v", c.session.URL)

	var res *http.Response
	c.conn, res, err = wss.DefaultDialer.Dial(c.session.URL, nil)
	if err != nil {
	RETRY:
		log.Errorf("会话：%s，错误：%v", c.session.ID, err)
		log.Warning("将在两秒后重试...")
		time.Sleep(2 * time.Second)
		c.Resume()
		goto RETRY  // 断线重连
	}
	res.Body.Close()

	go func() {
	// to do...
	}()

	log.Infof("会话：%s, 地址：%s 已成功连接！", c.session.ID, c.session.URL)

	return nil
}

func (c *GuildClient) Resume() error {
	event := &tencent_guild.WebsocketPayload{
		Data: &tencent_guild.WebsocketResumeData{
			Token:     c.session.Token.GetAuthToken(),
			SessionID: c.session.ID,
			Seq:       c.session.LastSeq,
		},
	}
	event.OPCode = tencent_guild.WSResume
	return c.Write(event)
}

func (c *GuildClient) Write(message *tencent_guild.WebsocketPayload) error {
	msg, _ := json.Marshal(message)

	c.mu.Lock()
	err := c.conn.WriteMessage(wss.TextMessage, msg)
	c.mu.Unlock()
	if err != nil {
		log.Warning("向服务器发送信息失败：", err.Error())
		return err
	}

	log.Debug("向服务器发送信息：", global.BytesToString(msg))
	select {
	case <- time.After(30 * time.Second):
		log.Warning("向服务器发送信息超时！")
		return err
	}
}

func (c *GuildClient) nextSeq() uint64 {
	return atomic.AddUint64(&c.seq, 1)
}

func (c *GuildClient) Listen() {
	for {
		t, payload, err := c.conn.ReadMessage()
		if err != nil {
			log.Warning("WebSocket 服务器连接断开...将于2秒后重试")
			time.Sleep(2 * time.Second)
			c.Connect()
		}

		if t == wss.TextMessage {
			rsp := gjson.Parse(global.BytesToString(payload))
			log.Info(rsp)
		}
	}
}

func (c *GuildClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Errorf("会话：%s 关闭连接", c.session.ID)
	}
	c.heartBeatTicker.Stop()
}

package tencent_guild

import "RouterForGuild/global"

type WebsocketAP struct {
	URL               string            `json:"url"`
	Shards            uint32            `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

type SessionStartLimit struct {
	Total          uint32 `json:"total"`
	Remaining      uint32 `json:"remaining"`
	ResetAfter     uint32 `json:"reset_after"`
	MaxConcurrency uint32 `json:"max_concurrency"`
}

type ShardConfig struct {
	ShardID    uint32
	ShardCount uint32
}

type Session struct {
	ID      string
	URL     string
	Token   global.Token
	Intent  int
	LastSeq uint64
	Shards  ShardConfig
}

type WebsocketUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Bot      bool   `json:"bot"`
}

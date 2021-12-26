package tencent_guild

type EventType string

type WebSocketPayload struct {
	WebSocketPayloadBase
	Data interface{} `json:"d,omitempty"`
}

type WebSocketPayloadBase struct {
	OPCode OPCode    `json:"op"`
	Seq    uint64    `json:"s,omitempty"`
	Type   EventType `json:"t,omitempty"`
}

type WebSocketIdentity struct {
	Token      string   `json:"token"`
	Intents    Intent   `json:"intents"`
	Shard      []uint32 `json:"shard"`
	Properties struct {
		OS      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"$properties,omitempty"`
}

type WebSocketResume struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       uint64 `json:"seq"`
}

type WebSocketHello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type WebSocketIsReady struct {
	Version   int    `json:"version"`
	SessionID string `json:"session_id"`
	User      struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Bot      bool   `json:"bot"`
	}
	Shard []uint32 `json:"shard"`
}

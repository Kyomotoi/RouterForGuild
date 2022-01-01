package tencent_guild

type EventType string

type WebsocketPayload struct {
	WebsocketPayloadBase
	Data interface{} `json:"d,omitempty"`
}

type WebsocketPayloadBase struct {
	OPCode int    `json:"op"`
	Seq    uint64    `json:"s,omitempty"`
	Type   EventType `json:"t,omitempty"`
}

type WebsocketIdentity struct {
	Token      string   `json:"token"`
	Intents    int      `json:"intents"`
	Shard      []uint32 `json:"shard"`
	Properties struct {
		OS      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"$properties,omitempty"`
}

type WebsocketResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       uint64 `json:"seq"`
}

type WebsocketHello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type WebsocketIsReady struct {
	Version   int    `json:"version"`
	SessionID string `json:"session_id"`
	User      struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Bot      bool   `json:"bot"`
	}
	Shard []uint32 `json:"shard"`
}

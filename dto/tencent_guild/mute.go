package tencent_guild

type MuteBody struct {
	MuteEndTimestamp string `json:"mute_end_timestamp,omitempty"`
	MuteSeconds string `json:"mute_seconds,omitempty"`
}

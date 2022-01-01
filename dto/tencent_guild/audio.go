package tencent_guild

const (
	AudioStatusStart = iota
	AudioStatusPause
	AudioStatusResume
	AudioStatusStop
)

type AudioControl struct {
	AudioURL string `json:"audio_url,omitempty"`
	Text     string `json:"text"`
	Status   uint32 `json:"status"`
}

type AudioAction struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	AudioURL  string `json:"audio_url"`
	Text      string `json:"text"`
}

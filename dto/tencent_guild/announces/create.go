package tencent_guild

// CreateGlobalBody 请求：创建频道全局公告。返回：type Announces
type CreateGlobalBody struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
}

// CreateChannelBody 请求：创建子频道公告。返回：type Announces
type CreateChannelBody struct {
	MessageID string `json:"message_id"`
}

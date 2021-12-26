package tencent_guild

type Intent int

const (
	// Guilds 频道相关事件
	Guilds Intent = 1 << iota
	// GuildMembers 频道成员相关事件
	GuildMembers Intent = 1 << 1
	// UserMessages 用户消息相关事件（私域可用）
	UserMessages Intent = 1 << 9
	// 	GuildMessageReactions ...
	GuildMessageReactions Intent = 1 << 10
	// DirectMessage ...
	DirectMessage Intent = 1 << 12
	// ForumEvent ...
	ForumEvent Intent = 1 << 28
	// AudioAction 音频相关事件
	AudioAction Intent = 1 << 29
	// GuildAtMessages at事件
	GuildAtMessages Intent = 1 << 30
)

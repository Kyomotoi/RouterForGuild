package tencent_guild

type Channel struct {
	ID          string `json:"id"`
	GuildID     string `json:"guild_id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	SubType     int    `json:"sub_type"`
	Position    int64  `json:"position"`
	ParentID    string `json:"parent_id"`
	OwnerID     string `json:"owner_id"`
	PrivateType uint32 `json:"private_type"`
}

const (
	ChannelTypeText int = iota
	_
	ChannelTypeVoice
	_
	ChannelTypeCategory
	ChannelTypeLive int = 10000 + iota
	ChannelTypeApplication
	ChannelTypeForum
)

const (
	ChannelSubTypeChat int = iota
	ChannelSubTypeNotice
	ChannelSubTypeGuide
	ChannelSubTypeTeamGame
)

type ChannelPermission struct {
	ChannelID  string `json:"channel_id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type ChannelRolePermission struct {
	ChannelID  string `json:"channel_id,omitempty"`
	RoleID     string `json:"role_id,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type ChannelUpdatePermission struct {
	Add    string `json:"add,omitempty"`
	Remove string `json:"remove,omitempty"`
}

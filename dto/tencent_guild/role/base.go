package tencent_guild

// Role 频道身份组
type Role struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Color       uint32 `json:"color"`
	Hoist       uint32 `json:"hoist"`
	Number      uint32 `json:"number,omitempty"`
	MemberLimit uint32 `json:"member_limit,omitempty"`
}

// GuildRoles 返回：频道身份组列表
type GuildRoles struct {
	GuildID      string  `json:"guild_id"`
	Roles        []*Role `json:"roles"`
	RoleNumLimit string  `json:"role_num_limit"`
}

// RoleFilter ...
type RoleFilter struct {
	Name  int32 `json:"name"`
	Color int32 `json:"color"`
	Hoist int32 `json:"hoist"`
}

// RoleInfo ...
type RoleInfo struct {
	Name  string `json:"name"`
	Color uint32 `json:"color"`
	Hoist int32  `json:"hoist"`
}

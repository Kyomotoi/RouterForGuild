package tencent_guild

// RoleUpdateBody 请求：修改频道身份组
type RoleUpdateBody struct {
	Filter *RoleFilter `json:"filter"`
	Info   *RoleInfo   `json:"info"`
}

// RoleUpdateResult 返回：修改频道身份组
type RoleUpdateResult struct {
	GuildID string `json:"guild_id"`
	RoleID  string `json:"role_id"`
	Role    *Role  `json:"role"`
}

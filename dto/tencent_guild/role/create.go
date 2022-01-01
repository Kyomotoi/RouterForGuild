package tencent_guild

// RoleCreateBody 请求：创建频道身份组
type RoleCreateBody struct {
	Filter *RoleFilter `json:"filter"`
	Info   *RoleInfo   `json:"info"`
}

// RoleCreateResult 返回：创建频道身份组
type RoleCreateResult struct {
	RoleID string  `json:"role_id"`
	Role   []*Role `json:"role"`
}

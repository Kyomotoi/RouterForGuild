package tencent_guild

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Bot      bool   `json:"bot"`

	// union_openid 与 union_user_account 只有在单独拉取 member 信息的时候才会提供
	// 在其他的事件中所携带的 user 对象，均无这两个字段的内容

	// 特殊关联应用的 openid，需要特殊申请并配置后才会返回
	UnionOpenid string `json:"union_openid,omitempty"`
	// 机器人关联的互联应用的用户信息，与union_openid关联的应用是同一个
	UnionUserAccount string `json:"union_user_account,omitempty"`
}

// GetUserInfoBody 请求：当前用户频道列表。返回：type Guild
type GetUserInfoBody struct {
	Before string `json:"before"`
	After  string `json:"after"`
	Limit  int    `json:"limit"`
}

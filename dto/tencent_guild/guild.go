package tencent_guild

type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	OwnerID     string `json:"owner_id"`
	Owner       bool   `json:"owner"`
	MemberCount int    `json:"member_count"`
	MaxMembers  int64  `json:"max_members"`
	Description string `json:"description"`
	JoinedAt    string `json:"joined_at"`

	// 接下来三个字段未出现在官方文档
	// 可能需要某种特殊的权限

	// 频道列表
	Channels []*Channel `json:"channels"`
	// 游戏绑定公会区服ID
	UnionWorldID string `json:"union_world_id"`
	// 游戏绑定公会/战队ID
	UnionOrgID string `json:"union_org_id"`
}

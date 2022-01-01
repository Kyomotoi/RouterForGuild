package tencent_guild

type SendBody struct {
	Content string  `json:"content"`
	Embed   *Embed  `json:"embed"`
	Ark     *MsgArk `json:"ark"`
	Image   string  `json:"image"`
	MsgID   string  `json:"msg_id"`
}

type SendOfArk struct {
	Ark *MsgArk `json:"ark"`
}

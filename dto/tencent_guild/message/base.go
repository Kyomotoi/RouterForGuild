package tencent_guild

import (
	"RouterForGuild/dto/tencent_guild"
	"RouterForGuild/global"
)

type Message struct {
	ID              string                `json:"id"`
	ChannelID       string                `json:"channel_id"`
	GuildID         string                `json:"guild_id"`
	Content         string                `json:"content"`
	Timestamp       global.Timestamp      `json:"timestamp"`
	EditedTimestamp global.Timestamp      `json:"edited_timestamp"`
	MentionEveryone bool                  `json:"mention_everyone"`
	Author          *tencent_guild.User   `json:"author"`
	Attachments     []*Attachment         `json:"attachments"`
	Embeds          []*Embed              `json:"embeds"`
	Mentions        []*tencent_guild.User `json:"mentions"`
	Member          *tencent_guild.Member `json:"member"`
	Ark             *MsgArk               `json:"ark"`
}

type Attachment struct {
	URL string `json:"url"`
}

type Embed struct {
	Title     string            `json:"title"`
	Prompt    string            `json:"prompt"`
	Thumbnail []*EmbedThumbnail `json:"thumbnail,omitempty"`
	Timestamp global.Timestamp  `json:"timestamp,omitempty"`
	Fields    []*EmbedField     `json:"fields"`
}

type EmbedThumbnail struct {
	URL string `json:"url"`
}

type EmbedField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MsgArk struct {
	TemplateID int         `json:"template_id"`
	Kv         []*MsgArkKv `json:"kv"`
}

type MsgArkKv struct {
	Key   string       `json:"key"`
	Value string       `json:"value"`
	Obj   []*MsgArkObj `json:"obj"`
}

type MsgArkObj struct {
	ObjKv []*MsgArkObjKv `json:"obj_kv"`
}

type MsgArkObjKv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResponseBody struct {
	ID              string              `json:"id"`
	ChannelID       string              `json:"channel_id"`
	GuildID         string              `json:"guild_id"`
	Content         string              `json:"content"`
	Timestamp       global.Timestamp    `json:"timestamp"`
	TTS             bool                `json:"tts"`
	MentionEveryone bool                `json:"mention_everyone"`
	Author          *tencent_guild.User `json:"author"`
	Pinned          bool                `json:"pinned"`
	Type            int                 `json:"type"`
	Flags           int                 `json:"flags"`
}

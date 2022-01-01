package tencent_guild

import "RouterForGuild/dto/tencent_guild"

type Schedule struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	StartTimestamp string                `json:"start_timestamp"`
	EndTimestamp   string                `json:"end_timestamp"`
	Creator        *tencent_guild.Member `json:"creator"`
	JumpChannelID  string                `json:"jump_channel_id,omitempty"`
	RemindType     string                `json:"remind_type,omitempty"`
}

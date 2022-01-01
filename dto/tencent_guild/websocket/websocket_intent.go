package tencent_guild

const (
	IntentGuilds int = 1 << iota

	IntentGuildMembers
	IntentGuildBans
	IntentGuildEmojis
	IntentGuildIntegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessages
	IntentGuildMessageReactions

	IntentGuildMessageTyping

	IntentDirectMessage

	IntentForumEvent      int = 1 << 28
	IntentAudioAction     int = 1 << 29
	IntentGuildAtMessages int = 1 << 30
)

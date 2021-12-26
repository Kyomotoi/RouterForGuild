package tencent_guild

type OPCode int

const (
	WebsocketDispatchEvent OPCode = iota
	WebsocketHeartbeat
	WebsocketIdentity
	_
	_
	_
	WebsocketResume
	WebsocketConnect
	_
	WebsocketInvalidSession
	WebsocketHello
	WebsocketHeartbeatAck
)

package server

const (
	WssDisconnected = iota
	WssConnecting
	WssConnected
)

type WsMessage string

const (
	WsmUpdateRoom      = WsMessage("UPDATE_ROOM")
	WsmUpdateBoard     = WsMessage("UPDATE_BOARD")
	WsmConnectionReady = WsMessage("CONNECTION_READY")
)


package socket

import (
	socketio "github.com/googollee/go-socket.io"
)

// RegisterHandlers registers all socket event handlers to the server
func RegisterHandlers(server *socketio.Server, handler SocketHandler) {
	// Activity events
	server.OnEvent("", EventActivityJoin, handler.ActivityJoin)
	server.OnEvent("", EventActivityUpsert, handler.ActivityUpsert)
	server.OnEvent("", EventActivityLeave, handler.ActivityLeave)
}

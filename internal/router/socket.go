package router

import (
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/socket"
)

func RegisterSocket(mux *http.ServeMux, server *socketio.Server, handler socket.SocketHandler) {
	mux.Handle("/socket.io/", server)

	server.OnEvent("/", "trip:query", handler.GetTrips)
}

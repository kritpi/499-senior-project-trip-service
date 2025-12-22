package socket

import (
	"context"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func (h *socketHandler) GetTrips(socket socketio.Conn, msg any) {
	log.Println(msg)
	trips, err := h.svc.GetTrips(context.Background())
	if err != nil {
		return
	}
	socket.Emit("trip:data", trips)
}

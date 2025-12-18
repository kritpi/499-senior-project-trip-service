package socket

import (
	"context"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
)

type SocketHandler interface {
	Test(ctx context.Context, socket socketio.Conn, id string) error
	GetTrips(socket socketio.Conn, msg any)
}

type socketHandler struct {
	svc port.Service
}

func NewSocketIO(svc port.Service) SocketHandler {
	return &socketHandler{
		svc: svc,
	}
}

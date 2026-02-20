package socket

import (
	"context"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/port"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

type SocketHandler interface {
	Test(ctx context.Context, socket socketio.Conn, id string) error
	GetTrips(socket socketio.Conn, msg any)

	// ActivityJoin(socket socketio.Conn, in dto.ActivityJoinRequest) (*dto.ActivityResponse, error)
	ActivityJoin(socket socketio.Conn, in dto.ActivityJoinRequest)
	ActivityUpsert(socket socketio.Conn, in dto.ActivityUpsertRequest)
	ActivityBroadcast(socket socketio.Conn, in dto.ActivityJoinRequest)
	ActivityLeave(socket socketio.Conn, in dto.ActivityLeaveRequest)
	// ActivityConnect(socket socketio.Conn, in dto.ActivityConnectRequest) (*dto.ActivityResponse, error)
	// Activities(socket socketio.Conn, msg any) (*dto.Activity, error)
}

type socketHandler struct {
	svc    port.Service
	server *socketio.Server
}

func NewSocketIO(svc port.Service, server *socketio.Server) SocketHandler {
	return &socketHandler{
		svc:    svc,
		server: server,
	}
}

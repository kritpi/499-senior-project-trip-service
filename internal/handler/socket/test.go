package socket

import (
	"context"

	socketio "github.com/googollee/go-socket.io"
)

func (s *socketHandler) Test(ctx context.Context, socket socketio.Conn, id string) error {
	panic("unimplement")
	// s.svc.GetTrips(ctx)

}

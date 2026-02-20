package socket

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *socketHandler) ActivityBroadcast(socket socketio.Conn, in dto.ActivityJoinRequest) {
	socketContext := socket.Context().(*dto.SocketContext)
	
	socket.Emit(EventActivityBroadcast, map[string]interface{}{

	})
	fmt.Printf("socketContext: %v\n", socketContext)
}

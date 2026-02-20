package socket

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *socketHandler) ActivityLeave(socket socketio.Conn, in dto.ActivityLeaveRequest) {
	socketContext := socket.Context().(*dto.SocketContext)

	// Leave the room for this trip and date
	roomName := GetActivityRoomName(in.TripId, in.TripDate)
	socket.Leave(roomName)

	// Update socket context
	socketContext.Joined = false
	socketContext.TripId = 0
	socketContext.TripDate = ""

	// Acknowledge the leave
	socket.Emit(EventActivityLeave, map[string]interface{}{
		"success": true,
		"message": "Left activity planning session",
	})
}

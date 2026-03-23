package socket

import (
	"context"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *socketHandler) ActivityJoin(socket socketio.Conn, in dto.ActivityJoinRequest) {
	socketContext := socket.Context().(*dto.SocketContext)

	// Update socket context
	socketContext.TripId = in.TripId
	socketContext.TripDate = in.TripDate
	socketContext.Joined = true

	// Join the room for this trip and date
	roomName := GetActivityRoomName(in.TripId, in.TripDate)
	socket.Join(roomName)

	// Call service to handle join logic
	resp, err := h.svc.SocketActivityJoin(context.Background(), *in.ToDomain(socketContext.MemberId))
	if err != nil {
		socket.Emit(EventActivityJoin, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	socket.Emit(EventActivityJoin, dto.ActivityResponse{}.FromDomain(resp))

}

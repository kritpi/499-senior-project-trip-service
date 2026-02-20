package socket

import (
	"context"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
)

func (h *socketHandler) ActivityUpsert(socket socketio.Conn, in dto.ActivityUpsertRequest) {
	socketContext := socket.Context().(*dto.SocketContext)

	err := h.svc.SocketActivityUpsert(context.Background(), *in.ToDomain(socketContext.MemberId))
	if err != nil {
		socket.Emit(EventActivityUpsert, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Send confirmation to the client who made the update
	socket.Emit(EventActivityUpsert, map[string]interface{}{
		"message": "activities scheduled",
	})

}

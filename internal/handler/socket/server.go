package socket

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/dto"
	"github.com/kritpi/499-senior-project-trip-service/property"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
)

func NewSocket(handler SocketHandler, cfg property.Property) *socketio.Server {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})

	// Connection handler for root namespace
	server.OnConnect("/", func(c socketio.Conn) error {
		// Get token from query parameter (Socket.IO v2 compatibility)
		url := c.URL()
		token := url.Query().Get("token")

		if token == "" {
			log.Warnf("socket connection attempt with no token in query, socket_id: %s", c.ID())
			return fmt.Errorf("missing authentication token")
		}

		// Authenticate using the token (prepend "Bearer " prefix for the utility function)
		authHeader := fmt.Sprintf("Bearer %s", token)
		claims, err := utils.AuthenticateFromHeader(authHeader, cfg)
		if err != nil {
			log.Warnf("socket authentication failed: %+v, socket_id: %s", err, c.ID())
			return err
		}
		c.SetContext(&dto.SocketContext{
			MemberId: claims.ID,
			Joined:   false,
		})
		log.Infof("socket connected: %s, member_id: %s", c.ID(), claims.ID)
		return nil
	})

	// Disconnect handler
	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		socketContext, ok := c.Context().(*dto.SocketContext)
		if !ok || socketContext == nil {
			log.Infof("socket disconnected: %s, no context available, reason: %s", c.ID(), reason)
			return
		}
		log.Infof("socket disconnected: %s, member_id: %s, reason: %s", c.ID(), socketContext.MemberId, reason)
	})

	// Error handler
	server.OnError("/", func(c socketio.Conn, err error) {
		log.Errorf("socket error: %+v", err)
	})

	// Event handlers are registered separately via RegisterHandlers function

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socket.io serve error: %+v", err)
		}
	}()
	return server
}

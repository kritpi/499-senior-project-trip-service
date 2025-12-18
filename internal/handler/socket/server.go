package socket

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func NewSocket(handler SocketHandler) *socketio.Server {
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

	server.OnConnect("/", func(c socketio.Conn) error {
		log.Infof("socket connected: %s", c.ID())
		return nil
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		log.Infof("socket disconnected: %s, reason: %s", c.ID(), reason)
	})

	server.OnError("/", func(c socketio.Conn, err error) {
		log.Errorf("socket error: %+v", err)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socket.io serve error: %+v", err)
		}
	}()
	return server
}

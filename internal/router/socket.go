package router

import (
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kritpi/499-senior-project-trip-service/internal/handler/socket"
	"github.com/kritpi/499-senior-project-trip-service/property"
)

func RegisterSocket(mux *http.ServeMux, server *socketio.Server, handler socket.SocketHandler, cfg property.Property) {
	// Wrap server with CORS handling
	corsHandler := allowCors(server)
	mux.Handle("/socket.io/", corsHandler)

	// Event handlers are registered in socket.RegisterHandlers() during server initialization
}

func allowCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Allow specific origins including localhost:3000
		allowedOrigins := map[string]bool{
			"http://localhost:3000": true,
			"http://localhost:5173": true,
			"http://localhost:4200": true,
		}

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

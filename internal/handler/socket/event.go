package socket

const (
	EventConnect    = "connection"
	EventDisconnect = "disconnect"
)

type TripActivityPayload struct {
	TripId string `json:"trip_id"`
}

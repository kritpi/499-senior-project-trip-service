package socket

const (
	EventConnect    = "connection"
	EventDisconnect = "disconnect"

	// Activity events
	EventActivityJoin = "activity:join"
	EventActivityUpsert    = "activity:upsert"
	EventActivityBroadcast = "activity:broadcast"
	EventActivityLeave     = "activity:leave"
)

type TripActivityPayload struct {
	TripId string `json:"trip_id"`
}

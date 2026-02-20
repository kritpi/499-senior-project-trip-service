package socket

import "fmt"

// GetActivityRoomName returns a consistent room name for a trip's activity planning session
// Format: "trip:{trip_id}:date:{date}"
func GetActivityRoomName(tripId int, date string) string {
	return fmt.Sprintf("trip:%d:date:%s", tripId, date)
}

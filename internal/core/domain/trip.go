package domain

import "time"

// Trip represents a trip entity
type Trip struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Destination string    `json:"destination"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"` // planning, ongoing, completed, cancelled
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetTripsRequest is the request for getting trips
type GetTripsRequest struct {
}

// GetTripsResponse is the response for getting trips
type GetTripsResponse struct {
	Trips []Trip `json:"trips"`
}

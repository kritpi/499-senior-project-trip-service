package domain

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

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

// Upsert trip request
type UpsertTripRequest struct {
	ID           *int       `json:"id"`
	OwnerId      string     `json:"owner_id"`
	TripName     string     `json:"trip_name"`
	Desciption   *string    `json:"description"`
	StartDate    *time.Time `json:"start_date"` // 2025-07-11
	EndDate      *time.Time `json:"end_date"`   // 2025-07-11
	MainLocation *string    `json:"main_location"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type UpsertTripResponse struct {
	TripId int `json:"trip_id"`
}

func (t UpsertTripResponse) FromEntity(e entity.UpsertTripResponse) *UpsertTripResponse {
	return &UpsertTripResponse{
		TripId: t.TripId,
	}
}

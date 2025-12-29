package entity

import "time"

type UpsertTripResponse struct {
	TripId int `json:"trip_id"`
}

type Trip struct {
	ID           int       `json:"id"`
	OwnerId      string    `json:"owner_id"`
	TripName     string    `json:"name"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	MainLocation string    `json:"main_location`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

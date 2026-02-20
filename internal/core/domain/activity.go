package domain

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

type Activity struct {
	ID               string                 `json:"id"`
	TripId           int                    `json:"trip_id"`
	ActivityDate     time.Time              `json:"activity_date"`
	StartTime        *time.Time             `json:"start_time"`
	EndTime          *time.Time             `json:"end_time"`
	Note             *string                `json:"note"`
	Description      *string                `json:"description"`
	ActivityLocation *Location              `json:"activity_location"`
	Category         *enum.ActivityCategory `json:"category"`
	Rank             int                    `json:"rank"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
}

type Location struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type ActivityMember struct {
	TripId     int    `json:"trip_id"`
	ActivityId string `json:"activity_id"`
	MemberId   string `json:"member_id"`
}

type ActivityJoinRequest struct {
	MemberId string    `json:"member_id"`
	TripDate time.Time `json:"trip_date"`
	TripId   int       `json:"trip_id"`
}

type ActivityJoinResponse struct {
	TripId     int        `json:"trip_id"`
	Date       string     `json:"date"`
	Activities []Activity `json:"activities"`
	IsEditable bool       `json:"is_editable"`
}

type ActivityUpsertRequest struct {
	MemberId   string     `json:"member_id"`
	TripId     int        `json:"trip_id"`
	Date       time.Time  `json:"date"`
	Activities []Activity `json:"activities"`
}

type ActivityUpsertProcessRequest struct {
	TripId     int        `json:"trip_id"`
	Date       time.Time  `json:"date"`
	Activities []Activity `json:"activities"`
}

type GetActivitiesOfTheDayRequest struct {
	TripId int       `json:"trip_id"`
	Date   time.Time `json:"date"`
}

type GetActivitiesOfTheDayResponse struct {
	Activities []Activity `json:"activities"`
}

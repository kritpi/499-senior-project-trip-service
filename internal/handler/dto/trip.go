package dto

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
)

// Trip represents a trip entity
type Trip struct {
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
	ID *int `json:"trip_id"`
	// OwnerId      string  `json:"owner_id"`
	TripName     string  `json:"trip_name"`
	Desciption   *string `json:"description"`
	StartDate    *string `json:"start_date"` // 2025-07-11
	EndDate      *string `json:"end_date"`   // 2025-07-11
	MainLocation *string `json:"main_location"`
}

func (t UpsertTripRequest) ToDomain(memberId string) *domain.UpsertTripRequest {
	var startDate *time.Time
	if t.StartDate != nil {
		var err error
		startDate, err = utils.DateStringToStartDate(*t.StartDate, utils.DATE_FORMAT)
		if err != nil {
			return nil
		}
	}

	var endDate *time.Time
	if t.EndDate != nil {
		var err error
		endDate, err = utils.DateStringToStartDate(*t.EndDate, utils.DATE_FORMAT)
		if err != nil {
			return nil
		}
	}

	return &domain.UpsertTripRequest{
		ID:           t.ID,
		OwnerId:      memberId,
		TripName:     t.TripName,
		Desciption:   t.Desciption,
		StartDate:    startDate,
		EndDate:      endDate,
		MainLocation: t.MainLocation,
	}
}

type UpsertTripResponse struct {
	TripId int `json:"trip_id"`
}

func (t UpsertTripResponse) FromDomain(dm *domain.UpsertTripResponse) *UpsertTripResponse {
	return &UpsertTripResponse{
		TripId: dm.TripId,
	}
}

type TripMember struct {
	MemberId string        `json:"member_id"`
	Role     enum.TripRole `json:"role"`
}

type CreateTripMemberRequest struct {
	TripId    int          `json:"trip_id"`
	Members   []TripMember `json:"members"`
	CreatedAt time.Time    `json:"created_at"`
}

type GetMemberTripsRequest struct {
	MemberId string `json:"member_id"`
}

func (t GetMemberTripsRequest) ToDomain(memberId string) *domain.GetMemberTripsRequest {
	return &domain.GetMemberTripsRequest{
		MemberId: memberId,
	}
}

type GetMemberTripsResponse struct {
	Trips []MemberTrips `json:"trips"`
}

type MemberTrips struct {
	TripId       int           `json:"trip_id"`
	TripName     string        `json:"trip_name"`
	StartDate    string        `json:"start_date"`
	EndDate      string        `json:"end_date"`
	MainLocation string        `json:"main_location"`
	Role         enum.TripRole `json:"role"`
}

func (t GetMemberTripsResponse) FromDomain(dm *domain.GetMemberTripsResponse) *GetMemberTripsResponse {
	trip := make([]MemberTrips, len(dm.Trips))
	for i, t := range dm.Trips {
		trip[i] = MemberTrips{
			TripId:       t.TripId,
			TripName:     t.TripName,
			StartDate:    utils.DateTimeToDateString(t.StartDate),
			EndDate:      utils.DateTimeToDateString(t.EndDate),
			MainLocation: t.MainLocation,
			Role:         t.Role,
		}
	}

	return &GetMemberTripsResponse{
		Trips: trip,
	}
}

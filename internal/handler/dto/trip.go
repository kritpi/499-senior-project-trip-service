package dto

import (
	"github.com/gofiber/fiber/v2/log"
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
			log.Errorf("error parsing date string: %+v", err)
			startDate = nil
		}
	}

	var endDate *time.Time
	if t.EndDate != nil {
		var err error
		endDate, err = utils.DateStringToStartDate(*t.EndDate, utils.DATE_FORMAT)
		if err != nil {
			log.Errorf("error parsing date string: %+v", err)
			endDate = nil
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
	MemberId string          `json:"member_id"`
	Role     enum.MemberRole `json:"role"`
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
	TripId       int             `json:"trip_id"`
	TripName     string          `json:"trip_name"`
	StartDate    string          `json:"start_date"`
	EndDate      string          `json:"end_date"`
	MainLocation string          `json:"main_location"`
	Role         enum.MemberRole `json:"role"`
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

// GetTripByIdRequest is the DTO for getting a trip by ID
type GetTripByIdRequest struct {
	TripId int `json:"trip_id"`
}

// TripMemberDetails represents a member in a trip with their role
type TripMemberDetails struct {
	MemberId string          `json:"member_id"`
	Name     string          `json:"name"`
	ImageUrl string          `json:"image_url"`
	Role     enum.MemberRole `json:"role"`
}

// GetTripByIdResponse is the DTO response for getting a trip by ID
type GetTripByIdResponse struct {
	ID           int                 `json:"trip_id"`
	OwnerId      string              `json:"owner_id"`
	TripName     string              `json:"trip_name"`
	Description  string              `json:"description"`
	StartDate    string              `json:"start_date"`
	EndDate      string              `json:"end_date"`
	MainLocation string              `json:"main_location"`
	Members      []TripMemberDetails `json:"members"`
}

func (t GetTripByIdResponse) FromDomain(resp *domain.GetTripByIdResponse) *GetTripByIdResponse {
	trip := resp.Trip

	// Convert members
	members := make([]TripMemberDetails, len(resp.Members))
	for i, m := range resp.Members {
		members[i] = TripMemberDetails{
			MemberId: m.MemberId,
			Name:     m.Name,
			ImageUrl: m.ImageUrl,
			Role:     m.Role,
		}
	}

	return &GetTripByIdResponse{
		ID:           trip.ID,
		OwnerId:      trip.OwnerId,
		TripName:     trip.TripName,
		Description:  trip.Description,
		StartDate:    utils.DateTimeToDateString(trip.StartDate),
		EndDate:      utils.DateTimeToDateString(trip.EndDate),
		MainLocation: trip.MainLocation,
		Members:      members,
	}
}

type TripInvitationRequest struct {
	TripId int             `json:"trip_id"`
	Member []InvitedMember `json:"member"`
}

type InvitedMember struct {
	Email string          `json:"email"`
	Role  enum.MemberRole `json:"role"`
}

func (t TripInvitationRequest) ToDomain(memberId string) *domain.TripInvitationRequest {
	member := make([]domain.InvitedMember, len(t.Member))
	for i, m := range t.Member {
		member[i] = domain.InvitedMember{
			Email: m.Email,
			Role:  m.Role,
		}
	}
	return &domain.TripInvitationRequest{
		MemberId: memberId,
		TripId:   t.TripId,
		Member:   member,
	}
}

type InviteMemberResponse struct {
	Message string   `json:"message"`
	Member  []string `json:"member"`
}

func (i InviteMemberResponse) FromDomain(resp *domain.InviteMemberResponse) *InviteMemberResponse {
	member := make([]string, len(resp.Member))
	for i, m := range resp.Member {
		member[i] = m
	}
	return &InviteMemberResponse{
		Message: resp.Message,
		Member:  member,
	}
}

package dto

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
)

type Activity struct {
	ID string `json:"id"`
	// TripId           int                    `json:"trip_id"`
	// ActivityDate     *string                `json:"activity_date"`
	StartTime        *string                `json:"start_time"`
	EndTime          *string                `json:"end_time"`
	Note             *string                `json:"note"`
	Description      *string                `json:"description"`
	ActivityLocation *Location              `json:"activity_location"`
	Category         *enum.ActivityCategory `json:"category"`
	Rank             int                    `json:"rank"`
	CreatedAt        string                 `json:"created_at"`
	UpdatedAt        string                 `json:"updated_at"`
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
	TripDate string `json:"trip_date"`
	TripId   int    `json:"trip_id"`
}

func (a ActivityJoinRequest) ToDomain(memberId string) *domain.ActivityJoinRequest {
	tripDate, err := utils.DateStringToDateTime(a.TripDate, utils.DATE_FORMAT)
	if err != nil {
		log.Errorf("error parsing date string: %+v", err)
		return nil
	}
	return &domain.ActivityJoinRequest{
		MemberId: memberId,
		TripDate: *tripDate,
		TripId:   a.TripId,
	}
}

type ActivityResponse struct {
	TripId     int        `json:"trip_id"`
	Date       string     `json:"date"`
	Activities []Activity `json:"activities"`
	IsEditable bool       `json:"is_editable"`
}

func (a ActivityResponse) FromDomain(dm *domain.ActivityJoinResponse) *ActivityResponse {
	activities := make([]Activity, len(dm.Activities))
	for i, act := range dm.Activities {
		// actDate := utils.DateTimeToDateString(&act.ActivityDate)

		var startTime *string
		if act.StartTime != nil {
			formattedStartTime := act.StartTime.Format(string(utils.TIME_FORMAT))
			startTime = &formattedStartTime
		}

		var endTime *string
		if act.EndTime != nil {
			formattedEndTime := act.EndTime.Format(string(utils.TIME_FORMAT))
			endTime = &formattedEndTime
		}

		createdAtStr := utils.DateTimeToDateString(&act.CreatedAt)
		updatedAtStr := utils.DateTimeToDateString(&act.UpdatedAt)

		activities[i] = Activity{
			ID: act.ID,
			// TripId:       act.TripId,
			// ActivityDate: &actDate,
			StartTime:   startTime,
			EndTime:     endTime,
			Note:        act.Note,
			Description: act.Description,
			ActivityLocation: &Location{
				Name:    act.ActivityLocation.Name,
				Address: act.ActivityLocation.Address,
				Lat:     act.ActivityLocation.Lat,
				Lng:     act.ActivityLocation.Lng,
			},
			Category:  act.Category,
			Rank:      act.Rank,
			CreatedAt: createdAtStr,
			UpdatedAt: updatedAtStr,
		}
	}
	return &ActivityResponse{
		TripId:     dm.TripId,
		Date:       dm.Date,
		Activities: activities,
		IsEditable: dm.IsEditable,
	}
}

type ActivityUpsertRequest struct {
	TripId     int        `json:"trip_id"`
	Date       string     `json:"date"`
	Activities []Activity `json:"activities"`
}

func (a ActivityUpsertRequest) ToDomain(memberId string) *domain.ActivityUpsertRequest {
	tripDate, err := utils.DateStringToDateTime(a.Date, utils.DATE_FORMAT)
	if err != nil {
		log.Errorf("error parsing date string: %+v", err)
		return nil
	}

	activities := make([]domain.Activity, len(a.Activities))
	for i, act := range a.Activities {
		actDate, err := utils.DateStringToDateTime(a.Date, utils.DATE_FORMAT)
		if err != nil {
			log.Errorf("error parsing date string: %+v", err)
			return nil
		}
		startTime, err := utils.TimeStringToTime(*act.StartTime, utils.TIME_FORMAT)
		if err != nil {
			log.Errorf("error parsing start time: %+v", err)
			return nil
		}
		endTime, err := utils.TimeStringToTime(*act.EndTime, utils.TIME_FORMAT)
		if err != nil {
			log.Errorf("error parsing end time: %+v", err)
			return nil
		}
		activities[i] = domain.Activity{
			ID:           act.ID,
			TripId:       a.TripId,
			ActivityDate: *actDate,
			StartTime:    startTime,
			EndTime:      endTime,
			Note:         act.Note,
			Description:  act.Description,
			ActivityLocation: &domain.Location{
				Name:    act.ActivityLocation.Name,
				Address: act.ActivityLocation.Address,
				Lat:     act.ActivityLocation.Lat,
				Lng:     act.ActivityLocation.Lng,
			},
			Category: act.Category,
			Rank:     act.Rank,
		}
	}
	return &domain.ActivityUpsertRequest{
		MemberId:   memberId,
		TripId:     a.TripId,
		Date:       *tripDate,
		Activities: activities,
	}
}

type ActivityLeaveRequest struct {
	TripDate string `json:"trip_date"`
	TripId   int    `json:"trip_id"`
}

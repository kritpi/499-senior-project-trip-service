package utils

import (
	"time"
)

const (
	DATETIME_WITHZONE_FORMAT = "2006-01-02T15:04:05-07:00"
)

type DateFormat string

const (
	DATE_FORMAT DateFormat = "2006-01-02"
	TIME_FORMAT DateFormat = "15:04"
)

func DateStringToStartDate(dateString string, format DateFormat) (*time.Time, error) {
	t, err := time.Parse(string(format), dateString)
	if err != nil {
		return nil, err
	}
	d := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return &d, nil
}

func DateStringToEndDate(dateString string, format DateFormat) (*time.Time, error) {
	t, err := time.Parse(string(format), dateString)
	if err != nil {
		return nil, err
	}
	d := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.UTC)
	return &d, nil
}

func DateTimeToDateString(t *time.Time) string {
	return t.Format(string(DATE_FORMAT))
}

func DateStringToDateTime(dateString string, format DateFormat) (*time.Time, error) {
	t, err := time.Parse(string(format), dateString)
	if err != nil {
		return nil, err
	}
	d := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return &d, nil
}

func TimeStringToTime(timeString string, format DateFormat) (*time.Time, error) {
	t, err := time.Parse(string(format), timeString)
	if err != nil {
		return nil, err
	}

	result := time.Date(1970, 1, 1, t.Hour(), t.Minute(), 0, 0, time.UTC)
	return &result, nil
}

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
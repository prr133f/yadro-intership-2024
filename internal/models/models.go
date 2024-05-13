package models

import (
	"time"
)

type State struct {
	Computers  int
	StartTime  time.Time
	EndTime    time.Time
	HourlyRate int
}

type Event struct {
	Time       time.Time
	ID         int
	ClientName string
	TableID    int
}

type Data struct {
	State  State
	Events []Event
}

type Table struct {
	Margin     int
	TimeInWork time.Time
}

type TableMap struct {
	Map map[int]Table
}

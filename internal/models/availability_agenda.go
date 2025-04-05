package models

import "time"

type AvailabilityAgenda struct {
	Id                     int
	Name                   string
	AvailabilityDays       []AvailabilityDay
	LimitedEventsPerPerson int
	FrequencyLimit         Frequency //Lifetime, Weekly, Monthly, Biweekly, Annually
}

type AvailabilityDay struct {
	WeekDay      time.Weekday
	StartingTime time.Time
	EndingTime   time.Time
}

type Frequency int

const (
	Lifetime Frequency = iota
	Weekly
	Monthly
	Biweekly
	Annually
)

func (aa AvailabilityAgenda) Init() {
	
}

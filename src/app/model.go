package main

import "time"

var employees = map[string]Employee{
	"962134": Employee{
		ID:        962134,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Position:  "CEO",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176158": Employee{
		ID:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		Position:  "CTO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  29,
	},
}

var TimeOff = map[string][]TimeOff{
	"962134": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		},
	},
	"176158": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		},
	},
}

type Employee struct {
	ID        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string
	Amount    float32
	StartDate time.Time
	Status    string
}

package models

import "time"

var Employees = map[string]Employee{
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

var TimesOff = map[string][]TimeOff{
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
		{
			Type:      "Holiday",
			Amount:    16.,
			StartDate: time.Date(2016, 7, 10, 0, 0, 0, 0, time.UTC),
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
	ID        int
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	StartDate time.Time
	Position  string `form:"position"`
	TotalPTO  float32
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string    `json:"type"`
	Amount    float32   `json:"amount"`
	StartDate time.Time `json:"start_date"`
	Status    string    `json:"status"`
}

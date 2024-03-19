package model

import "time"

type ReportPost struct {
	ReportText   string    `json:"reportTitle"`
	ReportDate   time.Time `json:"reportDate"`
	ReportStatus bool      `json:"reportStatus"`
}

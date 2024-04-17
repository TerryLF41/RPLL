package model

import "time"

// ReportPost represents a report post model
type ReportPost struct {
	PostNo       int       `json:"postNo"`
	UserID       int       `json:"userID"`
	ReportText   string    `json:"reportText"`
	ReportDate   time.Time `json:"reportDate"`
	ReportStatus bool      `json:"reportStatus"`
}

// ReportPostModelFactory interface defines methods for report post model
type ReportPostModelFactory interface {
	CreateReportPost(postNo int, userID int, reportText string, reportDate time.Time, reportStatus bool) *ReportPost
}

// ConcreteReportPostModelFactory struct implements ReportPostModelFactory interface
type ConcreteReportPostModelFactory struct{}

// CreateReportPost creates a new report post instance
func (factory *ConcreteReportPostModelFactory) CreateReportPost(postNo int, userID int, reportText string, reportDate time.Time, reportStatus bool) *ReportPost {
	return &ReportPost{
		PostNo:       postNo,
		UserID:       userID,
		ReportText:   reportText,
		ReportDate:   reportDate,
		ReportStatus: reportStatus,
	}
}

// NewReportPostModelFactory creates a new report post model factory
func NewReportPostModelFactory() ReportPostModelFactory {
	return &ConcreteReportPostModelFactory{}
}

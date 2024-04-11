package model

import "time"

// UserLog represents a user log model
type UserLog struct {
	LogNo     int       `json:"logNo"`
	UserID    int       `json:"userId"`
	LogType   string    `json:"logType"`
	IpAddress string    `json:"ipAddress"`
	LogTime   time.Time `json:"logTime"`
}

// UserLogModelFactory interface defines methods for user log model
type UserLogModelFactory interface {
	CreateUserLog(logNo, userID int, logType, ipAddress string, logTime time.Time) *UserLog
}

// ConcreteUserLogModelFactory struct implements UserLogModelFactory interface
type ConcreteUserLogModelFactory struct{}

// CreateUserLog creates a new user log instance
func (factory *ConcreteUserLogModelFactory) CreateUserLog(logNo, userID int, logType, ipAddress string, logTime time.Time) *UserLog {
	return &UserLog{
		LogNo:     logNo,
		UserID:    userID,
		LogType:   logType,
		IpAddress: ipAddress,
		LogTime:   logTime,
	}
}

// NewUserLogModelFactory creates a new user log model factory
func NewUserLogModelFactory() UserLogModelFactory {
	return &ConcreteUserLogModelFactory{}
}

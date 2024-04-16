package model

import "time"

type User struct {
	UserID         int          `json:"userId"`
	Username       string       `json:"username"`
	Password       string       `json:"password,omitempty"`
	Email          string       `json:"email"`
	ProfileDesc    string       `json:"profileDesc"`
	JoinDate       time.Time    `json:"joinDate"`
	UserType       int          `json:"usertype,omitempty"`
	BanStatus      bool         `json:"banstatus"`
	ProfilePicture string       `json:"profilePicture"`
	ListReportPost []ReportPost `json:"ListReportPost,omitempty"`
}

// UserModelFactory interface defines methods for user model
type UserModelFactory interface {
	CreateUser(userID int, username, password, email, profileDesc string, joinDate time.Time, userType int, banStatus bool, profilePicture string, listReportPost []ReportPost) *User
}

// ConcreteUserModelFactory struct implements UserModelFactory interface
type ConcreteUserModelFactory struct{}

func (factory *ConcreteUserModelFactory) CreateUser(userID int, username, password, email, profileDesc string, joinDate time.Time, userType int, banStatus bool, profilePicture string, listReportPost []ReportPost) *User {
	return &User{
		UserID:         userID,
		Username:       username,
		Password:       password,
		Email:          email,
		ProfileDesc:    profileDesc,
		JoinDate:       joinDate,
		UserType:       userType,
		BanStatus:      banStatus,
		ProfilePicture: profilePicture,
		ListReportPost: listReportPost,
	}
}

// NewUserModelFactory creates a new user model factory
func NewUserModelFactory() UserModelFactory {
	return &ConcreteUserModelFactory{}
}

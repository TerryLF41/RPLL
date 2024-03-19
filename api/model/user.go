package model

type User struct {
	UserID			int         `json:"userId"`
	Username        string      `json:"username"`
	Password        string      `json:"password,omitempty"`
	Email           string      `json:"email"`
	ProfileDesc		string		`json:"profileDesc"`
	JoinDate		date		`json:"joinDate"`
	UserType        int         `json:"usertype,omitempty"`
	BanStatus       bool        `json:"banstatus,omitempty"`
}

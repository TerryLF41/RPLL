package model

import "time"

type UserLog struct {
	LogNo     int       `json:"logNo"`
	UserID    int       `json:"userId"`
	LogType   string    `json:"logType"`
	IpAddress string    `json:"ipAddress"`
	LogTime   time.Time `json:"logTime"`
}

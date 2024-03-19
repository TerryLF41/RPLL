package model

import "time"

type Thread struct {
	ThreadNo    int       `json:"threadNo"`
	ThreadTitle string    `json:"threadTitle"`
	ThreadDesc  string    `json:"threadDesc"`
	CreateDate  time.Time `json:"createDate"`
	BanStatus   bool      `json:"banStatus"`
	PostList    []Post    `json:"post,omitempty"`
}

package model

import "time"

type Topic struct {
	TopicNo    int       `json:"topicNo"`
	TopicTitle string    `json:"topicTitle"`
	TopicDesc  string    `json:"topicDesc"`
	CreateDate time.Time `json:"createDate"`
	BanStatus  bool      `json:"banstatus"`
	ThreadList []Thread  `json:"threadList,omitempty"`
}

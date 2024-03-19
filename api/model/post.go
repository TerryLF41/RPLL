package model

import "time"

type Post struct {
	PostNo    int       `json:"postNo"`
	ThreadNo  int       `json:"threadNo"`
	UserId    int       `json:"userId"`
	Reply     int       `json:"reply"`
	PostText  string    `json:"postText"`
	PostImage string    `json:"postImage"`
	PostDate  time.Time `json:"postDate"`
	BanStatus bool      `json:"banStatus"`
}

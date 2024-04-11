package model

import "time"

// Thread represents a thread model
type Thread struct {
	ThreadNo    int       `json:"threadNo"`
	TopicNo     int       `json:"topicNo"`
	ThreadTitle string    `json:"threadTitle"`
	ThreadDesc  string    `json:"threadDesc"`
	CreateDate  time.Time `json:"createDate"`
	BanStatus   bool      `json:"banStatus"`
	PostList    []Post    `json:"post,omitempty"`
}

// ThreadModelFactory interface defines methods for thread model
type ThreadModelFactory interface {
	CreateThread(threadNo int, topicNo int, threadTitle string, threadDesc string, createDate time.Time, banStatus bool, postList []Post) *Thread
}

// ConcreteThreadModelFactory struct implements ThreadModelFactory interface
type ConcreteThreadModelFactory struct{}

// CreateThread creates a new thread instance
func (factory *ConcreteThreadModelFactory) CreateThread(threadNo int, topicNo int, threadTitle string, threadDesc string, createDate time.Time, banStatus bool, postList []Post) *Thread {
	return &Thread{
		ThreadNo:    threadNo,
		TopicNo:     topicNo,
		ThreadTitle: threadTitle,
		ThreadDesc:  threadDesc,
		CreateDate:  createDate,
		BanStatus:   banStatus,
		PostList:    postList,
	}
}

// NewThreadModelFactory creates a new thread model factory
func NewThreadModelFactory() ThreadModelFactory {
	return &ConcreteThreadModelFactory{}
}

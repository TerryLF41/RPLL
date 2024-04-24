package model

import "time"

// Topic represents a topic model
type Topic struct {
	TopicNo      int       `json:"topicNo"`
	TopicTitle   string    `json:"topicTitle"`
	TopicDesc    string    `json:"topicDesc"`
	CreateDate   time.Time `json:"createDate"`
	BanStatus    bool      `json:"banstatus"`
	TopicPicture string    `json:"topicPicture"`
	ThreadCount  *int      `json:"threadCount,omitempty"`
	ThreadList   []Thread  `json:"threadList,omitempty"`
}

// TopicModelFactory interface defines methods for topic model
type TopicModelFactory interface {
	CreateTopic(topicNo int, topicTitle, topicDesc string, createDate time.Time, banStatus bool, topicPicture string, threadCount *int, threadList []Thread) *Topic
}

// ConcreteTopicModelFactory struct implements TopicModelFactory interface
type ConcreteTopicModelFactory struct{}

// CreateTopic creates a new topic instance
func (factory *ConcreteTopicModelFactory) CreateTopic(topicNo int, topicTitle string, topicDesc string, createDate time.Time, banStatus bool, topicPicture string, threadCount *int, threadList []Thread) *Topic {
	return &Topic{
		TopicNo:      topicNo,
		TopicTitle:   topicTitle,
		TopicDesc:    topicDesc,
		CreateDate:   createDate,
		BanStatus:    banStatus,
		TopicPicture: topicPicture,
		ThreadCount:  threadCount,
		ThreadList:   threadList,
	}
}

// NewTopicModelFactory creates a new topic model factory
func NewTopicModelFactory() TopicModelFactory {
	return &ConcreteTopicModelFactory{}
}

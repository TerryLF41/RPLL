package model

import "time"

// Topic represents a topic model
type Topic struct {
	TopicNo      int       `json:"topicNo"`
	TopicTitle   string    `json:"topicTitle"`
	TopicPicture string    `json:"topicPicture"`
	TopicDesc    string    `json:"topicDesc"`
	CreateDate   time.Time `json:"createDate"`
	BanStatus    bool      `json:"banstatus"`
	ThreadList   []Thread  `json:"threadList,omitempty"`
}

// TopicModelFactory interface defines methods for topic model
type TopicModelFactory interface {
	CreateTopic(topicNo int, topicTitle, topicDesc, topicPicture string, createDate time.Time, banStatus bool, threadList []Thread) *Topic
}

// ConcreteTopicModelFactory struct implements TopicModelFactory interface
type ConcreteTopicModelFactory struct{}

// CreateTopic creates a new topic instance
func (factory *ConcreteTopicModelFactory) CreateTopic(topicNo int, topicTitle string, topicDesc, topicPicture string, createDate time.Time, banStatus bool, threadList []Thread) *Topic {
	return &Topic{
		TopicNo:      topicNo,
		TopicTitle:   topicTitle,
		TopicPicture: topicPicture,
		TopicDesc:    topicDesc,
		CreateDate:   createDate,
		BanStatus:    banStatus,
		ThreadList:   threadList,
	}
}

// NewTopicModelFactory creates a new topic model factory
func NewTopicModelFactory() TopicModelFactory {
	return &ConcreteTopicModelFactory{}
}

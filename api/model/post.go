package model

import "time"

// Post represents a post model
type Post struct {
	PostNo    int       `json:"postNo"`
	ThreadNo  int       `json:"threadNo"`
	UserId    int       `json:"userId"`
	ReplyTo   int       `json:"replyTo,omitempty"`
	PostText  string    `json:"postText"`
	PostImage string    `json:"postImage"`
	PostDate  time.Time `json:"postDate"`
	BanStatus bool      `json:"banStatus"`
}

// PostModelFactory interface defines methods for post model
type PostModelFactory interface {
	CreatePost(postNo, threadNo, userId, replyTo int, postText, postImage string, postDate time.Time, banStatus bool) *Post
}

// ConcretePostModelFactory struct implements PostModelFactory interface
type ConcretePostModelFactory struct{}

// CreatePost creates a new post instance
func (factory *ConcretePostModelFactory) CreatePost(postNo, threadNo, userId, replyTo int, postText, postImage string, postDate time.Time, banStatus bool) *Post {
	return &Post{
		PostNo:    postNo,
		ThreadNo:  threadNo,
		UserId:    userId,
		ReplyTo:   replyTo,
		PostText:  postText,
		PostImage: postImage,
		PostDate:  postDate,
		BanStatus: banStatus,
	}
}

// NewPostModelFactory creates a new post model factory
func NewPostModelFactory() PostModelFactory {
	return &ConcretePostModelFactory{}
}

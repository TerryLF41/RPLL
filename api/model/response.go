package model

// Response represents a generic response model
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//Generic Response
type SuccessResponseModelFactory interface {
	CreateSuccessResponse(message string, data interface{}) *Response
}

type ConcreteSuccessResponseModelFactory struct{}

func (factory *ConcreteSuccessResponseModelFactory) CreateSuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Status:  200,
		Message: message,
		Data:    data,
	}
}

// NewResponseModelFactory creates a new response model factory
func NewSuccessResponseModelFactory() SuccessResponseModelFactory {
	return &ConcreteSuccessResponseModelFactory{}
}

//Error Response
type ErrorResponseModelFactory interface {
	CreateErrorResponse(message string) *Response
}

type ConcreteErrorResponseModelFactory struct{}

func (factory *ConcreteErrorResponseModelFactory) CreateErrorResponse(message string) *Response {
	return &Response{
		Status:  400,
		Message: message,
		Data:    nil,
	}
}

func NewErrorResponseModelFactory() ErrorResponseModelFactory {
	return &ConcreteErrorResponseModelFactory{}
}

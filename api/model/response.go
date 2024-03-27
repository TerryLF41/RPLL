package model

// Response represents a generic response model
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//Generic Response
type GenericResponseModelFactory interface {
	CreateGenericResponse(data interface{}) *Response
}

type ConcreteGenericResponseModelFactory struct{}

func (factory *ConcreteGenericResponseModelFactory) CreateGenericResponse(data interface{}) *Response {
	return &Response{
		Status:  200,
		Message: "Success",
		Data:    data,
	}
}

// NewResponseModelFactory creates a new response model factory
func NewGenericResponseModelFactory() GenericResponseModelFactory {
	return &ConcreteGenericResponseModelFactory{}
}

//Error Response
type ErrorResponseModelFactory interface {
	CreateErrorResponse(message string) *Response
}

type ConcreteErrorResponseModelFactory struct{}

func (factory *ConcreteErrorResponseModelFactory) CreateErrorResponse(message string) *Response {
	return &Response{
		Status:  404,
		Message: message,
		Data:    nil,
	}
}

func NewErrorResponseModelFactory() ErrorResponseModelFactory {
	return &ConcreteErrorResponseModelFactory{}
}

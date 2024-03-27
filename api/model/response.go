package model

// GenericResponse represents a generic response model
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// GenericResponseModelFactory interface defines methods for generic response model
type ResponseModelFactory interface {
	CreateGenericResponse(status int, message string, data interface{}) *GenericResponse
}

// ConcreteResponseModelFactory struct implements GenericResponseModelFactory interface
type ConcreteResponseModelFactory struct{}

// CreateGenericResponse creates a new generic response instance
func (factory *ConcreteResponseModelFactory) CreateSuccessResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (factory *ConcreteResponseModelFactory) CreateErrorResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

// NewGenericResponseModelFactory creates a new generic response model factory
func NewResponseModelFactory() ResponseModelFactory {
	return &ConcreteGenericResponseModelFactory{}
}

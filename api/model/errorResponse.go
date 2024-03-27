package model

// ErrorResponse represents an error response model
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// // ErrorResponseModelFactory interface defines methods for error response model
// type ErrorResponseModelFactory interface {
// 	CreateErrorResponse(status int, message string) *ErrorResponse
// }

// // ConcreteErrorResponseModelFactory struct implements ErrorResponseModelFactory interface
// type ConcreteErrorResponseModelFactory struct{}

// // CreateErrorResponse creates a new error response instance
// func (factory *ConcreteErrorResponseModelFactory) CreateErrorResponse(status int, message string) *ErrorResponse {
// 	return &ErrorResponse{
// 		Status:  status,
// 		Message: message,
// 	}
// }

// // NewErrorResponseModelFactory creates a new error response model factory
// func NewErrorResponseModelFactory() ErrorResponseModelFactory {
// 	return &ConcreteErrorResponseModelFactory{}
// }

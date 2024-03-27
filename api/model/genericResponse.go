package model

// GenericResponse represents a generic response model
type GenericResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// // GenericResponseModelFactory interface defines methods for generic response model
// type GenericResponseModelFactory interface {
// 	CreateGenericResponse(status int, message string, data interface{}) *GenericResponse
// }

// // ConcreteGenericResponseModelFactory struct implements GenericResponseModelFactory interface
// type ConcreteGenericResponseModelFactory struct{}

// // CreateGenericResponse creates a new generic response instance
// func (factory *ConcreteGenericResponseModelFactory) CreateGenericResponse(status int, message string, data interface{}) *GenericResponse {
// 	return &GenericResponse{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}
// }

// // NewGenericResponseModelFactory creates a new generic response model factory
// func NewGenericResponseModelFactory() GenericResponseModelFactory {
// 	return &ConcreteGenericResponseModelFactory{}
// }

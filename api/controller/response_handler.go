package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, message string) {
	var errorResponse model.ErrorResponse
	errorResponse.Status = 400
	errorResponse.Message = message
	w.Header().Set("Content=Type", "application/json")
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		log.Println(err)
	}
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	var response model.ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content=Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendSuccessResponse(w http.ResponseWriter, message string, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var response model.GenericResponse
	response.Status = 200
	response.Message = message
	response.Data = value
	w.Header().Set("Content=Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

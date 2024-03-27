package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content=Type", "application/json")
	responseFactory := model.NewErrorResponseModelFactory()
	response := responseFactory.CreateErrorResponse(message)
	json.NewEncoder(w).Encode(response)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("Content=Type", "application/json")
	responseFactory := model.NewErrorResponseModelFactory()
	response := responseFactory.CreateErrorResponse("Unauthorized access")
	json.NewEncoder(w).Encode(response)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendSuccessResponse(w http.ResponseWriter, message string, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseFactory := model.NewSuccessResponseModelFactory()
	response := responseFactory.CreateSuccessResponse(
		"Successfully retrieved topic",
		value,
	)
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content=Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

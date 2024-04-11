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

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendUnauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("Content=Type", "application/json")
	responseFactory := model.NewErrorResponseModelFactory()
	response := responseFactory.CreateErrorResponse("Unauthorized access")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendSuccessResponse(w http.ResponseWriter, message string, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseFactory := model.NewSuccessResponseModelFactory()
	response := responseFactory.CreateSuccessResponse(
		message,
		value,
	)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}

func sendSuccessLoginResponse(w http.ResponseWriter, message string, data interface{}, token string) {
	response := map[string]interface{}{
		"status":  "200",
		"message": message,
		"data":    data,
		"token":   token, // Include the JWT token in the response
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		sendErrorResponse(w, "Failed to marshal JSON response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

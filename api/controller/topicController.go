package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllTopic...
func GetAllTopic(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM thread"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var topic model.Topic
	var topics []model.Topic
	for rows.Next() {
		if err := rows.Scan(
			&topic.TopicNo, &topic.TopicTitle, &topic.TopicDesc, &topic.CreateDate, &topic.BanStatus); err != nil {
			log.Println(err)
			return
		} else {
			topics = append(topics, topic)
		}
	}

	if len(topics) > 1 {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = topics
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Error"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Insert Topic...
func InsertTopic(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	topicTitle, _ := strconv.Atoi(r.Form.Get("topicTitle"))
	topicDesc, _ := strconv.Atoi(r.Form.Get("topicDesc"))

	_, errQuery := db.Exec("INSERT INTO transactions(topicTitle, topicDesc) values (?,?)",
		topicTitle,
		topicDesc,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Error"
		json.NewEncoder(w).Encode(response)
	}
	w.Header().Set("Content-Type", "application/json")
}

// Update Topic Title...
func UpdateTopicTitle(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	vars := mux.Vars(r)
	topicNo := vars["topicNo"]

	topicTitle, _ := strconv.Atoi(r.Form.Get("topicTitle"))

	sqlStatement := `
		UPDATE topic 
		SET topicTitle = ?
		WHERE topicNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		topicTitle,
		topicNo,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Error"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update Topic Description...
func UpdateTopicDescription(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	vars := mux.Vars(r)
	topicNo := vars["topicNo"]

	topicDesc, _ := strconv.Atoi(r.Form.Get("topicDesc"))

	sqlStatement := `
		UPDATE topic 
		SET topicDesc = ?
		WHERE topicNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		topicDesc,
		topicNo,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Error"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update Topic Status...
func UpdateTopicStatus(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	vars := mux.Vars(r)
	topicNo := vars["topicNo"]

	banStatus, _ := strconv.Atoi(r.Form.Get("banStatus"))

	sqlStatement := `
		UPDATE topic 
		SET banStatus =  ?
		WHERE topicNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		banStatus,
		topicNo,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Error"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Delete Topic
func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	topicNo := vars["topicNo"]

	_, errQuery := db.Exec("DELETE FROM topic WHERE topicNo=?",
		topicNo,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Error"
		json.NewEncoder(w).Encode(response)
	}
	w.Header().Set("Content-Type", "application/json")
}

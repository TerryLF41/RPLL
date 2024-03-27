package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Get semua topic yang ada
func GetAllTopic(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM topic"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var topic model.Topic
	var topicList []model.Topic
	for rows.Next() {
		if err := rows.Scan(
			&topic.TopicNo, &topic.TopicTitle, &topic.TopicDesc, &topic.CreateDate, &topic.BanStatus); err != nil {
			log.Println(err)
			return
		} else {
			topicList = append(topicList, topic)
		}
	}

	if len(topicList) > 1 {
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to get data from the database")
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Insert sebuah topic baru
func InsertTopic(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}

	topicFactory := model.NewTopicModelFactory()

	// Create a new topic instance
	newTopic := topicFactory.CreateTopic(
		0,
		r.Form.Get("topicTitle"),
		r.Form.Get("topicDesc"),
		time.Now(),
		false,
		nil,
	)

	_, errQuery := db.Exec("INSERT INTO transactions(topicTitle, topicDesc) values (?,?)",
		newTopic.TopicTitle,
		newTopic.TopicDesc,
	)

	if errQuery == nil {
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to insert new topic")
		json.NewEncoder(w).Encode(response)
	}
	w.Header().Set("Content-Type", "application/json")
}

// Update judul/title sebuah topic
func UpdateTopicTitle(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
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
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to update topic title")
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update description sebuah topic
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
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to update topic description")
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update status topic menjadi 0(unbanned) atau 1(banned)
// Digunakan oleh admin untuk menban topic offensive
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
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to update topic ban status")
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Delete sebuah topic
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
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to delete topic")
		json.NewEncoder(w).Encode(response)
	}
	w.Header().Set("Content-Type", "application/json")
}

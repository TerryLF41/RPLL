package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get semua thread yang ada
func GetAllThreads(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM thread"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var thread model.Thread
	var threadList []model.Thread
	for rows.Next() {
		if err := rows.Scan(&thread.ThreadNo, &thread.ThreadTitle, &thread.ThreadDesc,
			&thread.CreateDate, &thread.BanStatus); err != nil {
			log.Println(err)
			return
		} else {
			threadList = append(threadList, thread)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if len(threadList) >= 1 {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = threadList
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Failed to get data from the database"
		json.NewEncoder(w).Encode(response)
	}

	json.NewEncoder(w).Encode(threadList)
}

// Insert sebuah thread baru
func InsertThread(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to insert a new thread")
		return
	}
	threadTitle := r.Form.Get("title")
	topicNo, _ := strconv.Atoi(r.Form.Get("topik_No"))
	threadDesc := r.Form.Get("deskripsi")

	_, errQuery := db.Exec("INSERT INTO thread(threadTitle, topicNo, threadDesc) values (?,?,?)",
		threadTitle,
		topicNo,
		threadDesc,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Failed to insert new topic"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update judul/title sebuah thread
func UpdateThreadTitle(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)
	threadId := vars["threadNo"]

	title := r.Form.Get("title")

	sqlStatement := `
		UPDATE thread 
		SET threadTitle = ?
		WHERE threadNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		title,
		threadId,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Failed to update thread title"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update description sebuah thread
func UpdateThreadDesc(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	vars := mux.Vars(r)
	threadId := vars["threadNo"]

	desc := r.Form.Get("desc")

	sqlStatement := `
		UPDATE thread 
		SET threadDesc = ?
		WHERE threadNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		desc,
		threadId,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Failed to update thread description"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Update status thread menjadi 0(unbanned) atau 1(banned)
// Digunakan oleh admin untuk menban thread offensive
func UpdateThreadBanStatus(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	vars := mux.Vars(r)
	threadId := vars["threadNo"]

	ban, _ := strconv.Atoi(r.Form.Get("ban"))

	sqlStatement := `
		UPDATE thread 
		SET banStatus = ?
		WHERE threadNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		ban,
		threadId,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Failed to update thread ban status"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Delete sebuah thread
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	threadId := vars["threadNo"]

	_, errQuery := db.Exec("DELETE FROM thread WHERE threadNo=?",
		threadId,
	)

	if errQuery == nil {
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "Failed to delete thread"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

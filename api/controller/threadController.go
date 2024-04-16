package controller

import (
	"RPLL/api/model"
	"log"
	"net/http"
	"strconv"
	"time"

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
		if err := rows.Scan(&thread.ThreadNo, &thread.TopicNo, &thread.ThreadTitle,
			&thread.ThreadDesc, &thread.CreateDate, &thread.BanStatus); err != nil {
			log.Println(err)
			return
		} else {
			threadList = append(threadList, thread)
		}
	}

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if len(threadList) >= 1 {
		sendSuccessResponse(w, "Successfully retrieved thread", threadList)
	} else {
		sendErrorResponse(w, "Failed to retrieve thread")
	}
}

func GetTopicThreads(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)

	topicId := vars["topicno"]

	sqlStatement := `
		Select * 
		From thread
		where topicNo = ` + topicId

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
		return
	}

	var thread model.Thread
	var threadList []model.Thread
	for rows.Next() {
		if err := rows.Scan(&thread.ThreadNo, &thread.TopicNo, &thread.ThreadTitle,
			&thread.ThreadDesc, &thread.CreateDate, &thread.BanStatus); err != nil {
			log.Println(err)
			return
		} else {
			threadList = append(threadList, thread)
		}
	}

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if len(threadList) >= 1 {
		sendSuccessResponse(w, "Successfully retrieved thread", threadList)
	} else {
		sendErrorResponse(w, "Failed to retrieve thread")
	}
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
	topicNo, _ := strconv.Atoi(r.Form.Get("topicNo"))
	threadFactory := model.NewThreadModelFactory()

	// Create a new thread instance
	newThread := threadFactory.CreateThread(
		1,
		topicNo,
		r.Form.Get("threadTitle"),
		r.Form.Get("threadDesc"),
		time.Now(),
		false,
		nil,
	)

	sqlStatement := "INSERT INTO thread(threadTitle, topicNo, threadDesc) values (?,?,?)"
	_, errQuery := db.Exec(sqlStatement,
		newThread.ThreadTitle,
		newThread.TopicNo,
		newThread.ThreadDesc,
	)

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new thread", nil)
	} else {
		sendErrorResponse(w, "Failed to insert thread to database")
	}
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated thread title", nil)
	} else {
		sendErrorResponse(w, "Failed to update thread title")
	}
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated thread description", nil)
	} else {
		sendErrorResponse(w, "Failed to update thread description")
	}
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
	threadNo := vars["threadNo"]

	banStatus, _ := strconv.Atoi(r.Form.Get("banStatus"))

	sqlStatement := `
		UPDATE thread 
		SET banStatus = ?
		WHERE threadNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		banStatus,
		threadNo,
	)

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated thread ban status", nil)
	} else {
		sendErrorResponse(w, "Failed to update thread ban status")
	}
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully deleted thread", nil)
	} else {
		sendErrorResponse(w, "Failed to delete thread")
	}
}

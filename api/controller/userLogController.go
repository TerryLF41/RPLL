package controller

import (
	"RPLL/api/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// File ini memuat function-function yang berkaitan dengan userLog
// Asumsi tidak ada update log karena log harus bersifat apa adanya
// Delete juga tidak ada karena log sebaiknya selalu ada untuk keperluan tertentu

// Select userLog dari seorang user dari database berdasarkan IDnya
func GetUserLogUsingId(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Ambil user id dari mux vars
	vars := mux.Vars(r)
	userId := vars["userId"]

	// Query ke SQL
	query := "SELECT * FROM userlog WHERE userId = '" + userId + "' ORDER BY logTime DESC"

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var userLog model.UserLog
	var userLogList []model.UserLog

	for rows.Next() {
		if err := rows.Scan(&userLog.LogNo, &userLog.UserID, &userLog.LogType, &userLog.IpAddress, &userLog.LogTime); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			userLogList = append(userLogList, userLog)
		}
	}

	sendSuccessResponse(w, "Successfully retrieved userlog", userLogList)
}

// Insert/create sebuah userLog baru
func InsertUserLog(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Ambil data log dari form HTML
	// logNo tidak diperlukan karena sudah auto increment di database
	// logTime juga sudah diatur default timestamp
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}
	userId, _ := strconv.Atoi(r.Form.Get("userId"))

	userLogFactory := model.NewUserLogModelFactory()

	// Create a new user log instance
	newUserLog := userLogFactory.CreateUserLog(
		0,
		userId,
		r.Form.Get("logType"),
		r.Form.Get("ipAddress"),
		time.Now(),
	)

	// Query ke SQL
	_, errQuery := db.Exec("INSERT INTO userlog(userId,logType,ipAddress)values(?,?,?)",
		newUserLog.UserID,
		newUserLog.LogType,
		newUserLog.IpAddress,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new userlog", nil)
		log.Println("Succ")
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert userLog to database")
	}
}

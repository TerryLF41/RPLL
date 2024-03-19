package controller

import (
	"RPLL/api/model"
	"log"
	"net/http"
)

// File ini memuat function-function yang berkaitan dengan userLog
// Asumsi tidak ada update log karena log harus bersifat apa adanya
// Delete juga tidak ada karena log sebaiknya selalu ada untuk keperluan tertentu

// Select userLog dari seorang user dari database berdasarkan IDnya
func GetUserLogUsingId(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Ambil user id dari form HTML
	userId := r.Form.Get("userId")

	// Query ke SQL
	query := "SELECT * FROM userlog WHERE userId = " + userId

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var userLog model.UserLog

	for rows.Next() {
		if err := rows.Scan(&userLog.LogNo, &userLog.UserID, &userLog.LogType, &userLog.IpAddress, &userLog.LogTime); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		}
	}

	sendSuccessResponse(w, "Success", userLog)
}

// Insert/create sebuah userLog baru
func InsertUserLog(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Ambil data log dari form HTML
	// logNo tidak diperlukan karena sudah auto increment di database
	// logTime juga sudah diatur default timestamp
	userId := r.Form.Get("userId")
	logType := r.Form.Get("logType")
	IpAddress := r.Form.Get("ipAddress")

	// Query ke SQL
	_, errQuery := db.Exec("INSERT INTO userlog(userId,logType,ipAddress)values(?,?,?)",
		userId,
		logType,
		IpAddress,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new userLog", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert userLog")
	}
}

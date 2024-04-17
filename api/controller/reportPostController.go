package controller

import (
	"RPLL/api/model"
	"log"
	"net/http"
	// "strconv"
	// "time"
)

func GetAllReportPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Query ke SQL
	query := "SELECT * FROM reportpost"

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var reportPost model.ReportPost
	var reportPostList []model.ReportPost
	for rows.Next() {
		if err := rows.Scan(&reportPost.PostNo, &reportPost.UserID, &reportPost.ReportText, &reportPost.ReportDate, &reportPost.ReportStatus); err != nil {
			log.Println(err)
			return
		} else {
			reportPostList = append(reportPostList, reportPost)
		}
	}

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if len(reportPostList) >= 1 {
		sendSuccessResponse(w, "Successfully retrieved report", reportPostList)
	} else {
		sendErrorResponse(w, "Failed to retrieve report")
	}
}

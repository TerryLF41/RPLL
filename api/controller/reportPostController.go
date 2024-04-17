package controller

import (
	"RPLL/api/model"
	"log"
	"net/http"

	// "strconv"

	"github.com/gorilla/mux"
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

// Update status report menjadi 0(unsolved) atau 1(solved)
// Digunakan oleh admin untuk menban post offensive
func UpdateReportStatus(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)
	postNo := vars["postNo"]
	// report, _ := strconv.Atoi(r.Form.Get("report"))

	sqlStatement := `
		UPDATE reportpost 
		SET reportStatus = 1
		WHERE postNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		postNo,
	)

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated report status", nil)
	} else {
		sendErrorResponse(w, "Failed to update report status")
	}
}

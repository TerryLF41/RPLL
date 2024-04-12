package controller

import (
	"RPLL/api/model"
	_ "image"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Get semua topic yang ada
func GetAllTopic(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var topicList []model.Topic

	searchType := r.URL.Query().Get("orderBy")

	context := &SearchContext{}

	if searchType == "time" {
		// Cari dari topic yang paling baru dibuat
		context.SetStrategy(&LatestTopicSearchStrategy{})
		topicList = context.PerformSearch("SELECT * FROM topic ORDER BY createDate DESC")
	} else {
		// Cari dari dengan jumlah thread terbanyak
		// TODO : Ganti query agar bisa cari berdasarkan jumlah thread terbanyak
		context.SetStrategy(&PopularitySearchStrategy{})
		topicList = context.PerformSearch("SELECT * FROM topic ORDER BY createDate ASC")
	}

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if len(topicList) >= 1 {
		sendSuccessResponse(w, "Successfully retrieved topic", topicList)
	} else {
		sendErrorResponse(w, "Failed to get data from the database")
	}
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

	log.Println(r.Form.Get("topicPicturePath"))

	// Create a new topic instance
	newTopic := topicFactory.CreateTopic(
		0,
		r.Form.Get("topicTitle"),
		r.Form.Get("topicDesc"),
		time.Now(),
		false,
		r.Form.Get("topicPicture"),
		nil,
	)

	_, errQuery := db.Exec("INSERT INTO topic(topicTitle, topicDesc, topicPicture) values (?,?,?)",
		newTopic.TopicTitle,
		newTopic.TopicDesc,
		newTopic.TopicPicture,
	)

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new topic", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert topic to database")
	}
}

func SaveTopicPicture(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	//fileName := r.FormValue("file_name")
	//filepath := "../vue-frontend/src/assets/userUploadedFiles/topicPicture/" + fileName
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	log.Println("Succ here")
	defer f.Close()
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated topic title", nil)
	} else {
		sendErrorResponse(w, "Failed to update topic title")
	}
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated topic description", nil)
	} else {
		sendErrorResponse(w, "Failed to update topic description")
	}
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated topic ban status", nil)
	} else {
		sendErrorResponse(w, "Failed to update topic ban status")
	}
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

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully deleted topic", nil)
	} else {
		sendErrorResponse(w, "Failed to delete topic")
	}
}

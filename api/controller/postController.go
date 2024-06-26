package controller

import (
	"RPLL/api/model"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Get semua post berdasarkan threadNo
func GetAllPostByThreadNo(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)

	threadNo := vars["threadNo"]

	// Query ke SQL
	query := "SELECT * FROM post WHERE threadNo = '" + threadNo + "'"

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var post model.Post
	var postList []model.Post
	for rows.Next() {
		if err := rows.Scan(&post.PostNo, &post.ThreadNo, &post.UserId, &post.ReplyTo, &post.PostText, &post.PostImage, &post.PostDate, &post.BanStatus); err != nil {
			log.Println(err)
			return
		} else {
			postList = append(postList, post)
		}
	}

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if len(postList) >= 1 {
		sendSuccessResponse(w, "Successfully retrieved post", postList)
	} else {
		sendErrorResponse(w, "Failed to retrieve post")
	}
}

// Insert sebuah post baru
func InsertPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Ambil data log dari form HTML
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}
	threadNo, _ := strconv.Atoi(r.Form.Get("threadNo"))
	userId, _ := strconv.Atoi(r.Form.Get("userId"))

	var replyTo *int

	// Get the value from the form
	replyValue := r.Form.Get("replyTo")

	// Check if the value is empty
	if replyValue == "" {
		replyTo = nil
	} else {
		// Convert the string to an int
		if val, err := strconv.Atoi(replyValue); err == nil {
			replyTo = &val
		} else {
			log.Println("Error converting replyTo:", err)
			replyTo = nil
		}
	}

	postFactory := model.NewPostModelFactory()

	// Create a new post instance
	newPost := postFactory.CreatePost(
		0,
		threadNo,
		userId,
		replyTo,
		r.Form.Get("postText"),
		r.Form.Get("postImage"),
		time.Now(),
		false,
	)

	// Query ke SQL
	_, errQuery := db.Exec("INSERT INTO post(threadNo,userId,reply,postText,postImage)values(?,?,?,?,?)",
		newPost.ThreadNo,
		newPost.UserId,
		newPost.ReplyTo,
		newPost.PostText,
		newPost.PostImage,
	)

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new post", nil)
	} else {
		sendErrorResponse(w, "Failed to insert post to database")
		log.Println(errQuery)
	}
}

func SavePostPicture(w http.ResponseWriter, r *http.Request) {
	postId, _ := strconv.Atoi(r.FormValue("postNo"))
	formFileId := "postImage"

	if postId != 0 {
		formFileId = "postImage" + r.FormValue("postNo")
	}

	file, _, _ := r.FormFile(formFileId)

	filename := r.FormValue("filename")
	//open a file for writing
	filepath := "vue-frontend/src/assets/userUploadedFiles/postPicture/" + filename
	log.Println("Filepath = ", filepath)

	// Buat file temp yang akan menyimpan gambar nantinya
	tempFile, err := os.Create(filepath)

	if err != nil {
		log.Println(err)
	}

	// Copy data gambar dari variable fileke path yang telah dibuat
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(tempFile, file)
	if err == nil {
		sendSuccessResponse(w, "Successfully saved post picture", nil)
	} else {
		log.Println(err)
		sendErrorResponse(w, "Failed to save post picture")
	}
	tempFile.Close()
}

// Update status post menjadi 0(unbanned) atau 1(banned)
// Digunakan oleh admin untuk menban post offensive
func UpdatePostBanStatus(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}

	vars := mux.Vars(r)

	postNo := vars["postNo"]

	banStatus, _ := strconv.Atoi(r.Form.Get("banStatus"))

	sqlStatement := `
		UPDATE post 
		SET banStatus = ?
		WHERE postNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		banStatus,
		postNo,
	)

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated post ban status", nil)
	} else {
		sendErrorResponse(w, "Failed to update post ban status")
	}
}

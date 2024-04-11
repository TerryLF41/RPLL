package controller

import (
	"RPLL/api/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Get semua post berdasarkan threadNo
func GetAllPostByThreadNo(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Ambil thread number dari url query
	threadNo := r.URL.Query().Get("threadNo")

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
	replyTo, _ := strconv.Atoi(r.Form.Get("replyTo"))

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
	_, errQuery := db.Exec("INSERT INTO post(threadNo,userId,replyTo,postText,postImage)values(?,?,?,?,?)",
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
	}
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

	postNo, _ := strconv.Atoi(r.Form.Get("postNo"))
	ban, _ := strconv.Atoi(r.Form.Get("ban"))

	sqlStatement := `
		UPDATE post 
		SET banStatus = ?
		WHERE postNo = ?`

	_, errQuery := db.Exec(sqlStatement,
		ban,
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

package controller

import (
	"RPLL/api/model"
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")
	if len(postList) >= 1 {
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			postList,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to get data from the database")
		json.NewEncoder(w).Encode(response)
	}

	json.NewEncoder(w).Encode(postList)
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

	if errQuery == nil {
		sendSuccessResponse(w, "Success", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert post")
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

	if errQuery == nil {
		responseFactory := model.NewSuccessResponseModelFactory()

		response := responseFactory.CreateSuccessResponse(
			"success",
			nil,
		)
		json.NewEncoder(w).Encode(response)
	} else {
		responseFactory := model.NewErrorResponseModelFactory()

		response := responseFactory.CreateErrorResponse("Failed to update post ban status")
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

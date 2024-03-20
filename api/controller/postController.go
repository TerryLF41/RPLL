package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GET all posts
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
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = postList
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "failed to get data from the database"
		json.NewEncoder(w).Encode(response)
	}

	json.NewEncoder(w).Encode(postList)
}

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
	postText := r.Form.Get("postText")
	postImage := r.Form.Get("postImage")

	// Query ke SQL
	_, errQuery := db.Exec("INSERT INTO post(threadNo,userId,replyTo,postText,postImage)values(?,?,?,?,?)",
		threadNo,
		userId,
		replyTo,
		postText,
		postImage,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully inserted new Post", nil)
	} else {
		log.Println(errQuery)
		sendErrorResponse(w, "Failed to insert post")
	}
}

func UpdatePostBanStatus(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
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
		var response model.GenericResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.ErrorResponse
		response.Status = 400
		response.Message = "fail to update post ban status"
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
}

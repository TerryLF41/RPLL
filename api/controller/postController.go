package controller

import (
	"RPLL/api/model"
	"encoding/json"
	"log"
	"net/http"
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

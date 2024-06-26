package main

import (
	"RPLL/api/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	//Note untuk UserType
	//0 = User biasa
	//1 = Admin

	// Handler untuk user
	router.HandleFunc("/user", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/ban/{userId}", controller.BanUser).Methods("POST")
	router.HandleFunc("/user/picture", controller.SaveProfilePicture).Methods("POST")
	router.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")
	router.HandleFunc("/password/{userId}", controller.ChangePassword).Methods("PUT")
	router.HandleFunc("/profile/{userId}", controller.ChangeProfile).Methods("PUT")

	// Handler untuk userLog
	router.HandleFunc("/userLog/{userId}", controller.GetUserLogUsingId).Methods("GET")
	router.HandleFunc("/userLog", controller.InsertUserLog).Methods("POST")

	// Handler untuk topic
	router.HandleFunc("/topic/{orderBy}", controller.GetAllTopic).Methods("GET")
	router.HandleFunc("/topic", controller.InsertTopic).Methods("POST")
	router.HandleFunc("/topic/picture", controller.SaveTopicPicture).Methods("POST")
	router.HandleFunc("/topic/title/{topicNo}", controller.UpdateTopicDescription).Methods("PUT")
	router.HandleFunc("/topic/desc/{topicNo}", controller.UpdateTopicDescription).Methods("PUT")
	router.HandleFunc("/topic/ban/{topicNo}", controller.UpdateTopicStatus).Methods("PUT")
	router.HandleFunc("/topic/{topicNo}", controller.GetUserLogUsingId).Methods("DELETE")

	// Handler untuk thread
	router.HandleFunc("/thread", controller.GetAllThreads).Methods("GET")
	router.HandleFunc("/thread/{topicno}", controller.GetTopicThreads).Methods("GET")
	router.HandleFunc("/thread", controller.InsertThread).Methods("POST")
	router.HandleFunc("/thread/Title/{threadNo}", controller.UpdateThreadTitle).Methods("PUT")
	router.HandleFunc("/thread/Desc/{threadNo}", controller.UpdateThreadDesc).Methods("PUT")
	router.HandleFunc("/thread/ban/{threadNo}", controller.UpdateThreadBanStatus).Methods("PUT")
	router.HandleFunc("/thread/{threadNo}", controller.DeleteThread).Methods("DELETE")

	// Handler untuk post
	router.HandleFunc("/post/{threadNo}", controller.GetAllPostByThreadNo).Methods("GET")
	router.HandleFunc("/post", controller.InsertPost).Methods("POST")
	router.HandleFunc("/post/picture", controller.SavePostPicture).Methods("POST")
	router.HandleFunc("/post/ban/{postNo}", controller.UpdatePostBanStatus).Methods("PUT")

	// Handler untuk reportpost
	router.HandleFunc("/reportpost", controller.GetAllReportPost).Methods("GET")
	router.HandleFunc("/reportpostu", controller.GetAllReportPostU).Methods("GET")
	router.HandleFunc("/reportpost", controller.InsertReport).Methods("POST")
	router.HandleFunc("/reportpost/resolve/{postNo}", controller.SolveReport).Methods("PUT")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8181"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	fmt.Println("Connected to port 8181")
	log.Println("Connected to port 8181")

	log.Fatal(http.ListenAndServe(":8181", handler))
}

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

	//User
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")
	router.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/password", controller.Authenticate(controller.ChangePassword, 1)).Methods("PUT")
	router.HandleFunc("/registerseller", controller.Authenticate(controller.RegisterSeller, 1)).Methods("PUT")

	// Handler untuk userLog
	router.HandleFunc("/userLog", controller.GetUserLogUsingId).Methods("GET")
	router.HandleFunc("/userLog", controller.InsertUserLog).Methods("POST")

	// Handler untuk topic
	router.HandleFunc("/topic", controller.GetAllTopic).Methods("GET")
	router.HandleFunc("/topic", controller.InsertTopic).Methods("POST")
	router.HandleFunc("/topic/Title/{topicNo}", controller.UpdateTopicDescription).Methods("PUT")
	router.HandleFunc("/topic/Desc/{topicNo}", controller.UpdateTopicDescription).Methods("PUT")
	router.HandleFunc("/topic/Status/{topicNo}", controller.UpdateTopicStatus).Methods("PUT")
	router.HandleFunc("/topic/{topicNo}", controller.GetUserLogUsingId).Methods("DELETE")

	// Handler untuk thread
	router.HandleFunc("/thread", controller.GetAllThreads).Methods("GET")
	router.HandleFunc("/thread", controller.InsertThread).Methods("POST")
	router.HandleFunc("/thread/Title/{threadNo}", controller.UpdateThreadTitle).Methods("PUT")
	router.HandleFunc("/thread/Desc/{threadNo}", controller.UpdateThreadDesc).Methods("PUT")
	router.HandleFunc("/thread/Status/{threadNo}", controller.UpdateThreadBanStatus).Methods("PUT")
	router.HandleFunc("/thread/{threadNo}", controller.DeleteThread).Methods("DELETE")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:8181"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	http.Handle("/", router)

	fmt.Println("Connected to port 8181")
	log.Println("Connected to port 8181")

	log.Fatal(http.ListenAndServe(":8181", handler))
}

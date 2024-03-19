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

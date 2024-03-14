package main

import (
	"PBP-Tubes-API-Tokopedia/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	//Scheduler monthly report
	controller.SetMonthlyReportScheduler()

	//Note
	//0 = Admin
	//1 = Pembeli
	//2 = Penjual

	//User
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")
	router.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/password", controller.Authenticate(controller.ChangePassword, 1)).Methods("PUT")
	router.HandleFunc("/registerseller", controller.Authenticate(controller.RegisterSeller, 1)).Methods("PUT")

	//Cart
	router.HandleFunc("/cart", controller.Authenticate(controller.GetCart, 1)).Methods("GET")
	router.HandleFunc("/cart", controller.Authenticate(controller.InsertItemToCart, 1)).Methods("POST")
	router.HandleFunc("/cart", controller.Authenticate(controller.UpdateCart, 1)).Methods("PUT")
	router.HandleFunc("/cart/{item_id}", controller.Authenticate(controller.DeleteItemFromCart, 1)).Methods("DELETE")

	//Product
	//Produk bisa bisa dilihat siapapunn
	router.HandleFunc("/item", controller.GetItem).Methods("GET")
	router.HandleFunc("/item", controller.Authenticate(controller.InsertItem, 2)).Methods("POST")
	router.HandleFunc("/item/{item_id}", controller.Authenticate(controller.UpdateItem, 2)).Methods("PUT")
	router.HandleFunc("/item/{item_id}", controller.Authenticate(controller.DeleteItem, 2)).Methods("DELETE")

	//Shop
	//Shop profile bisa dilihat siapapun
	router.HandleFunc("/shop", controller.GetShopProfile).Methods("GET")
	router.HandleFunc("/shoplist", controller.Authenticate(controller.GetUserShop, 2)).Methods("GET")
	router.HandleFunc("/shop", controller.Authenticate(controller.RegisterShop, 2)).Methods("POST")
	router.HandleFunc("/shop/{shop_id}", controller.Authenticate(controller.UpdateShopProfile, 2)).Methods("PUT")
	router.HandleFunc("/shop_admin", controller.Authenticate(controller.InsertShopAdmin, 2)).Methods("POST")

	//Transaction
	router.HandleFunc("/transaction", controller.Authenticate(controller.GetAllTransaction, 1)).Methods("GET")
	router.HandleFunc("/transaction", controller.Authenticate(controller.InsertItemToTransaction, 1)).Methods("POST")
	//Update transaksi hanya bisa dilakukan penjual saja
	router.HandleFunc("/transaction/{transaction_id}", controller.Authenticate(controller.UpdateTransaction, 2)).Methods("PUT")

	//Profile
	router.HandleFunc("/profile", controller.Authenticate(controller.GetUserProfile, 1)).Methods("GET")
	router.HandleFunc("/profile", controller.Authenticate(controller.UpdateUserProfile, 1)).Methods("PUT")

	//Review
	router.HandleFunc("/review", controller.Authenticate(controller.ReviewItem, 1)).Methods("POST")
	//Review sebuah produk bisa dilihat siapapun
	router.HandleFunc("/review", controller.GetItemReview).Methods("GET")

	//Admin
	router.HandleFunc("/banshop/{shop_id}", controller.Authenticate(controller.BanShop, 0)).Methods("PUT")
	router.HandleFunc("/banuser/{user_id}", controller.Authenticate(controller.BanUser, 0)).Methods("PUT")
	router.HandleFunc("/useradmin", controller.Authenticate(controller.GetUser, 0)).Methods("GET")

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

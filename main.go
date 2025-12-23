package main

import (
	"log"
	"net/http"

	"halo-streetwear/internal/db"
	"halo-streetwear/internal/handlers"
)

func main() {
	database := db.Connect()
	handlers.DB = database

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/login/email", handlers.EmailLogin)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/checkout", handlers.CheckoutPage)
	http.HandleFunc("/pay", handlers.FakePayment)
	http.HandleFunc("/success", handlers.SuccessPage)

	log.Println("HALO running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

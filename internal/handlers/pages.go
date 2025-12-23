package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/about.html")
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, "templates/dashboard.html")
}

func CheckoutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/checkout.html")
}

func SuccessPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/success.html")
}

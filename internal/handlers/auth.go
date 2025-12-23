package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func EmailLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := strings.ToLower(strings.TrimSpace(r.Form.Get("email")))
	password := strings.TrimSpace(r.Form.Get("password"))

	var hash string
	err := DB.QueryRow(
		"SELECT password_hash FROM public.users WHERE email=$1",
		email,
	).Scan(&hash)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    email,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func FakePayment(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

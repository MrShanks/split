package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}
	var user *User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	var password string

	row := db.QueryRow("SELECT password FROM users WHERE email = $1;", user.Email)

	err = row.Scan(&password)
	if err != nil {
		log.Println(err)
		return
	}

	if password != user.Password {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(map[string]string{"message": "User not found!"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to your dashboard"})
	if err != nil {
		log.Println(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/register.html")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}
	var user *User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)", user.Firstname, user.Lastname, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
	if err != nil {
		log.Println(err)
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {

}

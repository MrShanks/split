package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB
var connStr = "postgres://split:split@localhost:5432/split?sslmode=disable"

func Login(w http.ResponseWriter, r *http.Request) {
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

func main() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Connected to PostgreSQL")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/signup", Signup)

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("✅ Listening on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

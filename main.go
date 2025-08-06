package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	var firstName, lastName string

	row := db.QueryRow("SELECT first_name, last_name FROM users WHERE email = $1;", user.Email)

	err = row.Scan(&firstName, &lastName)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(firstName, lastName, "Logged in succesffully!")

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to your dashboard"})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
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

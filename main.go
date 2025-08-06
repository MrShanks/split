package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	simone := User{Firstname: "simone"}
	marco := User{Firstname: "Marco"}
	pippo := User{Firstname: "Pippo"}

	netflix := Group{
		Members: []*User{
			&simone,
			&marco,
			&pippo,
		},
	}

	netflix.AddExpense(450, "CHF", &simone, false)
	netflix.AddExpense(200, "CHF", &marco, false)
	netflix.AddExpense(100, "CHF", &pippo, false)
	netflix.AddExpense(400, "CHF", &marco, false)

	for _, m := range netflix.Members {
		fmt.Println(m.Firstname, m.Credit, m.Debt)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

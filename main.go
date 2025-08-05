package main

import (
	"fmt"
	"time"
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Debt      float64
	Credit    float64
}

type Metadata struct {
	CreatedDate  time.Time
	LastModified time.Time
}

type Expense struct {
	Amount    float64
	Currency  string
	Payer     User
	Recurrent bool
	Metadata
}

func (u *User) Add(amount float64) {
	u.Credit += amount
}

func (g *Group) AddExpense(amount float64, currency string, user *User, recurrent bool) *Expense {
	metadata := Metadata{
		CreatedDate:  time.Now(),
		LastModified: time.Now(),
	}
	expense := &Expense{
		Amount:    amount,
		Currency:  currency,
		Payer:     *user,
		Recurrent: recurrent,
		Metadata:  metadata,
	}

	for _, m := range g.Members {
		if m.Firstname == user.Firstname {
			m.Add(amount)
		}
		m.Add(-amount / float64(len(g.Members)))
	}

	return expense
}

type Group struct {
	Members []*User
}

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
}

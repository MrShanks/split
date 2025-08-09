package main

import (
	"time"
)

type User struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     string  `json:"email"`
	Debt      float64 `json:"debt"`
	Credit    float64 `json:"credit"`
	Password  string  `json:"password"`
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

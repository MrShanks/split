package main

import (
	"time"
)

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
